package controllers

import (
	"github.com/labstack/echo/v4"
	"os"
	"path"
	"strings"
)

const (
	uiPath = "./ui-dist/"
)

// Essentially e.Static()
func UI() echo.HandlerFunc {
	return func(c echo.Context) error {
		fullPath := uiPath + strings.TrimPrefix(path.Clean(c.Request().URL.String()), "/")
		if c.Request().URL.String() != "/" {
			_, err := os.Stat(fullPath)
			if err != nil {
				if !os.IsNotExist(err) {
					panic(err)
				}
				// Requested file does not exist, so we return the default
				return c.File(uiPath)
			}
		}
		return c.File(fullPath)
	}
}
