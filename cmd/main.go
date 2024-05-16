package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	manager "simple-server/src/managers"
	route "simple-server/src/routes"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	rm := manager.NewRequestManager()

	router := mux.NewRouter()

	// Add application routes
	route.AddRoutes(rm, router)

	router.Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		fmt.Printf("%v %s\n", methods, path)
		return nil
	})

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", "6543"),

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is running on :%s\n", "6543")
	log.Fatal(srv.ListenAndServe())
}
