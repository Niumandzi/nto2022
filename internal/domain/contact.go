package domain

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

type ContactUsecase interface {
	CreateContact(ctx context.Context, contact model.Contact) (int, error)
	GetContact(ctx context.Context, contactId int) (model.Contact, error)
	GetContactByType(ctx context.Context, contactType string) (model.Contact, error)
	GetAllContacts(ctx context.Context) ([]model.Contact, error)
	UpdateContact(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error
	DeleteContact(ctx context.Context, contactId int) error
}

type ContactRepository interface {
	Create(ctx context.Context, contact model.Contact) (int, error)
	Get(ctx context.Context, contactId int) (model.Contact, error)
	GetByType(ctx context.Context, contactType string) (model.Contact, error)
	GetAll(ctx context.Context) ([]model.Contact, error)
	Update(ctx context.Context, contactId int, contactInput model.UpdateContactInput) error
	Delete(ctx context.Context, contactId int) error
}
