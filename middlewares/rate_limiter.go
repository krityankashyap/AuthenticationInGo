package middlewares

import (
	"net/http"
	"time"
	"golang.org/x/time/rate"
)

 var limiter = rate.NewLimiter(rate.Every(1*time.Second) , 5) // 5 request per sec

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter , r *http.Request){
      // Rate limiting logic goes here

			if !limiter.Allow(){
				http.Error(w, "Too many req are coming up", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w , r)
	})
}