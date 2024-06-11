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

func TestCreatePost(t *testing.T) {
	// Test successful request
	postData := map[string]string{
		"title":   "testpost",
		"content": "testcontent",
	}
	jsonData, _ := json.Marshal(postData)

	resp, err := http.Post(DefaultServerURL+"/posts/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"Created post testpost"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}

	// Test unsuccessful request (missing content)
	postData = map[string]string{
		"title": "testpost",
	}
	jsonData, _ = json.Marshal(postData)

	resp, err = http.Post(DefaultServerURL+"/posts/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestGetPostBySlug(t *testing.T) {
	resp, err := http.Get(DefaultServerURL + "/posts/slug/testpost")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"Post with slug testpost"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}

func TestGetPostByFilter(t *testing.T) {
	resp, err := http.Get(DefaultServerURL + "/posts/filter/title/testpost")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"Post filtered by title with value testpost"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}

func TestUpdatePost(t *testing.T) {
	// Test successful request
	postData := map[string]string{
		"title":   "testpost",
		"content": "testcontent",
	}
	jsonData, _ := json.Marshal(postData)

	req, _ := http.NewRequest("PUT", DefaultServerURL+"/posts/update/testpost", bytes.NewBuffer(jsonData))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"Updated post testpost, with content: testcontent"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}

	// Test unsuccessful request (missing content)
	postData = map[string]string{
		"title": "testpost",
	}
	jsonData, _ = json.Marshal(postData)

	req, _ = http.NewRequest("PUT", DefaultServerURL+"/posts/update/testpost", bytes.NewBuffer(jsonData))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeletePost(t *testing.T) {
	req, _ := http.NewRequest("DELETE", DefaultServerURL+"/posts/remove/testpost", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"status":true,"msg":"Deleted post testpost"}`
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}
