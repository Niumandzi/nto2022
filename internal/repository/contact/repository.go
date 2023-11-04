package contact

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

func NewContactRepository(db *sql.DB, logger logging.Logger) ContactRepository {
	return ContactRepository{
		db:     db,
		logger: logger,
	}
}

func (s ContactRepository) Create(ctx context.Context, contact model.Contact) (int, error) {
	result, err := s.db.ExecContext(
		ctx,
		"INSERT INTO contact (contact_type, Name, Number, Email) VALUES (?, ?, ?, ?)",
		contact.ContactType,
		contact.Name,
		contact.Number,
		contact.Email,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}

	return int(id), nil
}

func (s ContactRepository) GetById(ctx context.Context, contactId int) (model.Contact, error) {
	row := s.db.QueryRowContext(ctx, "SELECT * FROM contact WHERE id = ?", contactId)

	var contact model.Contact

	err := row.Scan(
		&contact.Id,
		&contact.ContactType,
		&contact.Name,
		&contact.Number,
		&contact.Email,
	)

	if err != nil {
		s.logger.Error(err.Error())
		return model.Contact{}, err
	}

	return contact, nil
}

func (s ContactRepository) GetByType(ctx context.Context, contactType string) ([]model.Contact, error) {
	var query string
	var args []interface{}

	if contactType == "all" {
		query = "SELECT * FROM contact"
	} else {
		query = "SELECT * FROM contact WHERE contact_type = ?"
		args = append(args, contactType)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var contacts []model.Contact
	for rows.Next() {
		var contact model.Contact
		err := rows.Scan(&contact.Id, &contact.ContactType, &contact.Name, &contact.Number, &contact.Email)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return contacts, nil
}

func (s ContactRepository) Update(ctx context.Context, contactInput model.Contact) error {
	res, err := s.db.ExecContext(
		ctx,
		"UPDATE contact SET contact_type=?, Name=?, Number=?, Email=? WHERE Id=?",
		contactInput.ContactType,
		contactInput.Name,
		contactInput.Number,
		contactInput.Email,
		contactInput.Id,
	)

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	if count != 1 {
		resErr := errcode.Wrap("update contact", "row counter not equals 1")
		s.logger.Error(resErr.Error())
		return err
	}

	return nil
}

func (s ContactRepository) Delete(ctx context.Context, contactId int) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM contact WHERE id = ?`, contactId)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	if count != 1 {
		resErr := errcode.Wrap("delete contact", "row counter not equals 1")
		s.logger.Error(resErr.Error())
		return err
	}

	return nil
}
