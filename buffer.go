package sqlitego

type InputBuffer struct {
	Buffer       string
	BufferLength int
	InputLength  int
}

func NewInputBuffer() InputBuffer {
	newInputBuffer := InputBuffer{}
	newInputBuffer.Buffer = ""
	newInputBuffer.BufferLength = 0
	newInputBuffer.InputLength = 0
	return newInputBuffer
}
