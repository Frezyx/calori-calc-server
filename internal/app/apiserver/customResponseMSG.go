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
	errNotFoundDiet             = errors.New("diet id is not found")
	msgUserDeleted              = "user is deleted"
	msgUserProductDeleted       = "user product is deleted"
	msgDietCreate               = "diet is created"
	msgAuthorized               = "user is authorized"
	msgChangesSave              = "changes is saved"
	msgDietDeleted              = "diet is deleted"
	msgDietUpdated              = "diet is updated"
)
