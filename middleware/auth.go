package middleware

import (
	"context"
	"github.com/lpcheng1208/pkgutils/staticdefine"
	"github.com/lpcheng1208/pkgutils/tokenhelper"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func AuthToken(redisConn *redis.Redis) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Middleware", "auth")

			token := r.Header.Get("X-Token")
			if token == "" {
				httpx.WriteJson(w, http.StatusOK, &ApiResponse{Msg: "auth err, please login", Code: 13})
				return
			}

			userId, err := tokenhelper.ParseTokenToUserId(token)

			if err != nil {
				httpx.WriteJson(w, http.StatusOK, &ApiResponse{Msg: "auth err, please login", Code: 13})
				return
			}
			keyName := staticdefine.GetUserTokenKey(userId)
			rdsToken, _ := redisConn.Get(keyName)

			if rdsToken != token {
				httpx.WriteJson(w, http.StatusOK, &ApiResponse{Msg: "auth err, please login", Code: 13})
				return
			}

			ctx := r.Context()

			newCtx := context.WithValue(ctx, "userId", userId)
			//没报错就继续
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
