package printers

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultResponsePrinter(t *testing.T) {
	expectedResponsePrinter := &defaultResponsePrinter{}

	actualResponsePrinter := NewDefaultResponsePrinter()

	assert.Equal(t, expectedResponsePrinter, actualResponsePrinter)
}

func TestDefaultResponsePrinter_PrintCity_success(t *testing.T) {
	responsePrinter := &defaultResponsePrinter{}

	cityFormatted := givenACityFormatted()
	stringPrinted := cityFormatted + "\n"

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	responsePrinter.PrintCity(cityFormatted)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, stringPrinted, string(out))
}

func givenACityFormatted() string {
	return "Processed city Berlin | Sunny - Sunny"
}
