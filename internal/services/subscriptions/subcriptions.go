package subscriptions

import "main/internal/interfaces"

type sub struct {
	storage interfaces.Storage
}

func New(s interfaces.Storage) interfaces.Subscriptions {
	return &sub{
		storage: s,
	}
}
