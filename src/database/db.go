package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"log"
	"vicinity-tinymesh-vas-co2/src/config"
)

const (
	dialect = "postgres"
)

func New(dbConfig *config.DBConfig, logger io.Writer) *gorm.DB {
	db, err := gorm.Open(dialect, dbConfig.String())

	if err != nil {
		log.Fatalln(err.Error())
	}

	db.SetLogger(log.New(logger, "", log.Ldate|log.Ltime))
	return db
}