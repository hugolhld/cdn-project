package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	responses "cdn-project/Responses"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Expected format: "Bearer <TOKEN>"
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return ""
	}

	return splitToken[1]
}

// ProtectedHandler - Middleware générique pour protéger les routes avec JWT
func ProtectedHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// ✅ Extraire le token JWT
		tokenString := ExtractToken(r)
		if tokenString == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    map[string]interface{}{"data": "Missing token"},
			}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// ✅ Valider le token JWT
		_, err := ValidateJWT(tokenString)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{
				Status:  http.StatusUnauthorized,
				Message: "error",
				Data:    map[string]interface{}{"data": "Invalid token"},
			}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// ✅ Si le token est valide, exécuter le handler
		handlerFunc.ServeHTTP(rw, r)
	}
}
