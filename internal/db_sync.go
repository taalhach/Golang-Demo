package internal

import (
	"fmt"
	"gorm.io/gorm"
	"os"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/taalhach/Golang-Demo/internal/models"
)

const (
	dbUrl   = "https://data.cityofnewyork.us/api/views/7yay-m4ae/rows.json"
	dataKey = "data"
)

var dbSyncCmd = &cobra.Command{
	Use:   "sync_db",
	Short: "Syncs database",
	Long:  "Syncs city government database",
	Run: func(cmd *cobra.Command, args []string) {

		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v\n", err)
			os.Exit(1)
		}
		client := resty.New()
		//.SetTimeout(1 * time.Second)
		resp, err := client.R().SetHeaders(map[string]string{"Accept": "application/json"}).Get(dbUrl)
		if err != nil {
			fmt.Printf("failed to fetch data from %v error occured %v", dbUrl, err)
			os.Exit(1)
		}
		data, _, _, err := jsonparser.Get(resp.Body(), dataKey)
		if err != nil {
			fmt.Printf("failed to extract data object from response error occured %v", err)
			os.Exit(1)
		}
		session := DB.Session(&gorm.Session{})

		_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if err := sanitizeExpenseElement(value, session); err != nil {
				fmt.Printf("unable to process data: %v, skipping this row", string(value))
			}
		})

		if err != nil {
			fmt.Printf("failed to process data %v", err)
			os.Exit(1)
		}

		// now commit db session to save data
		session.Commit()

		fmt.Println("extracted data successfully")
	},
}

func sanitizeExpenseElement(data []byte, session *gorm.DB) error {
	var (
		expense models.Expense
		err     error
	)
	i := 0
	_, err = jsonparser.ArrayEach(data, func(index []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch i {
		case 1:
			expense.Uuid, err = jsonparser.ParseString(index)
		case 8:
			expense.PublicationDate, err = jsonparser.ParseString(index)
		case 9:
			expense.FiscalYear, err = jsonparser.ParseString(index)
		case 10:
			expense.AgencyCode, err = jsonparser.ParseString(index)
		case 11:
			expense.AgencyName, err = jsonparser.ParseString(index)
		case 12:
			expense.TotalFund, err = jsonparser.ParseString(index)
		case 13:
			expense.CityFund, err = jsonparser.ParseString(index)
		case 14:
			var t jsonparser.ValueType
			_, t, _, err = jsonparser.Get(index)
			if err == nil && t != jsonparser.Null {
				expense.Remark, err = jsonparser.ParseString(index)
			}

		}
		i++
		if err != nil {
			fmt.Printf("failed to extract a row: %v, skipping this row\n", err)
			return
		}
	})

	if err != nil {
		return err
	}
	session.Where(&models.Expense{Uuid: expense.Uuid}).FirstOrCreate(&expense)
	return nil
}

func init() {
	rootCommand.AddCommand(dbSyncCmd)
}
