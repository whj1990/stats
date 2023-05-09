package repo

import (
	"github.com/google/wire"
	"github.com/whj1990/go-core/repository"
	"github.com/whj1990/go-core/store"
	"github.com/whj1990/go-core/trace"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(trace.NewGormLogger, store.NewDB, NewReviewProjectRepo)

func NewReviewProjectRepo(db *gorm.DB) ReviewProjectRepo {
	return &reviewProjectRepo{
		BaseRepo: repository.BaseRepo{Db: db, Model: ReviewProjectData{}},
	}
}
