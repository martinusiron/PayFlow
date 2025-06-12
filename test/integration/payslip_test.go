package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPayslip(t *testing.T) {
	token, _ := getToken("employee1", "pass1")
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/api/payslip/get", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Request-ID", "test-payslip-fetch")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)
}
