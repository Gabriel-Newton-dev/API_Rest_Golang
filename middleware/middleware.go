package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content_type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ele vai receber a solicitacao do nosso roteador,
// para evitar duplicidade de código.
