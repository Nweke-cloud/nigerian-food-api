package main

import (
	"database/sql"
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

func TestHealthReturns503WhenDBPingFails(t *testing.T) {
	// Point db at an unreachable address (port 1, nothing listening)
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:secret@127.0.0.1:1/fooddb?sslmode=disable&connect_timeout=1")
	if err != nil {
		t.Fatalf("sql.Open failed: %v", err)
	}
	defer func() { db = nil }()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handleHealth(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when db.Ping() fails, got %d", w.Code)
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