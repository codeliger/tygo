package main

type Error string

func (e Error) Error() string {
	return string(e)
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

// Authorization errors.
const (
	ErrCycleWrongBrand             = Error("cycle does not belong to users brand")
	ErrColorwaySizeAssetWrongBrand = Error("colorway_size_asset does not belong to users brand")
)
