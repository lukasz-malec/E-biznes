package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/db"

	"github.com/labstack/echo/v4"
)

func setupEcho() *echo.Echo {
	e := echo.New()
	db.Init()
	return e
}


// testy API dla wczesniejsych napisanych endpointow, czyli GET i POST

// testy dla GET

func TestGetProducts_StatusOK(t *testing.T) {
	e := setupEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetProducts(c)

	if err != nil {
		t.Errorf("oczekiwano nil, dostano: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("oczekiwano 200, dostano: %d", rec.Code)
	}
}


func TestGetProducts_ZwracaJSON(t *testing.T) {
	e := setupEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetProducts(c)

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("oczekiwano application/json, dostano: %s", contentType)
	}
}


func TestGetProducts_ListaNiePusta(t *testing.T) {
	e := setupEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetProducts(c)

	var products []map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &products)
	if len(products) == 0 {
		t.Error("oczekiwano produktów, lista jest pusta")
	}
}



// scenariusz negatywny 
func TestGetProducts_ZłaMetoda(t *testing.T) {
	e := setupEcho()
	req := httptest.NewRequest(http.MethodPost, "/api/products", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code == http.StatusOK {
		t.Error("POST na /api/products powinien zwrócić błąd")
	}
}


// testy dla POST

func TestCreateOrder_PoprawneZamówienie(t *testing.T) {
	e := setupEcho()

	body := `{
		"items": [{"product_id": 1, "name": "Test", "quantity": 1, "price": 99.99}],
		"total": 99.99,
		"customer": {"name": "Jan", "email": "jan@test.pl", "address": "ul. Testowa 1"}
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateOrder(c)

	if rec.Code != http.StatusCreated {
		t.Errorf("oczekiwano 201, dostano: %d", rec.Code)
	}
}


func TestCreateOrder_ZwracaOrderID(t *testing.T) {
	e := setupEcho()

	body := `{
		"items": [{"product_id": 1, "name": "Test", "quantity": 1, "price": 99.99}],
		"total": 99.99,
		"customer": {"name": "Jan", "email": "jan@test.pl", "address": "ul. Testowa 1"}
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateOrder(c)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	if _, ok := response["order_id"]; !ok {
		t.Error("odpowiedź nie zawiera order_id")
	}
}



func TestCreateOrder_ZwracaStatusConfirmed(t *testing.T) {
	e := setupEcho()

	body := `{
		"items": [{"product_id": 1, "name": "Test", "quantity": 1, "price": 99.99}],
		"total": 99.99,
		"customer": {"name": "Jan", "email": "jan@test.pl", "address": "ul. Testowa 1"}
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateOrder(c)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	if response["status"] != "confirmed" {
		t.Errorf("oczekiwano confirmed, dostano: %v", response["status"])
	}
}


// scenariusz negatywny 
func TestCreateOrder_NieprawidłowyJSON(t *testing.T) {
	e := setupEcho()
	req := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader("nieprawidłowy json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateOrder(c)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("oczekiwano 400, dostano: %d", rec.Code)
	}
}