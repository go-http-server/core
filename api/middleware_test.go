package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-http-server/core/plugin/pkg/token"
	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func addAuthMiddlware(t *testing.T, request *http.Request, tokenMaker token.TokenMaker, typeAuthorizationHeader, username string, roleId int, duration time.Duration) {
	token, err := tokenMaker.CreateToken(username, roleId, duration)
	require.NoError(t, err)
	require.NotNil(t, token)

	authorizationHeader := fmt.Sprintf("%s %s", typeAuthorizationHeader, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(*testing.T, *http.Request, token.TokenMaker)
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "PASS_AUTH_MIDDLEWARE",
			setupAuth: func(t *testing.T, r *http.Request, tm token.TokenMaker) {
				username := utils.RandomString(6)
				roleId := utils.RandomInt(1, 10)
				addAuthMiddlware(t, r, tm, authorizationTypeBearer, username, roleId, time.Minute)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
	}

	for _, currentCase := range testCases {
		t.Run(currentCase.name, func(t *testing.T) {
			testServer := newTestServer(t, nil)
			testURL := "/api/v1/delete-database"
			testServer.router.GET(testURL, authMiddleware(testServer.tokenMaker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, testURL, nil)
			require.NoError(t, err)

			currentCase.setupAuth(t, request, testServer.tokenMaker)
			testServer.router.ServeHTTP(recorder, request)
			currentCase.checkResponse(t, recorder)
		})
	}
}
