package middlewares

import (
	"AuthInGo/dtos"
	"AuthInGo/utils"
	"net/http"
	"context"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.UserLogindto // Define the payload u want to expect
		
		// Read and decode the JSON body into the payload
		err := utils.Readjson(*r , &payload)

		if err != nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		
		// Validate the payload using the Validator instance
		validationerr := utils.Validator.Struct(payload)

		if validationerr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", validationerr)
			return
		}
    
		reqContext := r.Context() // this is the context of req object or to say parent context

		ctx := context.WithValue(reqContext, "payload", payload) // Created a new context with payload

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var payload dtos.CreateUserRequestDTO

		err:= utils.Readjson(*r , &payload)

		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		validationerr := utils.Validator.Struct(payload)

		if validationerr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", validationerr)
			return
		}

		reqContext := r.Context() // this is the context of req object or to say parent context

		ctx := context.WithValue(reqContext, "payload", payload) // Created a new context with payload

		next.ServeHTTP(w , r.WithContext(ctx))
	})
}