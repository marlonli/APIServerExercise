package handlers

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/marlonli/APIServerExercise/common"
	"github.com/marlonli/APIServerExercise/pkg/container"
)

// An endpoint for persist metadata
func PersistMetadata(w http.ResponseWriter, req *http.Request) {
	var m container.Metadata
	err := yaml.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if m.IsValid() {
		// If the payload is valid, persist the metadata
		common.MetaList = append(common.MetaList, m)

		w.Header().Set("Content-Type", "text/yaml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Persist successfully!"))
		//err := yaml.NewEncoder(w).Encode(m)
		//if err != nil {
		//	log.Fatalf("error: %v", err)
		//}
	} else {
		// Invalid payload
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request!"))
	}
}

func GetMetadata(w http.ResponseWriter, req *http.Request) {
	err := yaml.NewEncoder(w).Encode(common.MetaList)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// An endpoint for searching metadata
func SearchMetadata(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request!"))
	}

	var resultList []container.Metadata

	//search metadata matches the request
	for _, data := range common.MetaList {
		if data.Match(strings.TrimSpace(string(body))) {
			resultList = append(resultList, data)
		}
	}

	w.Header().Set("Content-Type", "text/yaml; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err = yaml.NewEncoder(w).Encode(resultList)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
