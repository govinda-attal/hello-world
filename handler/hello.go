package handler

import (
	"encoding/json"
	"net/http"

	"github.com/govinda-attal/hello-world/pkg/hw"
)

// HelloHandler implements methods to handle HTTP requests for the microservice.
type HelloHandler struct {
	hwAPI hw.HelloWorld
}

// NewHelloHandler returns a new HTTP handler instance for given HelloWorld API.
func NewHelloHandler(hwAPI hw.HelloWorld) *HelloHandler {
	return &HelloHandler{hwAPI}
}

// Hello on HTTP GET request returns a greeing message from the system.
func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) error {

	name := r.FormValue("name")
	rs, err := h.hwAPI.Hello(name)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)

	return nil
}
