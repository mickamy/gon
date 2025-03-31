package gon

import (
	"fmt"

	"github.com/mickamy/gon/internal/config"
)

var Config *config.Config

func init() {
	var err error
	Config, err = config.LoadConfig()
	if err != nil {
		fmt.Println("⚠️ failed to load config file:", err)
	}
}
