package main

import "net/http"

type ResponseWrapper struct {
	statusCode int
	realWriter http.ResponseWriter
}

func NewResponseWrapper(w http.ResponseWriter) *ResponseWrapper {
	return &ResponseWrapper{
		statusCode: 200,
		realWriter: w,
	}
}

func (w *ResponseWrapper) Header() http.Header {
	return w.realWriter.Header()
}
func (w *ResponseWrapper) Write(buf []byte) (int, error) {
	return w.realWriter.Write(buf)
}
func (w *ResponseWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.realWriter.WriteHeader(statusCode)
}
