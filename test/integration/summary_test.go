package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdminSummary(t *testing.T) {
	token, _ := getToken("admin", "admin123")
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/api/summary/admin?payroll_id=1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Request-ID", "test-summary")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)
}
