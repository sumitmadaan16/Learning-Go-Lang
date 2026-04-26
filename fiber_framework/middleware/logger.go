package middleware

//
//import (
//	"fmt"
//	"net/http"
//	"time"
//)
//
//// Logger middleware logs the request details and duration of the request processing.
//func Logger(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Record the start time
//		start := time.Now()
//
//		// Log the request details
//		fmt.Printf("Request started - Method: %s, Path: %s\n", r.Method, r.URL.Path)
//
//		// Pass the request to the next handler
//		next.ServeHTTP(w, r)
//
//		// Log the completion time
//		fmt.Printf("Request completed - Method: %s, Path: %s, Duration: %v\n", r.Method, r.URL.Path, time.Since(start))
//	})
//}

// no use of custom logger
