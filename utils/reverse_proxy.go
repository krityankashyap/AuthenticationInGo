package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, prefixPath string) http.HandlerFunc {
	target , err := url.Parse(targetBaseURL) // Parse parses a raw url into a [URL] structure.

   if err != nil {
		fmt.Println("Error parshing target url: ", err)
		return nil
	 }

  proxy := httputil.NewSingleHostReverseProxy(target)
	originalDirectory := proxy.Director

	proxy.Director = func(r *http.Request){
		originalDirectory(r)

		
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefixPath)

		r.Host = target.Host

		if userId , ok := r.Context().Value("userId").(string); ok{
			r.Header.Set("X-user-ID", userId)
		}
	}

	return proxy.ServeHTTP

}