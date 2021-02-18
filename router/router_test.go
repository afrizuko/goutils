package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {

	r := NewRouter()
	t.Run("it returns a string with status OK", func(t *testing.T) {
		r.Handle("GET", "/", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprint(rw, "Hi you just landed at the base")
		})

		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "Hi you just landed at the base", string(data))
	})

	t.Run("it returns a greeting with passed name", func(t *testing.T) {
		r.GET("/users/:name", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(rw, "Hi %s", GetParam(r, "name"))
		})

		req := httptest.NewRequest("GET", "/users/John", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "Hi John", string(data))
	})

}

func TestMultiPathRouter(t *testing.T) {

	r := NewRouter()
	t.Run("it returns a string with status OK", func(t *testing.T) {
		r.Handle("GET", "/:name", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(rw, "Hello %s", GetParam(r, "name"))
		})

		req := httptest.NewRequest("GET", "/Joseph", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "Hello Joseph", string(data))
	})

	t.Run("it returns a greeting with passed name", func(t *testing.T) {
		r.GET("/users/:login/status/:id", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(rw, "Hi %s Your id is %s", GetParam(r, "login"), GetParam(r, "id"))
		})

		req := httptest.NewRequest("GET", "/users/John/status/5", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "Hi John Your id is 5", string(data))
	})
}

func TestPostRouter(t *testing.T) {

	r := NewRouter()
	t.Run("it returns a string with status OK", func(t *testing.T) {
		r.POST("/users", func(rw http.ResponseWriter, r *http.Request) {
			data, _ := ioutil.ReadAll(r.Body)
			fmt.Fprintf(rw, string(data))
		})

		req := httptest.NewRequest("POST", "/users", bytes.NewBufferString("Name: JOhn"))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "Name: JOhn", string(data))
	})

	t.Run("it returns a greeting with passed name", func(t *testing.T) {
		r.PUT("/users/:id", func(rw http.ResponseWriter, r *http.Request) {
			id := GetParam(r, "id")

			var result map[string]string
			json.NewDecoder(r.Body).Decode(&result)
			result["id"] = id
			json.NewEncoder(rw).Encode(&result)
		})

		req := httptest.NewRequest("PUT", "/users/5", bytes.NewBufferString(`{"Name":"Peter"}`))
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, 200, res.Result().StatusCode)

		var result map[string]string
		err := json.NewDecoder(res.Body).Decode(&result)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, "5", result["id"])
	})
}
