package in_toto

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	"fmt"
	"io"
	//"io/ioutil"
	//"mime/multipart"
	//"net/http"
)

// in_totoClient represent a client for kubesec.io.
type in_totoClient struct {
}

// NewClient returns a new client for kubesec.io.
func NewClient() *in_totoClient {
	return &in_totoClient{}
}

// ScanDefinition scans the provided resource definition.
func (kc *in_totoClient) ScanContainer() (*inTotoResult, error) {

    result := inTotoResult{
        Error: "error",
        Retval: 0,
    }

	return &result, nil
}

// inTotoResult represents a result returned by kubesec.io.
type inTotoResult struct {
	Error   string `json:"error"`
	Retval  int    `json:"score"`
}

// Dump writes the result in a human-readable format to the specified writer.
func (r *inTotoResult) Dump(w io.Writer) {
	io.WriteString(w, "-----------------")
	io.WriteString(w, fmt.Sprintf("in-toto analysis score: %v", r.Retval))
	io.WriteString(w, "-----------------")
}
