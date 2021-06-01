package main

import (
	"github.com/Bastorx/test_lbc/logging"
	"os"
)

func main() {
	logging.Logger.SetOutput(os.Stdout)
	logging.Logger.Info("Starting LBC Test")

	r := setupRouter()
	err := r.Run(":4000")
	if err != nil {
		return 
	}
}
