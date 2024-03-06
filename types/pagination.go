package types

type PaginationResult[T any] struct {
	Items *[]T `json:"items"`
	Total uint `json:"total"`
}
