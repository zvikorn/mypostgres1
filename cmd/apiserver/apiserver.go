package apiserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"mypostgres1/cmd/service"
)

const (
	port = 8000
)

type Server struct {
	ScannerService *service.ScannerService
}

// NewServer creates a new server.
func NewServer() *Server {
	server := &Server{}
	server.ScannerService = service.NewScannerService()
	return server
}

func (s *Server) Init() error {
	// initialize mux router
	r := mux.NewRouter()
	//TODO start full scanner with one resource
	//err := s.ScannerService.Start()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Could not .... Reason: %v", err)
	//	os.Exit(1)
	//}

	// register handlers
	r.HandleFunc("/api/v1/resourcs", s.createResource).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/urns/{sname}", s.getURNsByServicName).Methods(http.MethodGet)

	// start listening ...
	fmt.Println("Server started. Waiting for requests....")
	err := http.ListenAndServe(":"+strconv.Itoa(port), r)
	if err != nil {
		fmt.Printf("Server failed starting. Error: %s", err)
	}
	return nil
}
