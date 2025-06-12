package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

var (
	baseURL       = "http://localhost:8080"
	adminToken    string
	employeeToken string
	defaultClient = &http.Client{}
)

func TestMain(m *testing.M) {
	var err error
	adminToken, err = getToken("admin", "admin123")
	if err != nil {
		fmt.Println("Failed to get admin token:", err)
		os.Exit(1)
	}
	employeeToken, err = getToken("employee1", "pass1")
	if err != nil {
		fmt.Println("Failed to get employee token:", err)
		os.Exit(1)
	}
	code := m.Run()
	os.Exit(code)
}

func getToken(username, password string) (string, error) {
	creds := map[string]string{
		"username": username,
		"password": password,
	}
	body, _ := json.Marshal(creds)
	resp, err := http.Post(baseURL+"/api/auth/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to login: status %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["accessToken"].(string)
	if !ok {
		return "", fmt.Errorf("accessToken not found")
	}
	return token, nil
}

func doRequest(t *testing.T, method, path, token string, body interface{}) (*http.Response, []byte) {
	var buf io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal("Failed to marshal body:", err)
		}
		buf = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, baseURL+path, buf)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("X-Request-ID", "test-request-id")

	resp, err := defaultClient.Do(req)
	if err != nil {
		t.Fatal("Failed to send request:", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Failed to read response:", err)
	}

	return resp, respBody
}
