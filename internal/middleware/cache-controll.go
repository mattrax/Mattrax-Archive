package middleware

import (
	"net/http"
)

// CacheHeader sets the 'Cache-Control' header on the response to the string parsed into the handler.
func CacheHeader(value string, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", value)
		h(w, r)
	})
}

/* TODO:
func Headers(serverName string) muxie.Wrapper {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", serverName)
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			// TODO: Finish This + Handle When Value In The Config Is Not Set
			// TODO: https://www.owasp.org/index.php/HTTP_Strict_Transport_Security_Cheat_Sheet
			// TODO: https://github.com/srikrsna/security-headers
			// TODO: https://csp.withgoogle.com/docs/strict-csp.html
			// TODO: https://github.com/justinas/nosurf
			// TODO: https://github.com/zeit/next.js/issues/256

			next.ServeHTTP(w, r)
		})
	}
}
*/
