package rest_error

import "net/http"

var ErrSomethingHappened = New(http.StatusInternalServerError, "somethingHappen")
var ErrCodeDidNotSent = New(http.StatusBadRequest, "codeDidNotSent")
var ErrCodesNotEqual = New(http.StatusBadRequest, "codesNotEqual")
var ErrCodeMaxAttempts = New(http.StatusBadRequest, "codeMaxAttempts")
var ErrUserExists = New(http.StatusBadRequest, "userExists")
var ErrUserNotExists = New(http.StatusBadRequest, "userNotExists")
var ErrIncorrectPassword = New(http.StatusBadRequest, "incorrectPassword")
var ErrNotAuthorized = New(http.StatusUnauthorized, "notAuthorized")
