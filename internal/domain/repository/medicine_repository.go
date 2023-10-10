package repository

import (
	"context"
	"go-layered-architecture-sample/internal/domain/entity"
	"go-layered-architecture-sample/internal/gateway/rdb"
)

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_$GOFILE
type MedicineRepository interface {
	Find(ctx context.Context, params rdb.MedicineFindParams) ([]*entity.Medicine, error)
	GetByID(ctx context.Context, id int) (*entity.Medicine, error)
	Create(ctx context.Context, medicine *entity.Medicine) (*entity.Medicine, error)
	Update(ctx context.Context, medicine *entity.Medicine) (*entity.Medicine, error)
	Delete(ctx context.Context, id int) error
}
