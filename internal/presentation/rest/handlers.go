package rest

import "go-layered-architecture-sample/internal/presentation/rest/handler"

type Handlers struct {
	MedicineHandler *handler.MedicineHandler
}
