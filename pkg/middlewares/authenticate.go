package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"github.com/golang-jwt/jwt"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := r.Cookie("jwtToken")
		if err != nil {
			writeHeaders(w, r, next, 0, false, false, "", "")
			return
		}
		token, err := utils.ValidateJWT(jwtToken.Value)
		if err != nil {
			writeHeaders(w, r, next, 0, false, false, "", "")
			return
		}

		if !token.Valid {
			writeHeaders(w, r, next, 0, false, false, "", "")
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["user"].(string)

		var payload types.JwtPayload
		err = json.Unmarshal([]byte(str), &payload)
		if err != nil {
			writeHeaders(w, r, next, 0, false, false, "", "")
			return
		}
		writeHeaders(w, r, next, int(payload.Id), payload.IsAdmin, true, payload.Username, payload.Email)
	}
}

func writeHeaders(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, userId int, IsAdmin bool, IsLoggedIn bool, username string, email string) {
	controllers.Data.UserId = userId
	controllers.Data.IsAdmin = IsAdmin
	controllers.Data.IsLoggedIn = IsLoggedIn
	controllers.Data.Username = username
	controllers.Data.Email = email
	controllers.Data.UserId = userId

	next(w, r)
}
