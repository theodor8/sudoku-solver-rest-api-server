package handlers

import (
	"encoding/json"
	"net/http"
	"sudoku-server/api"
	"sudoku-server/internal/solver"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)



func ValidHandler(w http.ResponseWriter, r *http.Request) {
    params := api.SolveParams{}
    decoder := schema.NewDecoder()
    
    err := decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error("failed to decode parameters: ", err)
        api.InternalErrorHandler(w)
        return
    }

    valid := solver.IsValid(params.Input)

    w.Header().Set("Content-Type", "application/json")
    response := api.ValidResponse{
        Code: http.StatusOK,
        Valid: valid,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("failed to encode response: ", err)
        api.InternalErrorHandler(w)
    }
}
