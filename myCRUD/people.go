package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

type People struct {
	PID     guuid.UUID `json:"PID"`
	Name    string     `json:"Name"`
	Surname string     `json:"Surname"`
	Age     int        `json:"Age"`
}

type Peoples struct {
	pp []People
}

// ReturnAllPeoples return all entries
func (p *Peoples) ReturnEntries(w http.ResponseWriter, r *http.Request) {

	CheckError(json.NewEncoder(w).Encode(p.pp))

}

// ReturnSinglePeople return single entri by ID
func (p *Peoples) ReturnSingleEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	find, pid := p.CheckForEntry(key)
	if find {
		CheckError(json.NewEncoder(w).Encode(p.pp[pid]))
	}
}

// CreateNewPeople create single entri
func (p *Peoples) CreateNewEntry(w http.ResponseWriter, r *http.Request) {
	var people People

	CheckError(json.NewDecoder(r.Body).Decode(&people))
	people.PID = guuid.New()
	p.pp = append(p.pp, people)
	CheckError(json.NewEncoder(w).Encode(people))

}

// DeletePeople delete single entry by ID
func (p *Peoples) DeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	find, idr := p.CheckForEntry(id)
	if find {
		p.pp = append(p.pp[:idr], p.pp[idr+1:]...)
		fmt.Fprintln(w, "Entry was deleted!")
	} else {
		fmt.Fprintln(w, "There is no such entry to delete!")
	}
}

// DeleteAllPeople delete all entries
func (p *Peoples) DeleteAllEntries(w http.ResponseWriter, r *http.Request) {

	p.pp = nil

	fmt.Fprintln(w, "All entries were deleted!")
}

// UpdatePeople update entry
func (p *Peoples) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	var newpeople People

	CheckError(json.NewDecoder(r.Body).Decode(&newpeople))
	find, pid := p.CheckForEntry(newpeople.PID.String())
	if find {
		p.pp[pid] = newpeople
		CheckError(json.NewEncoder(w).Encode(p.pp[pid]))
	} else {
		fmt.Fprintln(w, "There is no such people to update!")
	}
}

// CheckForEntry checking for entry with specific ID
func (p *Peoples) CheckForEntry(key string) (bool, int) {
	find := false
	j := -1
	for i, people := range p.pp {
		if people.PID.String() == key {
			j = i
			find = true
		}
	}
	return find, j
}
