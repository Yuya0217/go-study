package rdb_test

import (
	"context"
	"go-layered-architecture-sample/internal/domain/entity"
	"go-layered-architecture-sample/internal/gateway/rdb"
	"go-layered-architecture-sample/pkg/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MedicineSuite struct {
	medicineRepository *rdb.MedicineRepository
	*Fixture
}

func (s *MedicineSuite) SetupSuite() {
	s.RDBTestSuite.SetupSuite()
	s.medicineRepository = rdb.NewMedicineRepository(s.repository, s.repository)
	s.setupTestData()

}

func (s *MedicineSuite) TearDownSuite() {
	err := s.cleanupTestData()
	if err != nil {
		s.T().Fatalf("failed to cleanup test data: %v", err)
	}
}

func TestMedicineSuite(t *testing.T) {
	testSuite := &MedicineSuite{
		Fixture: &Fixture{
			RDBTestSuite: &RDBTestSuite{},
		},
	}

	suite.Run(t, testSuite)
}

func (s *MedicineSuite) setupTestData() {
	err := s.cleanupTestData()
	if err != nil {
		s.T().Fatalf("failed to cleanup test data: %v", err)
	}

	s.insertMedicine(entity.Medicine{
		MedicineName: "バファリン",
		MedicineType: "解熱鎮痛剤",
	})

	s.insertMedicine(entity.Medicine{
		MedicineName: "ロキソプロフェン",
		MedicineType: "消炎鎮痛剤",
	})
}

func (s *MedicineSuite) cleanupTestData() error {
	_, err := s.repository.ExecContext(context.Background(), "TRUNCATE medicines")
	s.Require().NoError(err)

	return nil
}

func (suite *MedicineSuite) TestMedicineRepository_Find() {
	suite.setupTestData()

	type args struct {
		ctx    context.Context
		params rdb.MedicineFindParams
	}

	type want struct {
		result []*entity.Medicine
		err    error
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Find MedicineName",
			args: args{
				ctx: context.Background(),
				params: rdb.MedicineFindParams{
					MedicineName: utils.StringPtr("バファリン"),
				},
			},
			want: want{result: []*entity.Medicine{
				{
					ID:           1,
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
			}, err: nil},
		},
		{
			name: "Find MedicineName",
			args: args{
				ctx: context.Background(),
				params: rdb.MedicineFindParams{
					MedicineType: utils.StringPtr("消炎鎮痛剤"),
				},
			},
			want: want{result: []*entity.Medicine{
				{
					ID:           2,
					MedicineName: "ロキソプロフェン",
					MedicineType: "消炎鎮痛剤",
				},
			}, err: nil},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			got, err := suite.medicineRepository.Find(tc.args.ctx, tc.args.params)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if tc.want.result != nil {
				assert.Equal(suite.T(), tc.want.result, got)
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineRepository_GetByID() {
	suite.setupTestData()

	type args struct {
		ctx context.Context
		id  int
	}

	type want struct {
		result *entity.Medicine
		err    error
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "succesed",
			args: args{
				ctx: newContext(),
				id:  1,
			},
			want: want{result: &entity.Medicine{
				ID:           1,
				MedicineName: "バファリン",
				MedicineType: "解熱鎮痛剤",
			}, err: nil},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			got, err := suite.medicineRepository.GetByID(tc.args.ctx, tc.args.id)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if tc.want.result != nil {
				assert.Equal(suite.T(), tc.want.result, got)
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineRepository_Create() {
	type args struct {
		ctx      context.Context
		medicine *entity.Medicine
	}

	type want struct {
		result *entity.Medicine
		err    error
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "succesed",
			args: args{
				ctx: newContext(),
				medicine: &entity.Medicine{
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
			},
			want: want{result: &entity.Medicine{
				ID:           1,
				MedicineName: "バファリン",
				MedicineType: "解熱鎮痛剤",
			}, err: nil},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			got, err := suite.medicineRepository.Create(tc.args.ctx, tc.args.medicine)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if diff := cmp.Diff(got, tc.want.result, ignoreFieldsFor(entity.Medicine{}, "ID")); diff != "" {
				suite.T().Errorf("Structs differ: (-got +want)\n%s", diff)
			}
		})
	}
}
