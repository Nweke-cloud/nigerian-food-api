package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthReturns503WhenDBIsNil(t *testing.T) {
	db = nil
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handleHealth(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when db is nil, got %d", w.Code)
	}
}

func TestFoodsFallbackWhenDBIsNil(t *testing.T) {
	db = nil
	req := httptest.NewRequest(http.MethodGet, "/api/foods", nil)
	w := httptest.NewRecorder()

	handleFoods(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 from fallback, got %d", w.Code)
	}
}
