package main

import (
	"context"
	"github.com/niumandzi/nto2022/internal/app"
	"github.com/niumandzi/nto2022/internal/config"
	contactRepository "github.com/niumandzi/nto2022/internal/repository/contact"
	hotelRepository "github.com/niumandzi/nto2022/internal/repository/hotel"
	"github.com/niumandzi/nto2022/internal/usecase"
	contactUsecase "github.com/niumandzi/nto2022/internal/usecase/contact"
	hotelUsecase "github.com/niumandzi/nto2022/internal/usecase/hotel"
	"github.com/niumandzi/nto2022/pkg/logging"
	"github.com/niumandzi/nto2022/pkg/sqlitedb"
	"time"
)

func main() {
	ctx := context.Background()

	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	cfg, err := config.InitConfig()

	db, err := sqlitedb.NewClient(cfg.DriverName, cfg.FilePath)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = sqlitedb.CreateTables(db)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	timeoutContext := time.Duration(cfg.Timeout) * time.Second

	contactRepo := contactRepository.NewContactRepository(db, logger)
	hotelRepo := hotelRepository.NewHotelRepository(db, logger)

	contact := contactUsecase.NewContacUsecase(contactRepo, timeoutContext, logger)
	hotel := hotelUsecase.NewHotelUsecase(hotelRepo, timeoutContext, logger)

	cases := usecase.NewUsecases(contact, hotel)
	println(ctx, cases)

	gui := app.NewGUI()
	gui.Run()
}
