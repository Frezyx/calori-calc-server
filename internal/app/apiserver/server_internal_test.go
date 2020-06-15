package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Frezyx/calory-calc-server/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name        string
		payload     interface{}
		exeptedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"email":    "test@gmail.com",
				"password": "123456",
			},
			exeptedCode: http.StatusCreated,
		},
		{
			name:        "invalid payload",
			payload:     "invalid",
			exeptedCode: http.StatusBadRequest,
		},
		{
			name: "invalid eamil",
			payload: map[string]interface{}{
				"email":    "invalid",
				"password": "123456",
			},
			exeptedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid password empty",
			payload: map[string]interface{}{
				"email": "invalid",
			},
			exeptedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.exeptedCode, rec.Code)
		})
	}
}
