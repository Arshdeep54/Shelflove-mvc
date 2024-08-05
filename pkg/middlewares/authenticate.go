package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/controllers"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
	"github.com/golang-jwt/jwt"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := r.Cookie("jwtToken")
		if err != nil {
			writeHeaders(w, r, next, 0, false, "", "")
			return
		}

		token, err := utils.ValidateJWT(jwtToken.Value)
		if err != nil {
			writeHeaders(w, r, next, 0, false, "", "")
			return
		}

		if !token.Valid {
			writeHeaders(w, r, next, 0, false, "", "")
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["user"].(string)

		var payload types.JwtPayload

		err = json.Unmarshal([]byte(str), &payload)
		if err != nil {
			writeHeaders(w, r, next, 0, false, "", "")
			return
		}

		writeHeaders(w, r, next, int(payload.Id), true, payload.Username, payload.Email)
	}
}

func writeHeaders(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, userId int, isLoggedIn bool, username string, email string) {
	isAdmin, err := models.IsAdmin(userId)
	if err != nil {
		isAdmin = false
	}

	controllers.Data.IsAdmin = isAdmin
	controllers.Data.UserId = userId
	controllers.Data.IsLoggedIn = isLoggedIn
	controllers.Data.Username = username
	controllers.Data.Email = email
	controllers.Data.UserId = userId

	next(w, r)
}
