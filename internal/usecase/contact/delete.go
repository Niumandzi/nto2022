package contact

import "context"

func (c ContactUsecase) DeleteContact(contactId int) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.contextTimeout)
	defer cancel()

	err := c.contactRepo.Delete(ctx, contactId)
	if err != nil {
		return err
	}
	return nil
}
