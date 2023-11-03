package contact

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/niumandzi/nto2022/model"
	"regexp"
)

func (c ContactUsecase) CreateContact(contact model.Contact) (int, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.contextTimeout)
	defer cancel()

	err := validation.ValidateStruct(&contact,
		validation.Field(&contact.ContactType, validation.Required, validation.In("worker", "private_client", "legal_client")),
		validation.Field(&contact.Name, validation.Required),
		validation.Field(&contact.Number, validation.Required, validation.Match(regexp.MustCompile(`^\+\d{1,2}\s\(\d{3}\)\s\d{3}-\d{2}-\d{2}$`))),
		validation.Field(&contact.Email, validation.Required, is.Email),
	)
	if err != nil {
		c.logger.Error(err.Error())
		return 0, err
	}

	id, err := c.contactRepo.Create(ctx, contact)
	if err != nil {
		c.logger.Error(err.Error())
		return 0, err
	}
	return id, nil
}
