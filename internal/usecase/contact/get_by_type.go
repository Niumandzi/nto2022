package contact

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/niumandzi/nto2022/model"
)

func (c ContactUsecase) GetContactsByType(ctx context.Context, contactType string) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := validation.Validate(contactType, validation.Required, validation.In("worker", "private_client", "legal_client"))
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}

	contacts, err := c.contactRepo.GetByType(ctx, contactType)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	return contacts, nil
}
