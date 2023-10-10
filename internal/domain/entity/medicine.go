package entity

type Medicine struct {
	ID           int    `json:"id" db:"id"`
	MedicineName string `json:"medicineName" db:"medicine_name"`
	MedicineType string `json:"medicineType" db:"medicine_type"`
}
