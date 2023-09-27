package api

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-server/data"
	"rest-server/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type APIConfig struct {
	URL  string
	Port string
	Cors []string
}

type API struct {
	service *service.Service
	config  APIConfig
	r       *chi.Mux
}

func NewAPI(config APIConfig, s *service.Service) *API {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if len(config.Cors) > 0 {
		c := cors.New(cors.Options{
			AllowedOrigins:   config.Cors,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token", "Remote-Token"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		r.Use(c.Handler)
	}

	api := API{
		service: s,
		config: config,
		r: r,
	}

	api.initRoutes(r)

	return &api
}

func (api *API) StartServer() {
	log.Println("Starting webserver at port", api.config.Port)
	err := http.ListenAndServe(api.config.Port, api.r)
	if err != nil {
		log.Println(err.Error())
	}
}
 
func (api *API) GetWords(w http.ResponseWriter, r *http.Request) {
	words, err := api.service.Words.GetAll()

	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		log.Printf("[ERROR] API.GetWords: %s", err)
		return
	}

	jsonResponse(w, words)
}

func (api *API) AddWord(w http.ResponseWriter, r *http.Request) {
	upd := data.WordProperties{}
	err := parseFrom(w, r, &upd)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	id, err := api.service.Words.Add(upd)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		log.Printf("[ERROR] API.AddWord: %s", err)
		return
	}

	response := Response{
		ID: id,
	}

	jsonResponse(w, response)
}

func (api *API) UpdateWord(w http.ResponseWriter, r *http.Request) {
	upd := data.WordProperties{}
	err := parseFrom(w, r, &upd)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	id := numberParam(r, "id")

	err = api.service.Words.Update(id, upd)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		log.Printf("[ERROR] API.UpdateWord: %s", err)
		return
	}

	response := Response{
		ID: id,
	}

	jsonResponse(w, response)
}

func (api *API) DeleteWord(w http.ResponseWriter, r *http.Request) {
	id := numberParam(r, "id")
	err := api.service.Words.Delete(id)
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		log.Printf("[ERROR] API.DeleteWord: %s", err)
		return
	}

	response := Response{
		ID: id,
	}

	jsonResponse(w, response)
}

func (api *API) DeleteAllWords(w http.ResponseWriter, r *http.Request) {
	err := api.service.Words.DeleteAll()
	if err != nil {
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		log.Printf("[ERROR] API.DeleteAllWords: %s", err)
		return
	}
}


func (api *API) initRoutes(r *chi.Mux) {
	r.Get("/words", api.GetWords)
	r.Post("/words", api.AddWord)
	r.Put("/words/{id}", api.UpdateWord)
	r.Delete("/words/{id}", api.DeleteWord)
	r.Delete("/words", api.DeleteAllWords)
}

func jsonResponse(w http.ResponseWriter, data any) {
	bytes, err := json.Marshal(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(bytes)
}