package handlers

import (
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"testing"
)

type testEnv struct {
	urlStorage  repository.URLStorage
	httpHandler *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	newStorage := repository.Storage{}
	te.urlStorage = newStorage.CreateURLStorage()
	te.httpHandler = NewHTTPHandler(
		te.urlStorage,
	)
	return te
}
