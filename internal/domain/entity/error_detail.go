package entity

import (
	"net/http"
)

type ErrorCode int

const (
	// Medicine関連エラーコード
	MedicineValidationFailed ErrorCode = 10001
	MedicineNotFound         ErrorCode = 10002
	MedicineFindFailed       ErrorCode = 10003
	MedicineCreateFailed     ErrorCode = 10004
	MedicineUpdateFailed     ErrorCode = 10005
	MedicineDeleteFailed     ErrorCode = 10006

	// Authentication関連エラーコード
	AuthInvalidCredentials ErrorCode = 11001
	AuthExpiredToken       ErrorCode = 11002
	AuthPermissionDenied   ErrorCode = 11003

	GenericInternalError ErrorCode = 90001
	GenericBadRequest    ErrorCode = 90002
)

type ErrorDetail struct {
	State   int
	Message string
}

var ErrorDetails = map[ErrorCode]ErrorDetail{
	MedicineValidationFailed: {
		State:   http.StatusBadRequest,
		Message: "薬情報の入力値が不正です。",
	},
	MedicineNotFound: {
		State:   http.StatusNotFound,
		Message: "薬情報が見つかりませんでした。",
	},
	MedicineFindFailed: {
		State:   http.StatusInternalServerError,
		Message: "薬情報の検索に失敗しました。",
	},
	MedicineCreateFailed: {
		State:   http.StatusInternalServerError,
		Message: "薬情報の登録に失敗しました。",
	},
	MedicineUpdateFailed: {
		State:   http.StatusInternalServerError,
		Message: "薬情報の更新に失敗しました。",
	},
	MedicineDeleteFailed: {
		State:   http.StatusInternalServerError,
		Message: "薬情報の削除に失敗しました。",
	},
	GenericInternalError: {
		State:   http.StatusInternalServerError,
		Message: "予期せぬエラーが発生しました。",
	},
}
