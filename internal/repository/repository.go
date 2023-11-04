package repository

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

type ContactRepository interface {
	Create(ctx context.Context, contact model.Contact) (int, error)
	GetById(ctx context.Context, contactId int) (model.Contact, error)
	GetByType(ctx context.Context, contactType string) ([]model.Contact, error)
	Update(ctx context.Context, contactInput model.Contact) error
	Delete(ctx context.Context, contactId int) error
}

type HotelRepository interface {
	Create(ctx context.Context, hotel model.Hotel) (int, error)
	GetById(ctx context.Context, hotelId int) (model.HotelWithContact, error)
	GetAll(ctx context.Context) ([]model.HotelWithContact, error)
	Update(ctx context.Context, hotel model.Hotel) error
	Delete(ctx context.Context, hotelId int) error
}
