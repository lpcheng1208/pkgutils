package middleware

import (
	"net/http"
	"strconv"

	"strings"
)

// HTTP Methods
const (
	GetMethod     = "GET"
	PostMethod    = "POST"
	PutMethod     = "PUT"
	DeleteMethod  = "DELETE"
	OptionsMethod = "OPTIONS"
	PatchMethod   = "PATCH"
	HeadMethod    = "HEAD"
)

// HTTP Headers
const (
	ContentType    = "Content-Type"
	ContentLength  = "Content-Length"
	AcceptEncoding = "Accept-Encoding"
	XCSRFToken     = "X-CSRF-Token"
	Authorization  = "Authorization"
	Accept         = "Accept"
	Origin         = "Origin"
	CacheControl   = "Cache-Control"
	XRequestedWith = "X-Requested-With"
)

// Default values for Options
var (
	defaultAllowOrigins     = []string{"*"}
	defaultAllowHeaders     = []string{ContentType, ContentLength, AcceptEncoding, XCSRFToken, Authorization, Accept, Origin, CacheControl, XRequestedWith}
	defaultAllowMethods     = []string{GetMethod, PostMethod, PutMethod, DeleteMethod, PatchMethod, HeadMethod} // Not managing OPTIONS as default method in order to manage it individually
	defaultAllowCredentials = true
)

// Options for Handler
type Options struct {
	AllowOrigins     []string
	AllowHeaders     []string
	AllowMethods     []string
	AllowCredentials bool
}

// CorsHandler for setting headers on every managed request
func CorsHandler(options Options) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		// Setting the default origins in case not specified
		if options.AllowOrigins == nil {
			options.AllowOrigins = defaultAllowOrigins
		}
		// Setting the default headers in case not specified
		if options.AllowHeaders == nil {
			options.AllowHeaders = defaultAllowHeaders
		}
		// Setting the default methods in case not specified
		if options.AllowMethods == nil {
			options.AllowMethods = defaultAllowMethods
		}
		// Request managing func
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(options.AllowOrigins) > 0 {
				w.Header().Set("Access-Control-Allow-Origin", strings.Join(options.AllowOrigins, " "))
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}

			if len(options.AllowHeaders) > 0 {
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(options.AllowHeaders, ","))
			} else {
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(defaultAllowHeaders, ","))
			}

			if len(options.AllowMethods) > 0 {
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(options.AllowMethods, ","))
			} else {
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(defaultAllowMethods, ","))
			}

			if options.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(defaultAllowCredentials))
			}

			/** OPTIONS Method returns no content status, this is important for example
			when requesting server when AngularJS Resources in order to avoid OPTIONS Request error
			*/
			if r.Method == OptionsMethod {
				w.WriteHeader(http.StatusNoContent)
			}

			next.ServeHTTP(w, r)
		})
	}
}
