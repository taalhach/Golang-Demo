package common

import (
	"github.com/labstack/echo/v4"
	"github.com/taalhach/Golang-Demo/internal/configs"
	"gorm.io/gorm"
)

type CustomContext struct {
	echo.Context
	DB          *gorm.DB
	MainConfigs *configs.MainConfig
}
