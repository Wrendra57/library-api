package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/be/perpustakaan/helper"
	"github.com/be/perpustakaan/model/webresponse"
	"github.com/julienschmidt/httprouter"
)
type ContextKey string

const (
	// KeyOne and KeyTwo are keys to access values in the context.
	Id ContextKey = "id"
	Email ContextKey = "email"
	Level ContextKey = "level"
	Token ContextKey = "token"
)
func AuthMiddleware(next httprouter.Handle)httprouter.Handle {
 
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("authHeader is empty")
			webResponse := webresponse.ResponseApi{
			Code:   http.StatusUnauthorized,
			Status: "unauthorized",
			Data:   nil,
		}
			helper.WriteToResponseBody(w, webResponse)

			return
		}

		// Check if the Authorization header has the Bearer prefix
		if !strings.HasPrefix(authHeader,"Bearer "){
			fmt.Println("authHeade not has the Bearer prefix")
			webResponse := webresponse.ResponseApi{
			Code:   http.StatusUnauthorized,
			Status: "unauthorized2",
			Data:   nil,
		}
			helper.WriteToResponseBody(w, webResponse)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		result ,err:=helper.ParseJWT(tokenString)
		

		if err != nil {
			fmt.Println("error parsing")
			webResponse := webresponse.ResponseApi{
			Code:   http.StatusUnauthorized,
			Status: "unauthorized3",
			Data:   nil,
		}
			helper.WriteToResponseBody(w, webResponse)
			return
		}
		 
		// Attach the parsed token to the request context for later use
		ctx := context.WithValue(r.Context(), "token", tokenString)
		ctx = context.WithValue(ctx,"id",result.Id)
		ctx = context.WithValue(ctx,"email",result.Email)
		ctx = context.WithValue(ctx,"level",result.Level)
		 
		next(w, r.WithContext(ctx), p)
		// next(w, r.WithContext(ctx), p)
	}
}

func RoleMiddleware(allowedRole string, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Retrieve the parsed token from the request context
		userRole, ok := r.Context().Value("level").(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the user has the required role
		
		if userRole != allowedRole {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// If the role is valid, call the next handler
		next(w, r, ps)
	}
}