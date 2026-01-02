package status

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
			assert.Contains(t, r.HeaderMap.Get("Content-Type"), "application/json")

			body := r.Body.Bytes()

			// Test required fields exist and have valid values
			goVersion, err := jsonparser.GetString(body, "go_version")
			assert.NoError(t, err)
			assert.NotEmpty(t, goVersion)

			goOs, err := jsonparser.GetString(body, "go_os")
			assert.NoError(t, err)
			assert.NotEmpty(t, goOs)

			goArch, err := jsonparser.GetString(body, "go_arch")
			assert.NoError(t, err)
			assert.NotEmpty(t, goArch)

			cpuNum, err := jsonparser.GetInt(body, "cpu_num")
			assert.NoError(t, err)
			assert.Greater(t, cpuNum, int64(0))

			goroutineNum, err := jsonparser.GetInt(body, "goroutine_num")
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, goroutineNum, int64(0))

			gomaxprocs, err := jsonparser.GetInt(body, "gomaxprocs")
			assert.NoError(t, err)
			assert.Greater(t, gomaxprocs, int64(0))

			cgoCalls, err := jsonparser.GetInt(body, "cgo_call_num")
			assert.NoError(t, err)
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
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, memoryAlloc, int64(0))

			memoryTotalAlloc, err := jsonparser.GetInt(body, "memory_total_alloc")
			assert.NoError(t, err)
			assert.Greater(t, memoryTotalAlloc, int64(0))

			memorySys, err := jsonparser.GetInt(body, "memory_sys")
			assert.NoError(t, err)
			assert.Greater(t, memorySys, int64(0))

			memoryLookups, err := jsonparser.GetInt(body, "memory_lookups")
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, memoryLookups, int64(0))

			memoryMallocs, err := jsonparser.GetInt(body, "memory_mallocs")
			assert.NoError(t, err)
			assert.Greater(t, memoryMallocs, int64(0))

			memoryFrees, err := jsonparser.GetInt(body, "memory_frees")
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, memoryFrees, int64(0))
		})
}

func TestGinHandler_MultipleRequests(t *testing.T) {
	router := httpRouter()

	// Test multiple sequential requests
	for i := 0; i < 5; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/status", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		goVersion, err := jsonparser.GetString(w.Body.Bytes(), "go_version")
		assert.NoError(t, err)
		assert.NotEmpty(t, goVersion)
	}
}

func BenchmarkGinHandler(b *testing.B) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/status", GinHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/status", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Body.Reset()
		router.ServeHTTP(w, req)
	}
}
