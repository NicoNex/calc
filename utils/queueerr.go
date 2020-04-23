package utils

type EmptyQueue struct {
	msg string
}

func NewEmptyQueue() error {
	return EmptyQueue{"error: empty queue"}
}

func (e EmptyQueue) Error() string {
	return e.msg
}
