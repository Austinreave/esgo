package main

import (
	"esgo/es"
	"esgo/models"
)

func main() {
	//初始化es
	if err := es.Init(); err != nil {
		return
	}
	defer es.Client.Stop()
	models.Create()
}