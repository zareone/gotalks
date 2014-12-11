type ResponseWriter interface {
	// Header returns the header map that will be sent by WriteHeader.
	Header() Header

	// Write writes the data to the connection as part of an HTTP reply.
	Write([]byte) (int, error)

	// WriteHeader sends an HTTP response header with status code.
	WriteHeader(int)
}
