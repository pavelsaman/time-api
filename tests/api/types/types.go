package test_types

import "net/http"

type Handler struct {
	Url     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

type Handlers struct {
	Handlers []*Handler
}
