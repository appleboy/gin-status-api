package status

import (
	"encoding/json"
	"github.com/appleboy/gin-jwt-server/tests"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {

	tests.RunSimpleGet("/api/status",
		StatusHandler,
		func(r *httptest.ResponseRecorder) {
			var rd map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&rd)

			if err != nil {
				log.Fatalf("JSON Decode fail: %v", err)
			}

			assert.NotEmpty(t, rd["go_version"])
			assert.Equal(t, r.Code, http.StatusOK)
		})
}
