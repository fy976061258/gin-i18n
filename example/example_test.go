package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	i18n "github.com/fy976061258/gin-i18n/v1"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestPingRoute(t *testing.T) {
	bundle := i18n.NewBundle(
		language.Chinese,
		"active.zh-CN.toml",
		"active.en-US.toml",
		"active.ja-JP.toml",
	)

	router := setupRouter(bundle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "你好,gin", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ping?lang=en-US", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello,gin", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/ping?lang=ja-JP", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "こんにちは,gin", w.Body.String())
}
