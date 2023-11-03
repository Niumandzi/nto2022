package tour

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	errcode "github.com/niumandzi/nto2022/internal/errors"
	"github.com/niumandzi/nto2022/model"
	"github.com/niumandzi/nto2022/pkg/logging"
)

type ContactRepository struct {
	db     *sql.DB
	logger logging.Logger
}

func (h TourRepository) GetById(ctx context.Context, tourId int) (model.TourWithHotel, error) {
	row := h.db.QueryRowContext(ctx, "SELECT * FROM tour LEFT JOIN contact ON tour.HotelId = hotel.id WHERE tour.id = ?", tourId)

	var hotel model.Tour
	var contact model.Hotel

	err := row.Scan(
		&tour.Id,
		&tour.TourName,
		&tour.HotelId,
		&tour.EntryDate,
		&tour.DepartureDate,
		&tour.Food,
		&tour.Price,
		&tour.TourDescription,
		&hotel.Id,
		&hotel.Name,
		&hotel.LocationId,
		&hotel.Number,
		&hotel.WorkerId,
		&hotel.Description,

	)

	if err != nil {
		h.logger.Error(err.Error())
		return model.TourWithHotel{}, err
	}

	TourWithHotel := model.TourWithHotel{tour.Id, tour.TourName, hotel, tour.EntryDate, tour.DepartureDate, tour.Food, tour.Price,tour.TourDescription}

	return TourWithHotel, nil
}