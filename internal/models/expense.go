package models

type Expense struct {
	Id              int64 `json:"id" gorm:"primaryKey"`
	PublicationDate string
	FiscalYear      string
	AgencyCode      string
	AgencyName      string
	TotalFund       float64
	CityFund        float64
	Remark          string
}
