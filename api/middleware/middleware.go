package middleware

import "net/http"

type Middleare func(http.Handler) http.Handler
