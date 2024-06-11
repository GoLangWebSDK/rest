package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	DefaultServerURL = "http://localhost:8080"
)

func TestMain(m *testing.M) {
	// Setup: start the server
	go func() {
		main()
	}()

	// give the server some time to start
	time.Sleep(time.Second * 2)

	// Run the tests
	code := m.Run()
	os.Exit(code)
}

func TestGetRequest(t *testing.T) {
	resp, err := http.Get(DefaultServerURL + "/users/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"User with id 1"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}

func TestPostRequest(t *testing.T) {
	// Test successful request
	userData := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonData, _ := json.Marshal(userData)

	resp, err := http.Post(DefaultServerURL+"/users/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"User testuser created"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}

	// Test unsuccessful request (missing password)
	userData = map[string]string{
		"username": "testuser",
	}
	jsonData, _ = json.Marshal(userData)

	resp, err = http.Post(DefaultServerURL+"/users/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestPutRequest(t *testing.T) {
	// Test successful request
	post := map[string]string{
		"title":   "testtitle",
		"content": "testcontent",
	}
	bytesRepresentation, _ := json.Marshal(post)
	req, _ := http.NewRequest(http.MethodPut, DefaultServerURL+"/users/1/post/1/update", bytes.NewBuffer(bytesRepresentation))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"User 1 updated post 1"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}

	// Test unsuccessful request (missing content)
	post = map[string]string{
		"title": "testtitle",
	}
	bytesRepresentation, _ = json.Marshal(post)
	req, _ = http.NewRequest(http.MethodPut, DefaultServerURL+"/users/1/post/1/update", bytes.NewBuffer(bytesRepresentation))
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteRequest(t *testing.T) {
	req, err := http.NewRequest("DELETE", DefaultServerURL+"/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"User 1 deleted"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}
