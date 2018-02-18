package git

import (
	"io"
	"fmt"
	"regexp"
)

var gitNameRegex = regexp.MustCompile(`(?i)([^.\/]+)\.git$`)

func pipeStdout(rc io.ReadCloser, _ error) {
	buf := make([]byte, 50)
	for {
		n, err := rc.Read(buf)
		fmt.Printf("%s", string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
}