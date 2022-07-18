package server_test

import (
	"MEND/pkg/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	for _, tt := range []struct {
		name string
		req  func() *http.Request
		prep func(repo *server.MockUserRepository)
	}{
		{},
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			repo := server.NewMockUserRepository(t)
			tt.prep(repo)
			controller := server.NewUserController(repo)
			router := server.NewRouter(*controller)

			router.ServeHTTP(w, tt.req())
		})
	}
}
