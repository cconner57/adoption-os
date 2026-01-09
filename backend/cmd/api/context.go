package main

import (
	"context"
	"net/http"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

func (app *application) contextSetRequestID(r *http.Request, id string) *http.Request {
	ctx := context.WithValue(r.Context(), requestIDKey, id)
	return r.WithContext(ctx)
}

func (app *application) contextGetRequestID(r *http.Request) string {
	id, ok := r.Context().Value(requestIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

func (app *application) contextSetUser(r *http.Request, userID string) *http.Request {
	ctx := context.WithValue(r.Context(), userIDKey, userID)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) int64 {
	userID, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		return 0
	}
	return userID
}
