package types

type RequestWriter struct {
	data []byte
}

func NewRequestWriter() *RequestWriter {
	return &RequestWriter{}
}

func (this *RequestWriter) Write(p []byte) (n int, err error) {
	this.data = p
	return 0, nil
}

func (this *RequestWriter) Data() []byte {
	return this.data
}
