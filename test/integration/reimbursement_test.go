package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubmitReimbursement(t *testing.T) {
	token, _ := getToken("employee1", "pass1")
	reqBody := map[string]interface{}{
		"date":        time.Now().Format("2006-01-02"),
		"amount":      150000.0,
		"description": "Internet reimbursement",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, baseURL+"/api/reimbursement/submit", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-reimbursement-123")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)
}
