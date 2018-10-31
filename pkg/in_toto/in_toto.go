package in_toto

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// in_totoClient represent a client for kubesec.io.
type in_totoClient struct {
}

// NewClient returns a new client for kubesec.io.
func NewClient() *in_totoClient {
	return &in_totoClient{}
}

// FIXME: actually return an error
// ScanDefinition scans the provided resource definition.
func (kc *in_totoClient) ScanContainer(imageName string) (*inTotoResult, error) {

	result := inTotoResult{
		Retval: 0,
		Error:  "success",
	}

	err := os.Chdir(imageName)
	if err != nil {
		result.Retval = 128
		result.Error = "Couldn't change directory"
	}

	cmd := exec.Command("in-toto-verify", "-v", "-k", "root_key.pub", "-l", "root.layout")
	_, err = cmd.CombinedOutput()
	if err != nil {
		result.Retval = 127
		result.Error = err.Error()
	}
	return &result, nil
}

// inTotoResult represents a result returned by kubesec.io.
type inTotoResult struct {
	Error  string `json:"error"`
	Retval int    `json:"score"`
}

// Dump writes the result in a human-readable format to the specified writer.
func (r *inTotoResult) Dump(w io.Writer) {
	io.WriteString(w, "-----------------")
	io.WriteString(w, fmt.Sprintf("in-toto analysis score: %v", r.Retval))
	io.WriteString(w, "-----------------")
}
