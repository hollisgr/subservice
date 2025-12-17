package db

import (
	"context"
	"fmt"
	"main/internal/dto"
	"main/internal/interfaces"
	"main/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
		return id, fmt.Errorf("db create sub query err: %v", err)
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
		return res, fmt.Errorf("db load sub query error: %v", err)
	}

	res, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Subscription])

	if err != nil {
		if err == pgx.ErrNoRows {
			return res, err
		}
		return res, fmt.Errorf("db load sub collect row error: %v", err)
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
		return res, fmt.Errorf("db load sub list query error: %v", err)
	}

	res, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.Subscription])

	if err != nil {
		return res, fmt.Errorf("db load sub list collect error: %v", err)
	}

	if len(res) == 0 {
		return res, pgx.ErrNoRows
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
		return fmt.Errorf("db update sub exec error: %v", err)
	}

	if result.RowsAffected() == 0 {
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
		return fmt.Errorf("db delete sub exec error: %v", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (d *db) Cost(ctx context.Context, data dto.CostRequestToDB) (model.Subscription, error) {
	res := model.Subscription{}
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
			user_id = @user_id
			AND 
				service_name = @service_name
			AND 
				(@start_date BETWEEN start_date AND end_date 
				OR 
				@end_date BETWEEN start_date AND end_date)
	`
	args := pgx.NamedArgs{
		"service_name": data.ServiceName,
		"user_id":      data.UserId,
		"start_date":   data.StartDate,
		"end_date":     data.EndDate,
	}

	rows, err := d.db.Query(ctx, query, args)
	defer rows.Close()

	if err != nil {
		return res, fmt.Errorf("db cost sub query error: %v", err)
	}

	res, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Subscription])

	if err != nil {
		if err == pgx.ErrNoRows {
			return res, err
		}
		return res, fmt.Errorf("db cost sub collect row error: %v", err)
	}

	return res, nil
}
