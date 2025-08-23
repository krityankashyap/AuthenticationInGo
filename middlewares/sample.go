package middlewares

import (
	"AuthInGo/utils"
	"net/http"
)

func RequestValidatorSample(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload any // Define the payload u want to expect
		
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

		next.ServeHTTP(w,r)
	})
}