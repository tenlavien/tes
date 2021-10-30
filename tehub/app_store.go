package tehub

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Schema   string `json:"schema"`
}

func LoadDBConfig() *DBConfig {
	username, _ := os.LookupEnv("DB_USER_NAME")
	password, _ := os.LookupEnv("DB_PASSWORD")
	server, _ := os.LookupEnv("DB_SERVER")
	schema, _ := os.LookupEnv("DB_SCHEMA")

	return &DBConfig{
		Username: username,
		Password: password,
		Server:   server,
		Schema:   schema,
	}
}

type Store struct {
	Config *DBConfig
	DB     *gorm.DB
}

func NewStore(config *DBConfig) *Store {
	return &Store{
		Config: config,
	}
}

func (s *Store) Connect() error {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.Config.Username,
		s.Config.Password,
		s.Config.Server,
		s.Config.Schema,
	)
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		return err
	}
	s.DB = db
	log.Println("[info] connect db success")
	return nil
}

func (s *Store) Disconnect() error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func Paginate(pageID, perPage int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if pageID == 0 {
			pageID = 1
		}

		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (pageID - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}