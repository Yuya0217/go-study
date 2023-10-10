package dtos

type Medicine struct {
	ID           int
	MedicineName string
	MedicineType string
}

type MedicineFindParams struct {
	MedicineName *string
	MedicineType *string
}
