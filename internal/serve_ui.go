package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taalhach/Golang-Demo/internal/common"
	"github.com/taalhach/Golang-Demo/internal/ui_handlers"
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

const port = 8081

var serveUi = &cobra.Command{
	Use:   "serve_ui",
	Short: "servers ui api",
	Long:  fmt.Sprintf("servers velocity works ui api on localhost port %v", port),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v", err)
			os.Exit(1)
		}

		e := echo.New()
		pattern := fmt.Sprintf("%s/expenses_list.html", MainConfigs.TemplatesDirectory)
		renderer := &TemplateRenderer{
			templates: template.Must(template.ParseGlob(pattern)),
		}

		// now attach template renderer
		e.Renderer = renderer

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("method=${method} uri=${uri} status=${status} time=${latency_human}\n"),
		}))

		e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := &common.CustomContext{
					Context:     c,
					DB:          DB,
					MainConfigs: MainConfigs,
				}
				return h(cc)
			}
		})

		e.GET("/", ui_handlers.RootHandler)
		e.GET("/expenses_list", ui_handlers.ExpensesList)

		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	},
}

func init() {
	rootCommand.AddCommand(serveUi)
}
