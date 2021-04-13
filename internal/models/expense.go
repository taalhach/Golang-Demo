package models

type Expense struct {
	Id              int64  `json:"id" gorm:"primaryKey"`
	Uuid            string `json:"-"`
	PublicationDate string `json:"publication_date"`
	FiscalYear      string `json:"fiscal_year"`
	AgencyCode      string `json:"agency_code"`
	AgencyName      string `json:"agency_name"`
	TotalFund       string `json:"total_fund"`
	CityFund        string `json:"city_fund"`
	Remark          string `json:"remark"`
}
