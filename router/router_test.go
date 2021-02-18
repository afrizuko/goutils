package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {

	r := NewRouter()
	t.Run("it returns a string with status OK", func(t *testing.T) {
		r.Handle("GET", "/", func(rw http.ResponseWriter, r *http.Request, v url.Values) {
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
		r.GET("/users/:name", func(rw http.ResponseWriter, r *http.Request, v url.Values) {
			fmt.Fprintf(rw, "Hi %s", v["name"][0])
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
