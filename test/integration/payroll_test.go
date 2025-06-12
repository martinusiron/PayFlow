package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunPayroll(t *testing.T) {
	token, _ := getToken("admin", "admin123")
	reqBody := map[string]interface{}{
		"start_date": "2025-05-01",
		"end_date":   "2025-05-30",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, baseURL+"/api/payroll/run", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-payroll-run")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)
}
