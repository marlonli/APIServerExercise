package main

import (
	"log"
	"net/http"

	"github.com/marlonli/APIServerExercise/pkg/routers"
)

func main() {
	// create mock data
	// metaList = append(metaList, m.Metadata{ID: "1",
	// 	Title:       "App w/ Invalid maintainer email",
	// 	Version:     "1.0.1",
	// 	Maintainers: []&Person{name: "Firstname Lastname", email: "apptwo@hotmail.com"},
	// 	Company:     "Upbound Inc.",
	// 	Website:     "https://upbound.io",
	// 	Source:      "https://github.com/upbound/repo",
	// 	License:     "Apache-2.0",
	// 	Description: "|### blob of markdown More markdown})"})

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

	// Test request
	// yamlFile, err := ioutil.ReadFile("validRequest.yaml")
	// fmt.Println(yamlFile)
	// if err != nil {
	// 	log.Printf("yamlFile err: #%v\n", err)
	// }
}
