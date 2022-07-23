package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called")
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("called", string(body))
	fmt.Fprintf(w, "vamshi")
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get object")
}
func HandleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/createStudent", createStudent)
	http.HandleFunc("/getStudent", getStudent)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
func main() {
	HandleRequests()
}
