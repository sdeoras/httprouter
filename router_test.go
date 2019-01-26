package httprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var rtr Router

func init() {
	rtr = NewRouter()
	rtr.Register("a", a)
	rtr.Register("a/b", ab)
	rtr.Register("a/b/c", abc)
	rtr.Register("/", root)
}

func a(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "a")
}

func ab(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ab")
}

func abc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "abc")
}

func entryPoint(w http.ResponseWriter, r *http.Request) {
	rtr.Route(w, r)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "root")
}

func TestNewRouterA(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/a", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(entryPoint)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "a" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "a")
	}
}

func TestNewRouterAB(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/a/b", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(entryPoint)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "ab" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "ab")
	}
}

func TestNewRouterABC(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/a/b/c", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(entryPoint)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "abc" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "abc")
	}
}

func TestNewRouterABC2(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/a/B/c", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(entryPoint)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "abc" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "abc")
	}
}

func TestRoot(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(entryPoint)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "root" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "root")
	}
}
