package contact

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (c ContactUsecase) GetAllContacts() ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.contextTimeout)
	defer cancel()

	contacts, err := c.contactRepo.GetAll(ctx)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	return contacts, nil
}
