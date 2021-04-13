package models

type Expense struct {
	Id              int64 `json:"id" gorm:"primaryKey"`
	Uuid            string
	PublicationDate string
	FiscalYear      string
	AgencyCode      string
	AgencyName      string
	TotalFund       string
	CityFund        string
	Remark          string
}
