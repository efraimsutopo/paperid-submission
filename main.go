package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/efraimsutopo/paperid-submission/config"
	"github.com/efraimsutopo/paperid-submission/controller"
	"github.com/efraimsutopo/paperid-submission/model"
	"github.com/efraimsutopo/paperid-submission/modules/customValidator"
	"github.com/efraimsutopo/paperid-submission/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	userController "github.com/efraimsutopo/paperid-submission/controller/user"
	userRepository "github.com/efraimsutopo/paperid-submission/repository/user"
	userService "github.com/efraimsutopo/paperid-submission/service/user"

	transactionController "github.com/efraimsutopo/paperid-submission/controller/transaction"
	transactionRepository "github.com/efraimsutopo/paperid-submission/repository/transaction"
	transactionService "github.com/efraimsutopo/paperid-submission/service/transaction"
)

func main() {
	// Init Logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Init config
	if err := config.New(); err != nil {
		logger.Error(err)
		return
	}

	errs := make(chan error)

	// Init Syscall Notifier
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Init Database
	var db *gorm.DB
	{
		var err error
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Get().DBUser,
			config.Get().DBPassword,
			config.Get().DBHost,
			config.Get().DBPort,
			config.Get().DBName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Error(err)
			return
		}

		// Migrate the schema
		db.AutoMigrate(&model.User{})
		db.AutoMigrate(&model.Session{})
		db.AutoMigrate(&model.Transaction{})
	}

	// Init Echo
	ec := echo.New()
	ec.Validator = customValidator.New()

	// Initialize Repository
	userRepo := userRepository.New(db)
	transactionRepo := transactionRepository.New(db)

	// Initialize Service
	userSvc := userService.New(userRepo)
	transactionSvc := transactionService.New(transactionRepo)

	// Initialize Controller
	userCtrl := userController.New(userSvc)
	transactionCtrl := transactionController.New(transactionSvc)
	ctrl := controller.New(userCtrl, transactionCtrl)

	// Register Routes
	r := routes.New(ctrl)
	r.RegisterRoutes(ec)

	// Start HTTP Server
	go func() {
		errs <- ec.Start(":" + config.Get().AppPort)
	}()

	logger.Errorf("Service Ended with err: %v", <-errs)
}
