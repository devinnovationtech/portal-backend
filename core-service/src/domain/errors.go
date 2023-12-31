package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
	// ErrInvalidCredentials
	ErrInvalidCredentials = errors.New("Invalid credentials")
	// ErrDuplicateNIP
	ErrDuplicateNIP = errors.New("NIP already exists")
	// ErrUserIsNotActive
	ErrUserIsNotActive = errors.New("Your account is not active")
	// Forbidden
	ErrForbidden = errors.New("You are not authorized to access this resource")
	// Has Publication
	ErrHasPublication = errors.New("Your requested master data is already used on publication")
	// ErrSlugAlreadyExist
	ErrSlugAlreadyExist = errors.New("Slug already exists")
	// Err for bad request
	ErrBadRequest = errors.New("Bad Request")
)

type ErrResponse struct {
	Message string `json:"message"`
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{Message: err.Error()}
}
