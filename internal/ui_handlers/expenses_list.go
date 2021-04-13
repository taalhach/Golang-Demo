package ui_handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/Golang-Demo/internal/common"
	"github.com/taalhach/Golang-Demo/internal/models"
)

func ExpensesList(c echo.Context) error {
	cc := c.(*common.CustomContext)
	// create session
	session := cc.DbSession()

	//fetch expenses list from db
	var expenses []*models.Expense
	session.Find(&expenses)

	// render and return template
	return c.Render(http.StatusOK, fmt.Sprintf("%s", "expenses_list.html"), expenses)
}
