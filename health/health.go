package health

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/damiannolan/eros/server"
	"github.com/go-chi/chi"
)

// healthResource implements server.Resource interface
type healthResource struct {
	path string
}

// NewResource constructor func
func NewResource(path string) server.Resource {
	return &healthResource{
		path: path,
	}
}

// Path returns the Resource base path
func (r *healthResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *healthResource) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/status", r.healthCheck())

	return router
}

func (r *healthResource) healthCheck() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		payload := struct {
			Message    string `json:"message"`
			StatusCode int    `json:"statusCode"`
		}{
			fmt.Sprintf("Healthy response from service at - %s", time.Now()),
			200,
		}

		json, _ := json.Marshal(payload)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
