package contact

import (
	"context"
	"github.com/niumandzi/nto2022/internal/repository"
	"github.com/niumandzi/nto2022/pkg/logging"
	"time"
)

type ContactUsecase struct {
	contactRepo    repository.ContactRepository
	contextTimeout time.Duration
	logger         logging.Logger
	ctx            context.Context
}

func NewContacUsecase(contact repository.ContactRepository, timeout time.Duration, logger logging.Logger, ctx context.Context) ContactUsecase {
	return ContactUsecase{
		contactRepo:    contact,
		contextTimeout: timeout,
		logger:         logger,
		ctx:            ctx,
	}
}
