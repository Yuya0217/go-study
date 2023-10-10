package handler_test

import (
	openapi "go-layered-architecture-sample/generated"
	"go-layered-architecture-sample/internal/presentation/rest/handler"
	"go-layered-architecture-sample/internal/usecase/dtos"
	mock_usecase "go-layered-architecture-sample/internal/usecase/mocks"
	mock_logger "go-layered-architecture-sample/pkg/logger/mocks"
	"go-layered-architecture-sample/pkg/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mocks struct {
	mockMedicineUseCase *mock_usecase.MockMedicineUseCase
	mockLogger          *mock_logger.MockLogger
}

type MedicineSuite struct {
	suite.Suite
	mocks           mocks
	mockCtrl        *gomock.Controller
	medicineHandler *handler.MedicineHandler
}

func (suite *MedicineSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mocks.mockMedicineUseCase = mock_usecase.NewMockMedicineUseCase(suite.mockCtrl)
	suite.mocks.mockLogger = mock_logger.NewMockLogger(suite.mockCtrl)
	suite.medicineHandler = handler.NewMedicineHandler(suite.mocks.mockMedicineUseCase, suite.mocks.mockLogger)
}

func (suite *MedicineSuite) TearDownSuite() {
	suite.mockCtrl.Finish()
}

func TestMedicineSuite(t *testing.T) {
	testSuite := &MedicineSuite{}
	suite.Run(t, testSuite)
}

func (suite *MedicineSuite) TestMedicineHandler_Find() {
	type want struct {
		err error
	}

	testCases := []struct {
		name         string
		want         want
		mockBehavior func()
	}{
		{
			name: "succesed",
			want: want{err: nil},
			mockBehavior: func() {
				suite.mocks.mockMedicineUseCase.EXPECT().Find(gomock.Any(), dtos.MedicineFindParams{
					MedicineName: utils.StringPtr("バファリン"),
				}).Return([]*dtos.Medicine{
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

			ctx, buf, err := makeMockContext(
				echo.New(),
				echo.GET,
				"/v1/medicines",
				nil,
				nil,
				map[string]string{
					"medicineName": "バファリン",
				},
			)

			if err != nil {
				suite.T().Fatalf("Failed to makeMockContext: %s", err)
			}

			err = suite.medicineHandler.Find(ctx)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			goldenFileName := makeGoldenFileName(suite.T(), "medicine")

			if *updateGolden {
				writeGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			} else {
				assertGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineHandler_GetByID() {
	type want struct {
		err error
	}

	testCases := []struct {
		name         string
		want         want
		mockBehavior func()
	}{
		{
			name: "succesed",
			want: want{err: nil},
			mockBehavior: func() {
				suite.mocks.mockMedicineUseCase.EXPECT().GetByID(gomock.Any(), 1).Return(
					&dtos.Medicine{
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

			ctx, buf, err := makeMockContext(
				echo.New(),
				echo.GET,
				"/v1/medicines/:id",
				nil,
				map[string]string{
					"id": "1",
				},
				nil,
			)

			if err != nil {
				suite.T().Fatalf("Failed to makeMockContext: %s", err)
			}

			err = suite.medicineHandler.GetByID(ctx)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			goldenFileName := makeGoldenFileName(suite.T(), "medicine")

			if *updateGolden {
				writeGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			} else {
				assertGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			}
		})
	}
}

func (suite *MedicineSuite) TestMedicineHandler_Create() {
	type args struct {
		e   *echo.Echo
		req openapi.MedicineRequest
	}

	type want struct {
		err error
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
				e: echo.New(),
				req: openapi.MedicineRequest{
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				},
			},
			want: want{err: nil},
			mockBehavior: func() {
				suite.mocks.mockMedicineUseCase.EXPECT().Create(gomock.Any(), dtos.Medicine{
					MedicineName: "バファリン",
					MedicineType: "解熱鎮痛剤",
				}).Return(
					&dtos.Medicine{
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

			ctx, buf, err := makeMockContext(
				tc.args.e,
				echo.POST,
				"/v1/medicines",
				tc.args.req,
				nil,
				nil,
			)

			if err != nil {
				suite.T().Fatalf("Failed to makeMockContext: %s", err)
			}

			err = suite.medicineHandler.Create(ctx)
			assert.Equal(suite.T(), tc.want.err, err, "Error mismatch")

			goldenFileName := makeGoldenFileName(suite.T(), "medicine")

			if *updateGolden {
				writeGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			} else {
				assertGoldenFile(suite.T(), goldenFileName, buf.Bytes())
			}
		})
	}
}
