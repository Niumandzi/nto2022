package usecase

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

type ContactUseCase interface {
	CreateContact(ctx context.Context, contact model.Contact) (int, error)
	GetContact(ctx context.Context, contactId int) (model.Contact, error)
	GetContactsByType(ctx context.Context, contactType string) ([]model.Contact, error)
	GetAllContacts(ctx context.Context) ([]model.Contact, error)
	UpdateContact(ctx context.Context, contactInput model.Contact) error
	DeleteContact(ctx context.Context, contactId int) error
}

type HotelUseCase interface {
	CreateHotel(ctx context.Context, hotel model.Hotel) (int, error)
	GetHotel(ctx context.Context, hotelId int) (model.HotelWithContact, error)
	GetAllHotels(ctx context.Context) ([]model.HotelWithContact, error)
	UpdateHotel(ctx context.Context, hotel model.Hotel) error
	DeleteHotel(ctx context.Context, hotelId int) error
}

type UseCases struct {
	Contact ContactUseCase
	Hotel   HotelUseCase
}

func NewUsecases(contact ContactUseCase, hotel HotelUseCase) *UseCases {
	return &UseCases{
		Contact: contact,
		Hotel:   hotel,
	}
}
