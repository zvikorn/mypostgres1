package apiserver

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *Server) createResource(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getURNsByServicName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//TODO API validations

	sname := mux.Vars(r)["sname"]
	urns, err := s.ScannerService.GetURNsByServiceName(ctx, sname)
	if err != nil {
		log.Fatal(err)
	}
	writeJSONResponse(w, http.StatusOK, urns)

}
