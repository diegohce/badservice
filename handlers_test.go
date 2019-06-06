package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diegohce/logger"
	"github.com/julienschmidt/httprouter"
)

func TestDelayHandlerConversionError(t *testing.T) {

	log = logger.New("badservice-test::")

	params := httprouter.Params{
		httprouter.Param{"delay", "AAA"},
	}

	req, err := http.NewRequest("GET", "/badservice/delay/AAA", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	delayHandler(res, req, params)

	if res.Code == http.StatusOK {
		t.Error("Response code was: 200. Want 400")
	}
}

func TestDelayHandlerOK(t *testing.T) {

	params := httprouter.Params{
		httprouter.Param{"delay", "1"},
	}

	req, err := http.NewRequest("GET", "/badservice/delay/1", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	delayHandler(res, req, params)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was: %v. Want 200", res.Code)
	}
}

func TestStatusCodeHandlerConversionError(t *testing.T) {

	log = logger.New("badservice-test::")

	params := httprouter.Params{
		httprouter.Param{"code", "AAA"},
	}

	req, err := http.NewRequest("GET", "/badservice/status/AAA", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	statusCodeHandler(res, req, params)

	if res.Code == http.StatusOK {
		t.Error("Response code was: 200. Want 400")
	}
}

func TestStatusCodeHandlerValueError(t *testing.T) {

	log = logger.New("badservice-test::")

	params := httprouter.Params{
		httprouter.Param{"code", "10000"},
	}

	req, err := http.NewRequest("GET", "/badservice/status/10000", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	statusCodeHandler(res, req, params)

	if res.Code == http.StatusOK {
		t.Error("Response code was: 200. Want 999")
	}
}

func TestStatusCodeHandlerOK(t *testing.T) {

	log = logger.New("badservice-test::")

	params := httprouter.Params{
		httprouter.Param{"code", "500"},
	}

	req, err := http.NewRequest("GET", "/badservice/status/500", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	statusCodeHandler(res, req, params)

	if res.Code != 500 {
		t.Errorf("Response code was: %v. Want 500", res.Code)
	}
}
func TestDropConnectionHandler(t *testing.T) {

	log = logger.New("badservice-test::")

	params := httprouter.Params{}

	req, err := http.NewRequest("GET", "/badservice/drop", nil)
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()

	dropConnectionHandler(res, req, params)

	if res.Code != 400 {
		t.Errorf("Response code was: %v. Want 400", res.Code)
	}
}

