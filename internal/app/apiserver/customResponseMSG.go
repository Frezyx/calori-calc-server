package apiserver

import "errors"

var (
	errNotAuthenticated         = errors.New("user is not authenticated")
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errSearch                   = errors.New("search error")
	errIncorrectPassword        = errors.New("incorrect password")
	errNotFoundUser             = errors.New("user with this id is not found")
	errNotFoundDate             = errors.New("date id is not found")
	errNotFoundUserProduct      = errors.New("user product with this id is not found")
	errEmptyUserProductList     = errors.New("user product list is empty")
	msgUserDeleted              = "user is deleted"
	msgUserProductDeleted       = "user product is deleted"
	msgAuthorized               = "user is authorized"
	msgChangesSave              = "changes is saved"
)
