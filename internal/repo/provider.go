package repo

import (
	"github.com/google/wire"
	"github.com/whj1990/go-core/repository"
	"github.com/whj1990/go-core/store"
	"github.com/whj1990/go-core/trace"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(trace.NewGormLogger, store.NewDB, NewOrderRepo)

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{
		BaseRepo: repository.BaseRepo{Db: db, Model: OrderData{}},
	}
}
