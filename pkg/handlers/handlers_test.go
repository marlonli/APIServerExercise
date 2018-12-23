package handlers

import (
	"bytes"
	"github.com/marlonli/APIServerExercise/common"
	"github.com/marlonli/APIServerExercise/pkg/container"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPersistMetadata(t *testing.T) {
	// Create a request
	reqBytes, err := ioutil.ReadFile("validRequest.yaml")
	reqBody := bytes.NewReader(reqBytes)

	req, err := http.NewRequest("POST", "/v1/persist", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(PersistMetadata)
	handler.ServeHTTP(r, req)

	// Check the status code.
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong status code!\n got: %v\n expected: %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Persist successfully!"
	if r.Body.String() != expected {
		t.Errorf("Unexpected body!\n got: %v\n expected: %v",
			r.Body.String(), expected)
	}
}

func TestSearchMetadata(t *testing.T) {
	// Persist test data first
	dataBytes, err := ioutil.ReadFile("validRequest.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var m container.Metadata
	err = yaml.Unmarshal(dataBytes, &m)
	if err != nil {
		t.Fatal(err)
	}
	common.MetaList = append(common.MetaList, m)

	// The expected response
	respBytes, err := yaml.Marshal(&common.MetaList)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request
	reqBody := strings.NewReader("firstmaintainer")
	req, err := http.NewRequest("POST", "/v1/search", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchMetadata)
	handler.ServeHTTP(r, req)

	// Check the status code.
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong status code!\n got: %v\n expected: %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := string(respBytes)
	if r.Body.String() != expected {
		t.Errorf("Unexpected body!\n got: %v\n expected: %v",
			r.Body.String(), expected)
	}
}
