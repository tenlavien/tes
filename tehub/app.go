package tehub

import (
	"log"
	"net/http"
	"os"
)

type AppConfig struct {
	ServerAddress string
	DBConfig      *DBConfig
}

func LoadAppConfig() *AppConfig {
	serverAddr, _ := os.LookupEnv("ROOT_URL")
	return &AppConfig{
		ServerAddress: serverAddr,
		DBConfig:      LoadDBConfig(),
	}
}

type App struct {
	Config *AppConfig
	Server *http.Server
	Store  *Store
}

func NewApp(config *AppConfig) *App {
	return &App{
		Config: config,
		Server: &http.Server{
			Addr: config.ServerAddress,
		},
		Store: NewStore(config.DBConfig),
	}
}

func (app *App) Start() error {
	app.InitAPIList()
	if err := app.Store.Connect(); err != nil {
		return err
	}
	log.Println("starting server ", app.Config.ServerAddress)
	if err := app.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	if err := app.Store.Disconnect(); err != nil {
		return err
	}
	return nil
}