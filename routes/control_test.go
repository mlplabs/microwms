package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	type PingResponseTest struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}

	r := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	w := httptest.NewRecorder()

	Ping(w, r)

	data := PingResponseTest{}
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}
	if data.Code != http.StatusOK {
		t.Errorf("Invalid")
	}
}

func TestGetHealth(t *testing.T) {
	// Формируем реквест параметры тут нас не интересуют совсем. мы их не проверяем
	//r := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	// Готовим рекордер ответв
	//w := httptest.NewRecorder()
	// выполняем наш обработчик
	// TODO: отключил проверку необходим mock
	//GetHealth(w, r)
	// тут проверяем ответ

}
