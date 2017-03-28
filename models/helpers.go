package models

import (
	"os/exec"
	"time"
	"strings"
)

func newUUID() string {
	p, _ := exec.Command("uuidgen").Output()
	return string(strings.Trim(string(p), "\n"))
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
