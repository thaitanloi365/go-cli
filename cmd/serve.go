package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start worker",
	Long:  `Starts worker`,
	Run: func(cmd *cobra.Command, args []string) {
		var e = echo.New()

		e.Use(middleware.RequestID())
		e.Use(middleware.Recover())
		e.Use(middleware.CORS())
		e.Use(middleware.BodyLimit("20M"))

		SetupRootRoute(e.Group(""))

		var port = fmt.Sprintf(":%s", "8080")
		if err := e.Start(port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// SetupRootRoute setup root's routes
func SetupRootRoute(g *echo.Group) {
	g.GET("", healthCheckHandler)
	g.GET("/health_check", healthCheckHandler)
}

func healthCheckHandler(c echo.Context) error {
	var response = map[string]interface{}{
		"status": "ok",
	}
	return c.JSON(200, response)
}
