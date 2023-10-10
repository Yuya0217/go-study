package usecase_test

import (
	"context"
	"go-layered-architecture-sample/internal/domain/entity"
	mock_repository "go-layered-architecture-sample/internal/domain/repository/mocks"
	"go-layered-architecture-sample/internal/gateway/rdb"
	"go-layered-architecture-sample/internal/usecase"
	"go-layered-architecture-sample/internal/usecase/dtos"
	"go-layered-architecture-sample/pkg/utils"
	"testing"

	mock_logger "go-layered-architecture-sample/pkg/logger/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mocks struct {
	mockMedicineRepository  *mock_repository.MockMedicineRepository
	mockTransactionExecutor *mock_repository.MockTransactionExecutor
	mockLogger              *mock_logger.MockLogger
}

type MedicineSuite struct {
	suite.Suite
	mocks           mocks
	mockCtrl        *gomock.Controller
	medicineUseCase usecase.MedicineUseCase
}

func (suite *MedicineSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mocks.mockMedicineRepository = mock_repository.NewMockMedicineRepository(suite.mockCtrl)
	suite.mocks.mockTransactionExecutor = mock_repository.NewMockTransactionExecutor(suite.mockCtrl)
	suite.mocks.mockLogger = mock_logger.NewMockLogger(suite.mockCtrl)
	suite.medicineUseCase = usecase.NewMedicineUseCase(suite.mocks.mockMedicineRepository, suite.mocks.mockTransactionExecutor, suite.mocks.mockLogger)
}

func (suite *MedicineSuite) TearDownSuite() {
	suite.mockCtrl.Finish()
}

func TestMedicineSuite(t *testing.T) {
	testSuite := &MedicineSuite{}
	suite.Run(t, testSuite)
}

func (suite *MedicineSuite) TestMedicineUseCase_Find() {
	type args struct {
		ctx    context.Context
		params dtos.MedicineFindParams
	}

	type want struct {
		result []*dtos.Medicine
		err    error
	}

	testCases := []struct {
		name         string
		args         args
		want         want
		mockBehavior func()
	}{
		{
			name: "succesed",
			args: args{
				ctx: context.Background(),
				params: dtos.MedicineFindParams{
					MedicineName: utils.StringPtr("バファリン"),
				},
			},
			want: want{result: []*dtos.Medicine{
				{
					ID:           1,
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
			}, err: nil},
			mockBehavior: func() {
				suite.mocks.mockMedicineRepository.EXPECT().Find(gomock.Any(), rdb.MedicineFindParams{
					MedicineName: utils.StringPtr("バファリン"),
				}).Return([]*entity.Medicine{
					{
						ID:           1,
						MedicineName: "バファリン",
						MedicineType: "解熱鎮痛剤",
					},
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockBehavior()

			got, err := suite.medicineUseCase.Find(tc.args.ctx, tc.args.params)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if tc.want.result != nil {
				assert.Equal(suite.T(), tc.want.result, got)
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineUseCase_GetByID() {
	type args struct {
		ctx context.Context
		id  int
	}

	type want struct {
		result *dtos.Medicine
		err    error
	}

	testCases := []struct {
		name         string
		args         args
		want         want
		mockBehavior func()
	}{
		{
			name: "succesed",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: want{
				result: &dtos.Medicine{
					ID:           1,
					MedicineName: "fugafuga",
					MedicineType: "piyopiyo",
				},
				err: nil,
			},
			mockBehavior: func() {
				suite.mocks.mockMedicineRepository.EXPECT().GetByID(gomock.Any(), 1).Return(
					&entity.Medicine{
						ID:           1,
						MedicineName: "fugafuga",
						MedicineType: "piyopiyo",
					},
					nil,
				)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockBehavior()

			got, err := suite.medicineUseCase.GetByID(context.Background(), tc.args.id)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if tc.want.result != nil {
				assert.Equal(suite.T(), tc.want.result, got)
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineUseCase_Create() {
	type args struct {
		ctx   context.Context
		input dtos.Medicine
	}

	type want struct {
		result *dtos.Medicine
		err    error
	}

	testCases := []struct {
		name         string
		args         args
		want         want
		mockBehavior func()
	}{
		{
			name: "succesed",
			args: args{
				ctx: context.Background(),
				input: dtos.Medicine{
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
			},
			want: want{
				result: &dtos.Medicine{
					ID:           1,
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
				err: nil,
			},
			mockBehavior: func() {
				suite.mocks.mockMedicineRepository.EXPECT().Create(gomock.Any(), &entity.Medicine{
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				}).Return(
					&entity.Medicine{
						ID:           1,
						MedicineName: "バファリン",
						MedicineType: "解熱鎮痛剤",
					},
					nil,
				)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.mockBehavior()

			got, err := suite.medicineUseCase.Create(context.Background(), tc.args.input)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			if tc.want.result != nil {
				assert.Equal(suite.T(), tc.want.result, got)
			}
		})
	}
}
