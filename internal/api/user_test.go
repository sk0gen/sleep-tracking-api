package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

const strongPassword = "Str0ngP4ssw0rd!@"

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			"OK",
			gin.H{
				"username": "registeruser1",
				"password": strongPassword,
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"Empty Password",
			gin.H{
				"username": "registeruser2",
				"password": "",
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			"Invalid username",
			gin.H{
				"username": "registeruser#3",
				"password": strongPassword,
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			url := "/api/v1/auth/register"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestLoginUser(t *testing.T) {

	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			"OK",
			gin.H{
				"username": "loginuser1",
				"password": strongPassword,
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			"Empty Password",
			gin.H{
				"username": "loginuser2",
				"password": "",
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			"Invalid password",
			gin.H{
				"username": "loginuser3",
				"password": "123",
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()
			createUser(t, tc.body["username"].(string), strongPassword)

			url := "/api/v1/auth/login"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func createUser(t *testing.T, username string, password string) {
	server := newTestServer(t)
	recorder := httptest.NewRecorder()

	url := "/api/v1/auth/register"
	data, err := json.Marshal(gin.H{
		"username": username,
		"password": password,
	})
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func loginUser(t *testing.T, username string, password string) loginUserResponse {
	server := newTestServer(t)
	recorder := httptest.NewRecorder()

	url := "/api/v1/auth/login"
	data, err := json.Marshal(gin.H{
		"username": username,
		"password": password,
	})
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	var response loginUserResponse
	err = json.NewDecoder(recorder.Body).Decode(&response)
	return response
}
