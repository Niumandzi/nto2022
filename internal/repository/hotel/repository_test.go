package hotel

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/niumandzi/nto2022/model"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestHotelRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logging.Init()
	logger := logging.GetLogger()

	r := NewHotelRepository(db, logger)

	type args struct {
		ctx   context.Context
		hotel model.Hotel
	}

	type mockBehaviour func(args args)

	testTable := []struct {
		name          string
		want          int
		mockBehaviour mockBehaviour
		args          args
		wantErr       bool
	}{
		{
			name: "OK",
			want: 1,
			mockBehaviour: func(args args) {
				mock.ExpectExec("INSERT INTO hotel").
					WithArgs(args.hotel.Name, args.hotel.LocationId, args.hotel.Number, args.hotel.WorkerId,
						args.hotel.Description).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: args{
				context.Background(),
				model.Hotel{Name: "test", LocationId: 1, Number: "number", WorkerId: 1, Description: "description"},
			},
			wantErr: false,
		},
		{
			name:    "Some fields are emtpy",
			wantErr: true,
			mockBehaviour: func(args args) {
				mock.ExpectExec("INSERT INTO hotel").WithArgs(args.hotel.Name, args.hotel.LocationId, args.hotel.Number, args.hotel.WorkerId,
					args.hotel.Description).WillReturnError(errors.New("insert error"))
			},
			args: args{
				context.Background(),
				model.Hotel{Name: "", LocationId: 0, Number: "", Description: ""},
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)

			got, err := r.Create(tt.args.ctx, tt.args.hotel)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
