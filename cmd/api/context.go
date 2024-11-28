package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mahesh-singh/greenlight/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {

	ctx := context.WithValue(r.Context(), userContextKey, user)

	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {

	value := r.Context().Value(userContextKey)

	if value == nil {
		return data.AnonymousUser
	}

	user, ok := value.(*data.User)

	if !ok {
		panic(fmt.Sprintf("missing user value in the request context: %T", value))
	}
	return user
}
