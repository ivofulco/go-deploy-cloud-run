package main

import (
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Handle(t *testing.T) {
	t.Run("deve retornar 422 quando o cep invalido", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.SetPathValue("cep", "123456")
		Handle(w, r)
		if w.Code != 422 {
			t.Errorf("esperado 422, em %d", w.Code)
		}
	})
 
	t.Run("deve retornar 404 quando em cep nao encontrado", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		r.SetPathValue("cep", "12345678")
		w := httptest.NewRecorder()

		Handle(w, r)
		assert.Equal(t, w.Code, 404)
	})

	t.Run("deve retornar 200 como esperado", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		r.SetPathValue("cep", "01153000")
		w := httptest.NewRecorder()

		Handle(w, r)
		assert.Equal(t, w.Code, 200)
		assert.NotNil(t, w.Body)
	})
}

func Test_validCep(t *testing.T) {
	t.Run("deve retornar verdadeiro quando ao cep valido", func(t *testing.T) {
		assert.True(t, validCep("29216070"))
	})

	t.Run("deve retornar falso quando ao cep invalido", func(t *testing.T) {
		assert.False(t, validCep("1234567"))
	})
}
