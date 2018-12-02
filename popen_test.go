// +build !windows

package process

import (
	"testing"
)

func TestPopen(t *testing.T) {
	Popen("echo 123", "r")
}
