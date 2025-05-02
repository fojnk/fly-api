package main

import (
	"context"
	"flyAPI/internal/repository"
	"flyAPI/internal/repository/db"
	"flyAPI/internal/service"
	"flyAPI/internal/transport"
	"flyAPI/pkg"
	"os"
	"os/signal"
	"syscall"

	_ "flyAPI/docs"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           flyAPI
// @version         1.0
// @host      localhost:8080
// @BasePath  /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("initConfig failed: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Load .env files failed: %s", err.Error())
	}

	db, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Connection to Postgres DB failed: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := transport.NewHandler(services)

	srv := new(pkg.ServerApi)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("srv run failed: %s", err.Error())
		}
	}()

	logrus.Print("flyAPI Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("flyAPI Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
