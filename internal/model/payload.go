package model

type DefaultPayload[T any] struct {
	Data T `json:"data"`
}
