package db

import (
	"context"
	"fmt"
	"main/internal/interfaces"
	"main/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type db struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) interfaces.Storage {
	return &db{
		db: pool,
	}
}

func (d *db) Create(ctx context.Context, sub model.Subscription) (int, error) {
	id := 0
	query := `
		INSERT INTO
			subscriptions 
			(
				service_name,
				price,
				user_id,
				start_date,
				end_date
			)
		VALUES
		(
			@service_name,
			@price,
			@user_id,
			@start_date,
			@end_date
		)
		RETURNING
			id
	`
	args := pgx.NamedArgs{
		"service_name": sub.ServiceName,
		"price":        sub.Price,
		"user_id":      sub.UserId,
		"start_date":   sub.StartDate,
		"end_date":     sub.EndDate,
	}
	row := d.db.QueryRow(ctx, query, args)
	err := row.Scan(&id)
	if err != nil {
		logrus.Error("db create sub err:", err)
		return id, fmt.Errorf("db create sub error: %v", err)
	}
	return id, nil
}

func (d *db) Load(ctx context.Context, id int) (model.Subscription, error) {
	var res model.Subscription
	query := `
		SELECT 
			id,
			service_name,
			price,
			user_id,
			start_date,
			end_date
		FROM
			subscriptions
		WHERE
			id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := d.db.Query(ctx, query, args)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	res, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Subscription])

	if err != nil {
		logrus.Error("db load err:", err)
		return res, err
	}

	return res, nil
}

func (d *db) LoadList(ctx context.Context, limit int, offset int) ([]model.Subscription, error) {
	var res []model.Subscription
	query := `
		SELECT 
			id,
			service_name,
			price,
			user_id,
			start_date,
			end_date
		FROM
			subscriptions
		WHERE
			id = @id
		ORDER BY 
			id
		LIMIT 
			@limit
		OFFSET 
			@offset
	`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := d.db.Query(ctx, query, args)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	res, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.Subscription])

	if err != nil {
		logrus.Error("db load list err:", err)
		return res, err
	}
	return res, nil
}

func (d *db) Update(ctx context.Context, sub model.Subscription) error {
	query := `
		UPDATE
			subscriptions
		SET
			service_name = @upd_service_name,
			price = @upd_price,
			user_id = @upd_user_id,
			start_date = @upd_start_date,
			end_date = @upd_end_date
		FROM
			subscriptions
		WHERE
			id = @id
	`
	args := pgx.NamedArgs{
		"upd_service_name": sub.ServiceName,
		"upd_price":        sub.Price,
		"upd_user_id":      sub.UserId,
		"upd_start_date":   sub.StartDate,
		"upd_end_date":     sub.EndDate,
		"id":               sub.Id,
	}

	result, err := d.db.Exec(ctx, query, args)
	if err != nil {
		logrus.Error("db update err:", err)
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (d *db) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM
			subscriptions
		WHERE
			id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	result, err := d.db.Exec(ctx, query, args)
	if err != nil {
		logrus.Error("db delete err:", err)
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
