package rdb

import (
	"context"
	"fmt"
	"go-layered-architecture-sample/internal/domain/entity"

	sq "github.com/Masterminds/squirrel"
)

type MedicineRepository struct {
	primary *Repository
	replica *Repository
}

func NewMedicineRepository(primary *Repository, replica *Repository) *MedicineRepository {
	return &MedicineRepository{
		primary: primary,
		replica: replica,
	}
}

type MedicineFindParams struct {
	MedicineName *string
	MedicineType *string
}

func (r *MedicineRepository) Find(ctx context.Context, params MedicineFindParams) ([]*entity.Medicine, error) {
	sb := NewBaseStatementBuilder()

	builder := sb.
		Select("id, medicine_name, medicine_type").
		From("medicines")

	if params.MedicineName != nil {
		builder = builder.Where(sq.Eq{"medicine_name": params.MedicineName})
	}

	if params.MedicineType != nil {
		builder = builder.Where(sq.Eq{"medicine_type": params.MedicineType})
	}

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.replica.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entities := make([]*entity.Medicine, 0)

	for rows.Next() {
		e := entity.Medicine{}

		err = rows.Scan(&e.ID, &e.MedicineName, &e.MedicineType)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, &e)
	}

	return entities, nil
}

func (r *MedicineRepository) GetByID(ctx context.Context, id int) (*entity.Medicine, error) {
	sb := NewBaseStatementBuilder()

	query, args, err := sb.
		Select("id, medicine_name, medicine_type").
		From("medicines").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.replica.QueryRowxContext(ctx, query, args...)

	e := &entity.Medicine{}

	err = row.Scan(&e.ID, &e.MedicineName, &e.MedicineType)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *MedicineRepository) Create(ctx context.Context, medicine *entity.Medicine) (*entity.Medicine, error) {
	sb := NewBaseStatementBuilder()

	query, args, err := sb.
		Insert("medicines").
		Columns("medicine_name", "medicine_type").
		Values(medicine.MedicineName, medicine.MedicineType).
		ToSql()

	if err != nil {
		return nil, err
	}

	result, err := r.primary.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, int(lastInsertId))
}

func (r *MedicineRepository) Update(ctx context.Context, medicine *entity.Medicine) (*entity.Medicine, error) {
	sb := NewBaseStatementBuilder()

	query, args, err := sb.
		Update("medicines").
		Set("medicine_name", medicine.MedicineName).
		Set("medicine_type", medicine.MedicineType).
		Where(sq.Eq{"id": medicine.ID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	result, err := r.primary.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetByID(ctx, int(lastInsertId))
}

func (r *MedicineRepository) Delete(ctx context.Context, id int) error {
	sb := NewBaseStatementBuilder()

	query, args, err := sb.
		Delete("medicines").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.primary.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}
