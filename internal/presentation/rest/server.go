package rest

import (
	"context"
	"fmt"
	"go-layered-architecture-sample/internal/config"
	"go-layered-architecture-sample/internal/domain/entity"
	"go-layered-architecture-sample/internal/gateway/rdb"
	"go-layered-architecture-sample/internal/presentation/rest/handler"
	"go-layered-architecture-sample/internal/usecase"
	"go-layered-architecture-sample/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e       *echo.Echo
	address string
}

func NewServer(config *config.Config, logger logger.Logger) (*Server, error) {
	e := echo.New()

	e.HTTPErrorHandler = errorHandler(logger)

	primaryRepo, replicaRepo, err := rdb.NewRepository(&config.Database, logger)
	if err != nil {
		logger.Fatal("failed to rdb.NewRepository:", err)
		return nil, err
	}

	medicineRepository := rdb.NewMedicineRepository(primaryRepo, replicaRepo)
	medicineUseCase := usecase.NewMedicineUseCase(medicineRepository, primaryRepo, logger)
	medicineHandler := handler.NewMedicineHandler(medicineUseCase, logger)

	handlers := Handlers{
		MedicineHandler: medicineHandler,
	}

	setupRoutes(e, handlers)

	return &Server{
		e:       e,
		address: fmt.Sprintf("%s:%s", config.App.Host, config.HTTP.Port),
	}, nil
}

func (s *Server) Start() {
	// サーバーの起動
	go func() {
		if err := s.e.Start(s.address); err != nil && err != http.ErrServerClosed {
			fmt.Println("REST server start error:", err)
		}
	}()

	// 終了シグナルを待つためのチャネルを作成
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down the REST server...")

	// グレースフルシャットダウンのためのタイムアウト設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		fmt.Println("REST server forced to shutdown:", err)
	}
}

type errorResponse struct {
	Status  int              `json:"status"`
	Code    entity.ErrorCode `json:"code"`
	Message string           `json:"message"`
}

func errorHandler(logger logger.Logger) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		var (
			status = http.StatusInternalServerError
			code   = entity.GenericInternalError
		)

		if he, ok := err.(*echo.HTTPError); ok {
			status = he.Code
		}

		if ae, ok := err.(*entity.AppError); ok {
			status = ae.State
			code = ae.Code
		}

		message := entity.ErrorDetails[code].Message

		if err := c.JSON(status, errorResponse{
			Code:    code,
			Status:  status,
			Message: message,
		}); err != nil {
			logger.Errorf("failed to send JSON response: %v", err)
		}
	}
}
