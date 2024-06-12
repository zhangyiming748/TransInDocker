package util

import "testing"

func TestIsExistCmd(t *testing.T) {
	cmd := "ffmpeg"
	cmd2 := "trans"
	cmd3 := "translate-shell"
	if IsExistCmd(cmd, cmd2, cmd3) {
		t.Log("yes")
	} else {
		t.Log("no")
	}
}
