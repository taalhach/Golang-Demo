package internal

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/taalhach/Golang-Demo/internal/models"
	"gorm.io/gorm"
	"os"
)

const dbUrl = "https://data.cityofnewyork.us/api/views/7yay-m4ae/rows.json"

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
		resp, _ := client.R().SetHeaders(map[string]string{"Accept": "application/json"}).Get(dbUrl)
		//fmt.Println(string(resp.Body()), err)
		data, _, _, err := jsonparser.Get(resp.Body(), "data")
		fmt.Println(err)
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
	},
}

func sanitizeExpenseElement(data []byte, session *gorm.DB) error {
	var (
		expense models.Expense
		err     error
	)
	i := 0
	_, err = jsonparser.ArrayEach(data, func(currentData []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(i, string(currentData))
		switch i {
		case 8:
			//data, _, _, _ := jsonparser.Get(currentData)
			//fmt.Println("whats that ", string(data))
			_, t, _, err := jsonparser.Get(currentData)
			fmt.Println("data ", err, t)
			fmt.Print(jsonparser.GetUnsafeString(currentData))
			fmt.Printf(" date %v \n", i)
		case 9:
			_, err = jsonparser.GetInt(currentData)
			fmt.Printf(" year %v \n", i)
		case 10:
			expense.AgencyCode, err = jsonparser.GetUnsafeString(currentData)
		case 11:
			fmt.Println(jsonparser.Get(currentData))
			expense.AgencyName, err = jsonparser.GetUnsafeString(currentData)
		case 12:
			expense.TotalFund, err = jsonparser.GetFloat(currentData)
		case 13:
			expense.CityFund, err = jsonparser.GetFloat(currentData)
		case 14:
			var t jsonparser.ValueType
			_, t, _, err = jsonparser.Get(currentData)
			if t != jsonparser.Null {
				expense.Remark, err = jsonparser.GetString(currentData)
			}

		}
		i++
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	if err != nil {
		return err
	}

	session.Create(&expense)
	return nil
}

func init() {
	rootCommand.AddCommand(dbSyncCmd)
}
