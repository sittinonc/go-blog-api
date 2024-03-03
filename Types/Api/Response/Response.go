package Response

type OperationResponse[T any] struct {
	Success bool
	Message string
	Data    T
}
