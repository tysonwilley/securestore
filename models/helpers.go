package models

import (
	"os/exec"
	"time"
)

func newUUID() string {
	p, _ := exec.Command("uuidgen").Output()
	return string(p)
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
