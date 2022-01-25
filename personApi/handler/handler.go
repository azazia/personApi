package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"personApi/conn"
	"personApi/data"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	db :=	conn.OpenConnection()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var people []data.Person

	for rows.Next() {
		var person data.Person
		rows.Scan(&person.Name, &person.Nickname)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	db := conn.OpenConnection()

	var p data.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sqlStatement := `INSERT INTO person(name, nickname) VALUES($1,$2)`
	_, err = db.Exec(sqlStatement, p.Name, p.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}