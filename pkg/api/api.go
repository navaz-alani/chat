package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/navaz-alani/chat/pkg/service"
)

type Api struct {
	cs service.ChatService
	M  *mux.Router
}

func NewApi(cs service.ChatService) (*Api, error) {
	api := new(Api)
	api.cs = cs
	api.M = mux.NewRouter()
	// configure routes for the web API
	api.M.HandleFunc("/new_ws", api.NewWsConnectionEP)

	return api, nil
}

func (api *Api) NewWsConnectionEP(w http.ResponseWriter, r *http.Request) {
	c, err := service.NewWsClient(api.cs, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// register participant with chat service
	if err := api.cs.AddParticipant(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *Api) Serve(addr string) {
	log.Printf("Attempting to serve on %s", addr)
	log.Fatalln(http.ListenAndServe(addr, api.M))
}
