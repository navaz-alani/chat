package main

import (
	"github.com/navaz-alani/chat/pkg/api"
	"github.com/navaz-alani/chat/pkg/service"
)

func main() {
	cs, _ := service.NewService()
	cs.Serve()
	api, _ := api.NewApi(cs)
	api.Serve(":5000")
}
