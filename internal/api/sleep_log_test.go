package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSleepLog(t *testing.T) {
	createUser(t, "sleeploguser1", strongPassword)
	response := loginUser(t, "sleeploguser1", strongPassword)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationTypeBearer, response.Token)

	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			"OK",
			gin.H{
				"startTime": "2021-01-01T00:00:00Z",
				"endTime":   "2021-01-01T08:00:00Z",
				"quality":   "Good",
			},
			func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			url := "/api/v1/sleep-logs"
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Set(authorizationHeaderKey, authorizationHeader)

			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
