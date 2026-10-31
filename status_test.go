package status

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func httpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/api/status", GinHandler)

	return r
}

func TestStatusHandler(t *testing.T) {
	r := gofight.New()

	r.GET("/api/status").
		Run(httpRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			value, _ := jsonparser.GetString(r.Body.Bytes(), "go_version")

			assert.NotEmpty(t, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGinHandler_ResponseStructure(t *testing.T) {
	r := gofight.New()

	r.GET("/api/status").
		Run(httpRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			// Test status code
			assert.Equal(t, http.StatusOK, r.Code)

			// Test content type
			assert.Contains(t, r.Result().Header.Get("Content-Type"), "application/json")

			body := r.Body.Bytes()

			// Test required fields exist and have valid values
			goVersion, err := jsonparser.GetString(body, "go_version")
			require.NoError(t, err)
			assert.NotEmpty(t, goVersion)

			goOs, err := jsonparser.GetString(body, "go_os")
			require.NoError(t, err)
			assert.NotEmpty(t, goOs)

			goArch, err := jsonparser.GetString(body, "go_arch")
			require.NoError(t, err)
			assert.NotEmpty(t, goArch)

			cpuNum, err := jsonparser.GetInt(body, "cpu_num")
			require.NoError(t, err)
			assert.Positive(t, cpuNum)

			goroutineNum, err := jsonparser.GetInt(body, "goroutine_num")
			require.NoError(t, err)
			assert.GreaterOrEqual(t, goroutineNum, int64(0))

			gomaxprocs, err := jsonparser.GetInt(body, "gomaxprocs")
			require.NoError(t, err)
			assert.Positive(t, gomaxprocs)

			cgoCalls, err := jsonparser.GetInt(body, "cgo_call_num")
			require.NoError(t, err)
			assert.GreaterOrEqual(t, cgoCalls, int64(0))
		})
}

func TestGinHandler_MemoryStats(t *testing.T) {
	r := gofight.New()

	r.GET("/api/status").
		Run(httpRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)

			body := r.Body.Bytes()

			// Test memory related fields
			memoryAlloc, err := jsonparser.GetInt(body, "memory_alloc")
			require.NoError(t, err)
			assert.GreaterOrEqual(t, memoryAlloc, int64(0))

			memoryTotalAlloc, err := jsonparser.GetInt(body, "memory_total_alloc")
			require.NoError(t, err)
			assert.Positive(t, memoryTotalAlloc)

			memorySys, err := jsonparser.GetInt(body, "memory_sys")
			require.NoError(t, err)
			assert.Positive(t, memorySys)

			memoryLookups, err := jsonparser.GetInt(body, "memory_lookups")
			require.NoError(t, err)
			assert.GreaterOrEqual(t, memoryLookups, int64(0))

			memoryMallocs, err := jsonparser.GetInt(body, "memory_mallocs")
			require.NoError(t, err)
			assert.Positive(t, memoryMallocs)

			memoryFrees, err := jsonparser.GetInt(body, "memory_frees")
			require.NoError(t, err)
			assert.GreaterOrEqual(t, memoryFrees, int64(0))
		})
}

func TestGinHandler_MultipleRequests(t *testing.T) {
	router := httpRouter()

	// Test multiple sequential requests
	for range 5 {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/api/status",
			nil,
		)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		goVersion, err := jsonparser.GetString(w.Body.Bytes(), "go_version")
		require.NoError(t, err)
		assert.NotEmpty(t, goVersion)
	}
}

func BenchmarkGinHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/status", GinHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/status", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Body.Reset()
		router.ServeHTTP(w, req)
	}
}
