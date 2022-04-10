package compiler

type InputBuffer struct {
	Buffer       string
	BufferLength int
	InputLength  int
}

func NewInputBuffer() InputBuffer {
	newInputBuffer := InputBuffer{}
	newInputBuffer.Buffer = ""
	newInputBuffer.BufferLength = 16
	newInputBuffer.InputLength = 16
	return newInputBuffer
}
