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

var server *http.Server

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

func TestGetUserRoute(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/users/1")
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

func TestPostUserRoute(t *testing.T) {
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	bytesRepresentation, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/users/create", "application/json", bytes.NewBuffer(bytesRepresentation))
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
}

func TestPutUserRoute(t *testing.T) {
	post := map[string]string{
		"title":   "testtitle",
		"content": "testcontent",
	}
	bytesRepresentation, _ := json.Marshal(post)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/users/1/post/1/update", bytes.NewBuffer(bytesRepresentation))
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
}

func TestDeleteUserRoute(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/users/1", nil)
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
