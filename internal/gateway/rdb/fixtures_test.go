package rdb_test

import "go-layered-architecture-sample/internal/domain/entity"

type Fixture struct {
	*RDBTestSuite
}

func (f *Fixture) insertMedicine(row entity.Medicine) int64 {
	ctx := newContext()

	result, err := f.repository.NamedExecContext(ctx,
		`INSERT INTO medicines (
			medicine_name,
			medicine_type
		) VALUES (
			:medicine_name,
			:medicine_type
		)`, row)

	f.Require().NoError(err)

	id, err := result.LastInsertId()

	f.Require().NoError(err)

	return id
}
