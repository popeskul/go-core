package api

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
)

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request_id", rand.Intn(1_000_000))
		newR := r.WithContext(ctx)

		b, _ := httputil.DumpRequest(newR, true)
		fmt.Printf("%+v", string(b))

		next.ServeHTTP(w, newR)
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("request_id").(int)
		log.Println(r.Method, r.RemoteAddr, r.RequestURI, id)
		next.ServeHTTP(w, r)
	})
}
