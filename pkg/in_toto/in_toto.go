package in_toto

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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
		Output: "",
	}

	oldWd, err := os.Getwd()
	if err == nil {

		dir := imageName
		if strings.Contains(dir, ":") {
			dir = dir[:strings.LastIndexByte(dir, ':')]
		}
		err := os.Chdir(dir)
		if err != nil {
			result.Retval = 128
			result.Error = "Couldn't change directory"
		} else {
			cmd := exec.Command("in-toto-verify", "-v", "-k", "root_key.pub", "-l", "root.layout")
			stdoutStderr, execErr := cmd.CombinedOutput()
			result.Output = fmt.Sprintf("%s", stdoutStderr)
			if execErr != nil {
				result.Retval = 127
				result.Error = execErr.Error()
			}
			err = os.Chdir(oldWd)
			if err != nil {
				result.Retval = 128
				result.Error = "Couldn't change to old directory"
			}
		}
	}
	return &result, nil
}

// inTotoResult represents a result returned by kubesec.io.
type inTotoResult struct {
	Error  string `json:"error"`
	Retval int    `json:"score"`
	Output string `json:"output"`
}

// Dump writes the result in a human-readable format to the specified writer.
func (r *inTotoResult) Dump(w io.Writer) {
	io.WriteString(w, "-----------------")
	io.WriteString(w, fmt.Sprintf("in-toto analysis score: %v", r.Retval))
	io.WriteString(w, "-----------------")
}
