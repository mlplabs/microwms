package main

import (
	"github.com/mlplabs/microwms/internal/app"
	"github.com/mlplabs/microwms/internal/config"
)

// @title WMS API
// @version 1.0
func main() {
	cfg := config.ReadEnv()
	app.Init(cfg)
}
