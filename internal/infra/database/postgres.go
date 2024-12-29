package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/config"
)

const maxRetryCount = 20

type GORMImpl struct {
	c *gorm.DB
}

func NewPGGORMImpl(config config.Config) (GORMImpl, error) {
	var err error
	var retryCount int

	gormimpl := GORMImpl{}

	DBURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Name,
		"disable",
		config.Database.Password,
	)

	for retryCount < maxRetryCount {
		gormimpl.c, err = gorm.Open("postgres", DBURL)

		if err != nil {
			retryCount++
			log.Printf("error connecting to database, retrying... %s", err)
		} else {
			break
		}

		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal(err)

		return gormimpl, err
	}

	fatalErrorAboutDB := fmt.Errorf("can't start the database %s", err)

	log.Fatal(fatalErrorAboutDB)

	return gormimpl, fatalErrorAboutDB
}

func (pg GORMImpl) Client() *gorm.DB {
	return pg.c
}

func (pg GORMImpl) Close() error {
	return pg.Client().Close()
}
