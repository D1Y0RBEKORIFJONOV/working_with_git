package git

import (
	"os/exec"
)

func GetUserEmail() (string,error) {
	cmd := exec.Command("git", "config", "user.email")
	str, err := cmd.Output()
	if err != nil {
		return "",err
	}
	return string(str),err
}

func GetsUserName() (string,error) {
	cmd := exec.Command("git", "config", "user.name")
	str, err := cmd.Output()
	if err != nil {
		return "",err
	}
	return string(str),err
}
