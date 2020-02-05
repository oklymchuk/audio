package main

import (
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {

	var ps = new(Peoples)

	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Petro", Surname: "Poroshenko", Age: 55})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Pirto", Surname: "Ponko", Age: 55})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Pavlo", Surname: "Kucenko", Age: 47})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Olga", Surname: "Kapusta", Age: 67})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Ivan", Surname: "Doroshenko", Age: 52})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Jon", Surname: "Yaroshenko", Age: 25})
	ps.pp = append(ps.pp, People{PID: guuid.New(), Name: "Sabrina", Surname: "Poroshenko", Age: 35})

	rout := mux.NewRouter()
	//	apiListRoute := rout.PathPrefix("/api/v1").Subrouter()
	rout.HandleFunc("/api/v1/list/{id}", ps.ReturnSingleEntry).Methods("GET")
	rout.HandleFunc("/api/v1/list", ps.ReturnEntries).Methods("GET")
	rout.HandleFunc("api/v1/list", ps.CreateNewEntry).Methods("POST")
	rout.HandleFunc("/api/v1/list/{id}", ps.DeleteEntry).Methods("DELETE")
	rout.HandleFunc("api/v1/list", ps.DeleteAllEntries).Methods("DELETE")
	rout.HandleFunc("/api/v1/list", ps.UpdateEntry).Methods("PUT")

	//http.Handle("/", rout)
	CheckError(http.ListenAndServe(":8080", rout))
}
