package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

func middleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// right not all this does is log like
		// "github.com/gorilla/handlers"
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		fmt.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		// However since this is middleware you can have it do other things
		// Examples, auth users, write to file, redirects, handle panics, ect
		// add code to log to statds, remove log.Printf if you want

		handler.ServeHTTP(w, r)
	})
}

type People struct {
	PID     guuid.UUID `json:"PID"`
	Name    string     `json:"Name"`
	Surname string     `json:"Surname"`
	Age     int        `json:"Age"`
}

var Peoples []People

func main() {

	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Petro", Surname: "Poroshenko", Age: 55})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Petro", Surname: "Poroshenko", Age: 55})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Pavlo", Surname: "Kucenko", Age: 47})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Olga", Surname: "Kapusta", Age: 67})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Ivan", Surname: "Doroshenko", Age: 52})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Jon", Surname: "Yaroshenko", Age: 25})
	Peoples = append(Peoples, People{PID: guuid.New(), Name: "Sabrina", Surname: "Poroshenko", Age: 35})

	rout := mux.NewRouter()
	//	rout.Methods("GET", "POST", "DELETE", "PUT")
	rout.HandleFunc("/api/v1/list", returnAllPeoples).Methods("GET")
	rout.HandleFunc("/api/v1/list/{id}", returnSinglePeople).Methods("GET")
	rout.HandleFunc("/api/v1/list", createNewPeople).Methods("POST")
	rout.HandleFunc("/api/v1/list/{id}", deletePeople).Methods("DELETE")
	rout.HandleFunc("/api/v1/list", deleteAllPeople).Methods("DELETE")
	rout.HandleFunc("/api/v1/list", updatePeople).Methods("PUT")

	//rout.Host("localhost:8080")
	http.Handle("/", rout)
	/*srv := &http.Server{
		Handler:      middleWare(rout),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())*/
	err = http.ListenAndServe(":8080", middleWare(rout))
}

func returnAllPeoples(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Peoples)
}

func returnSinglePeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, people := range Peoples {
		if people.PID.String() == key {
			json.NewEncoder(w).Encode(people)
		}
	}
}

func createNewPeople(w http.ResponseWriter, r *http.Request) {
	var people People
	_ = json.NewDecoder(r.Body).Decode(&people)
	people.PID = guuid.New()
	Peoples = append(Peoples, people)

}

func deletePeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, people := range Peoples {
		if people.PID.String() == id {
			Peoples = append(Peoples[:index], Peoples[index+1:]...)
		}
	}
}

func deleteAllPeople(w http.ResponseWriter, r *http.Request) {
	Peoples = nil
}

func updatePeople(w http.ResponseWriter, r *http.Request) {
	var newpeople People
	_ = json.NewDecoder(r.Body).Decode(&newpeople)
	find := false
	for i, people := range Peoples {
		if people.PID.String() == newpeople.PID.String() {
			Peoples[i] = poeple
			//			people.Name = newpeople.Name
			//			people.Surname = newpeople.Surname
			//			people.Age = newpeople.Age
			find = true
		}
	}
	if !find {
		fmt.Fprintln(w, "There is no such peple to update!")
	}

}
