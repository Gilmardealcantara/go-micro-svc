package http

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/data"
)

// ResponseWriterWrapper struct is used to log the response
type ResponseWriterWrapper struct {
	w          http.ResponseWriter
	body       *bytes.Buffer
	statusCode *int
	account    *data.Account
	error      *error
}

// NewResponseWriterWrapper static function creates a wrapper for the http.ResponseWriter
func NewResponseWriterWrapper(w http.ResponseWriter) ResponseWriterWrapper {
	return ResponseWriterWrapper{
		w:          w,
		body:       new(bytes.Buffer),
		statusCode: new(int),
		account:    new(data.Account),
		error:      new(error),
	}
}

func (rww ResponseWriterWrapper) Write(buf []byte) (int, error) {
	rww.body.Write(buf)
	return rww.w.Write(buf)
}

// Header function overwrites the http.ResponseWriter Header() function
func (rww ResponseWriterWrapper) Header() http.Header {
	return rww.w.Header()

}

// WriteHeader function overwrites the http.ResponseWriter WriteHeader() function
func (rww ResponseWriterWrapper) WriteHeader(statusCode int) {
	*rww.statusCode = statusCode
	rww.w.WriteHeader(statusCode)
}

func (rww ResponseWriterWrapper) Code() int {
	return *rww.statusCode
}

func (rww ResponseWriterWrapper) Body() *bytes.Buffer {
	return rww.body
}

func (rww ResponseWriterWrapper) WriteError(err error) {
	*rww.error = err
}

func (rww ResponseWriterWrapper) Error() error {
	return *rww.error
}

func (rww ResponseWriterWrapper) WriteAccount(acc data.Account) {
	*rww.account = acc
}

func (rww ResponseWriterWrapper) Account() *data.Account {
	return rww.account
}

func (rww ResponseWriterWrapper) String() string {
	var buf bytes.Buffer

	buf.WriteString("\n---Response Debug:---\n")

	buf.WriteString("Headers:\n")
	for k, v := range rww.Header() {
		buf.WriteString(fmt.Sprintf("%s: %v", k, v))
	}

	buf.WriteString(fmt.Sprintf("\n\nStatus Code: %d", *rww.statusCode))

	buf.WriteString("\n\nBody:\n")
	buf.WriteString(rww.body.String())
	buf.WriteString("\n-----\n")

	return buf.String()
}

func WriteHttpError(w http.ResponseWriter, err error) {
	if rww, ok := w.(ResponseWriterWrapper); ok {
		rww.WriteError(err)
	}
}
