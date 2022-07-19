package server_test

import (
	"MEND/pkg/models"
	"MEND/pkg/server"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	for _, tt := range []struct {
		name           string
		req            func() *http.Request
		prep           func(repo *server.MockUserRepository)
		assertResponse func(statusCode int, data []byte)
	}{
		{
			name: "GET - 200 successful",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
				require.NoError(t, err)
				return req
			},
			prep: func(repo *server.MockUserRepository) {
				repo.
					On("Get", mock.AnythingOfType("*context.emptyCtx"), 1).
					Return(&models.User{ID: 1, Name: "test name", Surname: "test surname"}, nil)
			},
			assertResponse: func(statusCode int, data []byte) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Equal(t, `{"id":1,"name":"test name","surname":"test surname"}`, string(data))
			},
		},
		{
			name: "GET - 400 failed, invalid param",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/users/1qw", nil)
				require.NoError(t, err)
				return req
			},
			prep: func(repo *server.MockUserRepository) {},
			assertResponse: func(statusCode int, data []byte) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Equal(t, `{"error":"invalid user ID passed in request"}`, string(data))
			},
		},
		{
			name: "GET - 500 failed, error from repository",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
				require.NoError(t, err)
				return req
			},
			prep: func(repo *server.MockUserRepository) {
				repo.
					On("Get", mock.AnythingOfType("*context.emptyCtx"), 1).
					Return(nil, errors.New("test error"))
			},
			assertResponse: func(statusCode int, data []byte) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.Equal(t, `{"details":"test error","error":"unknown error"}`, string(data))
			},
		},
		// NOTE: rest of endpoints would be tested in exactly the same way, skipping.
	} {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			repo := server.NewMockUserRepository(t)
			tt.prep(repo)
			controller := server.NewUserController(repo)
			router := server.NewRouter(controller)

			router.ServeHTTP(w, tt.req())

			body, err := ioutil.ReadAll(w.Body)
			require.NoError(t, err)
			tt.assertResponse(w.Result().StatusCode, body)
			repo.AssertExpectations(t)
		})
	}
}
