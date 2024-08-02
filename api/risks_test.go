package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vivekgeorgemathew/aw/db/models"
	"github.com/vivekgeorgemathew/aw/db/store"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRisk(t *testing.T) {
	dbStore := store.NewStore()
	server := NewServer(dbStore)

	tests := []struct {
		req        CreateRiskRequest
		statusCode int
	}{
		{
			req: CreateRiskRequest{
				Title:       "Data Risk",
				Description: "Public S3 Bucket Risk",
				State:       "open",
			},
			statusCode: http.StatusOK,
		}, {
			req: CreateRiskRequest{
				Title:       "Network Risk",
				Description: "VPC SG Risk",
				State:       "invalid state"},
			statusCode: http.StatusBadRequest,
		},
	}
	for i := range tests {
		requestBody := tests[i].req
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(requestBody)
		if err != nil {
			t.Fatalf("Failed to encode request body %v", err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", &b)
		rec := httptest.NewRecorder()
		server.router.ServeHTTP(rec, req)
		if rec.Code != tests[i].statusCode {
			t.Errorf("Expected %d but got %d", tests[i].statusCode, rec.Code)
		}
	}
}

func TestGetRisk(t *testing.T) {
	dbStore := store.NewStore()
	server := NewServer(dbStore)

	// Create a risk
	requestBody := CreateRiskRequest{
		Title:       "Data Risk",
		Description: "Public S3 Bucket Risk",
		State:       "open",
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(requestBody)
	if err != nil {
		t.Fatalf("Failed to encode request body %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/risks", &b)
	rec := httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
	}
	data, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("Failed to read response body with error %v", err)
	}
	var risk models.Risk
	err = json.Unmarshal(data, &risk)
	if err != nil {
		t.Fatalf("Failed to unmarshall response data with error %v", err)
	}
	// Test using an invalid risk id
	req = httptest.NewRequest(http.MethodGet, "/api/v1/risks/invalid", nil)
	rec = httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected %d but got %d", http.StatusNotFound, rec.Code)
	}

	// Test using a valid risk id
	req = httptest.NewRequest(http.MethodGet,
		fmt.Sprintf("/api/v1/risks/%s", risk.RiskID), nil)
	rec = httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
	}
}

func TestGetAllRisks(t *testing.T) {
	dbStore := store.NewStore()
	server := NewServer(dbStore)

	// Test without any data in the store
	req := httptest.NewRequest(http.MethodGet, "/api/v1/risks", nil)
	rec := httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected %d but got %d", http.StatusInternalServerError, rec.Code)
	}

	// Test after creating a risk record
	requestBody := CreateRiskRequest{
		Title:       "Data Risk",
		Description: "Public S3 Bucket Risk",
		State:       "open",
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(requestBody)
	if err != nil {
		t.Fatalf("Failed to encode request body %v", err)
	}
	req = httptest.NewRequest(http.MethodPost, "/api/v1/risks", &b)
	rec = httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/risks", nil)
	rec = httptest.NewRecorder()
	server.router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Expected %d but got %d", http.StatusOK, rec.Code)
	}
	data, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Fatalf("Failed to read response body with error %v", err)
	}
	var risks []models.Risk
	err = json.Unmarshal(data, &risks)
	if err != nil {
		t.Fatalf("Failed to unmarshall response data with error %v", err)
	}
	if len(risks) != 1 {
		t.Errorf("Expected 1 record but got %d", len(risks))
	}
}
