package hotel

import (
	"context"
	"database/sql"
	errcode "github.com/niumandzi/nto2022/internal/errors"
	"github.com/niumandzi/nto2022/model"
	"github.com/niumandzi/nto2022/pkg/logging"
)

type HotelRepository struct {
	db     *sql.DB
	logger logging.Logger
}

func NewHotelRepository(db *sql.DB, logger logging.Logger) HotelRepository {
	return HotelRepository{
		db:     db,
		logger: logger,
	}
}

func (h HotelRepository) Create(ctx context.Context, hotel model.Hotel) (int, error) {
	result, err := h.db.ExecContext(
		ctx,
		"INSERT INTO hotel (name, location_id, number, worker_id, description) VALUES (?, ?, ?, ?, ?)",
		hotel.Name,
		hotel.LocationId,
		hotel.Number,
		hotel.WorkerId,
		hotel.Description,
	)
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}

	return int(id), nil
}

func (h HotelRepository) GetById(ctx context.Context, hotelId int) (model.HotelWithContact, error) {
	row := h.db.QueryRowContext(ctx, "SELECT * FROM hotel LEFT JOIN contact ON hotel.worker_id = contact.id WHERE hotel.id = ?", hotelId)

	var hotel model.Hotel
	var contact model.Contact

	err := row.Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.LocationId,
		&hotel.Number,
		&hotel.WorkerId,
		&hotel.Description,
		&contact.Id,
		&contact.ContactType,
		&contact.Name,
		&contact.Number,
		&contact.Email,
	)

	if err != nil {
		h.logger.Error(err.Error())
		return model.HotelWithContact{}, err
	}

	hotelWithContact := model.HotelWithContact{hotel.Id, hotel.Name, hotel.LocationId, hotel.Number, contact, hotel.Description}

	return hotelWithContact, nil
}

func (h HotelRepository) GetAll(ctx context.Context) ([]model.HotelWithContact, error) {
	rows, err := h.db.QueryContext(ctx, "SELECT * FROM hotel LEFT JOIN contact ON hotel.worker_id = contact.id")

	if err != nil {
		h.logger.Error(err.Error())
		return []model.HotelWithContact{}, err
	}
	defer rows.Close()

	var hotels []model.HotelWithContact

	for rows.Next() {
		var hotel model.Hotel
		var contact model.Contact
		var hotelWithContact model.HotelWithContact

		err = rows.Scan(&hotel.Id,
			&hotel.Name,
			&hotel.LocationId,
			&hotel.Number,
			&hotel.WorkerId,
			&hotel.Description,
			&contact.Id,
			&contact.ContactType,
			&contact.Name,
			&contact.Number,
			&contact.Email,
		)

		if err != nil {
			h.logger.Error(err.Error())
			return hotels, err
		}

		hotelWithContact = model.HotelWithContact{hotel.Id, hotel.Name, hotel.LocationId, hotel.Number, contact, hotel.Description}

		hotels = append(hotels, hotelWithContact)
	}

	return hotels, nil
}

func (h HotelRepository) Update(ctx context.Context, hotel model.Hotel) error {
	res, err := h.db.ExecContext(ctx, `UPDATE hotel SET name = ?, location_id = ?, number = ?, worker_id = ?, description = ? WHERE id = ?`, hotel.Name, hotel.LocationId, hotel.Number, hotel.WorkerId, hotel.Description, hotel.Id)
	if err != nil {
		h.logger.Error(err.Error())
	}

	count, err := res.RowsAffected()
	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	if count != 1 {
		resErr := errcode.Wrap("update hotel", "row counter not equals 1")
		h.logger.Error(resErr.Error())
		return err
	}

	return nil
}

func (h HotelRepository) Delete(ctx context.Context, hotelId int) error {
	res, err := h.db.ExecContext(ctx, `DELETE FROM hotel WHERE id = ?`, hotelId)
	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	if count != 1 {
		resErr := errcode.Wrap("delete hotel", "row counter not equals 1")
		h.logger.Error(resErr.Error())
		return err
	}

	return nil
}
