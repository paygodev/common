package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contctKey string

const userKey contctKey = "userID"

func CreateJWT(secret []byte, userID string, jwtExpirationInSeconds int64) (string, error) {
	expiration := time.Second * time.Duration(jwtExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromReq(r)

		token, err := validateToken(tokenString, secret)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("Invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, str)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromReq(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateToken(t string, secret string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteError(w, http.StatusForbidden, "permission denied")
}

func GetUserIDFromCtx(ctx context.Context) string {
	userID, ok := ctx.Value(userKey).(string)
	if !ok {
		return ""
	}
	return userID
}
