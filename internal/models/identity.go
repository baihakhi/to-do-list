package models

// Global identity struct
type Identity struct {
	ID   uint64 `json:"id"`
	Code string `json:"code" query:"code" form:"code"`
	Name string `json:"Name" query:"name" form:"name"`
}
