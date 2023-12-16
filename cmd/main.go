package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paper.id.disbursement/config/database"
	"github.com/paper.id.disbursement/constants"
	"github.com/paper.id.disbursement/models"
	"github.com/paper.id.disbursement/routers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	db *gorm.DB
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	if err := godotenv.Load(constants.ENVPATH); err != nil {
		logrus.Error(err.Error())
		panic(err.Error())
	}

	appMode := os.Getenv("APP_MODE")

	if appMode == constants.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
}

func createInitialWallet(name string, balance float64) error {
	var existingUser models.Wallet

	// check if the user with the given name already exists
	result := db.Where("name = ?", name).First(&existingUser)
	if result.RowsAffected > 0 {
		logrus.Info("User with name " + name + " already exists")
		return nil
	}

	wallet := &models.Wallet{
		Name:    name,
		Balance: balance,
	}

	result = db.Create(wallet)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func main() {
	var (
		portBuilder strings.Builder
		err         error
	)

	logrus.Info("starting application")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err = database.NewDBConnection()
	if err != nil {
		logrus.Error("can not connected database. reason : ", err.Error())
		return
	}

	route := gin.Default()
	route.NoRoute(routers.NoRoute)
	routers.RouterConfig(route)
	routers.API(&route.RouterGroup, db)

	portBuilder.WriteString(":")
	portBuilder.WriteString(os.Getenv("APP_PORT"))

	srv := &http.Server{
		Addr:         portBuilder.String(),
		Handler:      route,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// create sample wallet with initial balance
	err = createInitialWallet("wallet", 5000.0)
	if err != nil {
		logrus.Error(err.Error())
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Error("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()

	logrus.Warn("shutting down gracefully, press Ctrl+C again to force 🔴")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown 🔴: ", err)
	}

	logrus.Warn("application closed 🔴")
}