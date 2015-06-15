package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

type Record struct {

	//DB schema:
	// +-------+-------------+------+-----+---------+----------------+
	// | Field | Type        | Null | Key | Default | Extra          |
	// +-------+-------------+------+-----+---------+----------------+
	// | id    | int(11)     | NO   | PRI | NULL    | auto_increment |
	// | name  | varchar(45) | NO   |     | NULL    |                |
	// | age   | int(11)     | NO   |     | NULL    |                |
	// | email | varchar(45) | YES  |     | NULL    |                |
	// +-------+-------------+------+-----+---------+----------------+

	Id    string
	Name  string
	Age   string
	Email string
}

var (
	lock = sync.RWMutex{} //mutex for thread safety
)

func main() {
	fmt.Println("..")
	fmt.Print(http.DefaultClient)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...) //use the native http stack
	router, err := rest.MakeRouter(
		rest.Post("/records", PostRecord),
	//	rest.Get("/records/:id", GetRecord),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler())) //Listening on port 8080 on localhost

}

func PostRecord(w rest.ResponseWriter, r *rest.Request) {
	record := Record{} //instanciating record struct
	err := r.DecodeJsonPayload(&record)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if record.Id == "" {
		rest.Error(w, "record ID is required", 400)
		return
	}
	if record.Name == "" {
		rest.Error(w, "human's name is required", 400)
		return
	}
	if record.Age == "" {
		rest.Error(w, "human's age is required", 400)
		return
	}
	if record.Email == "" {
		rest.Error(w, "human's e-mail is required", 400)
		return
	}
	lock.Lock()
	// add record.id = &record in a db
	lock.Unlock()
	w.WriteJson(&record)
	fmt.Println("OUTPUT: ", record)
}
