package tour

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (h TourUsecase) GetHotel(ctx context.Context, tourId int) (model.TourWithHotel, error) {
	ctx, cancel := context.WithTimeout(ctx, h.contextTimeout)
	defer cancel()

	TourWithHotel, err := h.tourRepo.GetById(ctx, tourId)
	if err != nil {
		h.logger.Error(err.Error())
		return model.TourWithHotel{}, err
	}

	return TourWithHotel, nil
}