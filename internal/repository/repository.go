package repository

import (
	"useservice/internal/database"
	"useservice/internal/definitions"
)

type DatabaseAccessLayer struct {
	DB  *database.Database
	log definitions.Logger
}

func NewDatabaseAccessLayer(log definitions.Logger) (dal *DatabaseAccessLayer, err error) {
	db, err := database.New(log)
	if err != nil {
		return nil, err
	}

	dal = &DatabaseAccessLayer{
		DB:  db,
		log: log,
	}
	return dal, err

}

func (dal *DatabaseAccessLayer) CloseAll() {
	dal.DB.CloseAll()
}
