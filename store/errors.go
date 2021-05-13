package store

// ErrNotFound indicates that a model was not found
type ErrNotFound struct {
	modelName string
}

func NewErrNotFound(modelName string) *ErrNotFound {
	return &ErrNotFound{modelName}
}

func (e *ErrNotFound) Error() string {
	return e.modelName + " not found"
}
