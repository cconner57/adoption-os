package main

import (
	"context"
	"net/http"
)

type contextKey string

const requestIDContextKey = contextKey("requestID")

func (app *application) contextSetRequestID(r *http.Request, id string) *http.Request {
	ctx := context.WithValue(r.Context(), requestIDContextKey, id)
	return r.WithContext(ctx)
}

func (app *application) contextGetRequestID(r *http.Request) string {
	id, ok := r.Context().Value(requestIDContextKey).(string)
	if !ok {
		return ""
	}
	return id
}
