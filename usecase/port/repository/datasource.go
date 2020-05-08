package repository

import (
	"../../../domain/entity"
	"../port"
)

type DataSourceRepository interface {
	FindAll() ([]entity.DataSource, port.Error)
	Find(entity.DataSourceID) (*entity.DataSource, port.Error)
}
