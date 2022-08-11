package httpclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	actualHTTPClient := NewHTTPClient()

	assert.Equal(t, http.DefaultClient, actualHTTPClient)
}
