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
	DefaultServerURL = "http://localhost:8080/api"
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

func TestCreateUser(t *testing.T) {
	// Test successful request
	userData := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonData, _ := json.Marshal(userData)

	resp, err := http.Post(DefaultServerURL+"/users", "application/json", bytes.NewBuffer(jsonData))
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

	resp, err = http.Post(DefaultServerURL+"/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetUser(t *testing.T) {
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

func TestGetAllUsers(t *testing.T) {
	resp, err := http.Get(DefaultServerURL + "/users")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"All users"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}

func TestUpdateUser(t *testing.T) {
	// Test successful request
	userData := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonData, _ := json.Marshal(userData)

	req, _ := http.NewRequest(http.MethodPut, DefaultServerURL+"/users/1", bytes.NewBuffer(jsonData))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"User testuser, with id 1 updated"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}

	// Test unsuccessful request (missing password)
	userData = map[string]string{
		"username": "testuser",
	}
	jsonData, _ = json.Marshal(userData)

	req, _ = http.NewRequest(http.MethodPut, DefaultServerURL+"/users/1", bytes.NewBuffer(jsonData))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, DefaultServerURL+"/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
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
