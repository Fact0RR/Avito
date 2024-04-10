package middlewares

import "net/http"

func UserMiddleWare(admin_token,user_token string,next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		current_token := r.Header.Get("token")
		switch{
		case admin_token == current_token:
			next(w,r)
		case user_token == current_token:
			next(w,r)
		default:
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Пользователь не авторизован"))
		}
	}
}