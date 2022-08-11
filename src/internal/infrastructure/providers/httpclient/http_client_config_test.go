package httpclient

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	actualHTTPClient := NewHTTPClient()

	assert.Equal(t, http.DefaultClient, actualHTTPClient)
}
