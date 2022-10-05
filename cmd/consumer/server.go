package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./client/index.html")
}

func ordersHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	order, ok := cache[vars["uid"]]
	if !ok {
		var err error
		order, err = getByUid(vars["uid"])
		if err != nil {
			http.NotFound(w, req)
		}
	}
	w.Write([]byte(order))
}

func runServer() {
	r := mux.NewRouter()
	// статика
	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/api/orders/{uid}", ordersHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
