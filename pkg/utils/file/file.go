package file

import (
	"os"
	"os/exec"
	"strings"
)

func RootPath() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "错误", err
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path, nil
}
