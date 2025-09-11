package database

import "useservice/internal/definitions"

type IDal interface {
	New(log definitions.Logger) (*Database, error)
}

type DatabaseAccessLayer struct {
	DB IDatabase
}

func NewDatabaseAccessLayer(log definitions.Logger) (dal *DatabaseAccessLayer, err error) {
	db, err := New(log)
	if err != nil {
		return nil, err
	}

	dal = &DatabaseAccessLayer{
		DB: db,
	}
	return dal, err

}

func (dal *DatabaseAccessLayer) CloseAll() {
	dal.DB.CloseAll()
}
