package git

import (
	"os/exec"
)

<<<<<<< HEAD
=======
func GetUserEmail() (string,error) {
	cmd := exec.Command("git", "config", "user.email")
	str, err := cmd.Output()
	if err != nil {
		return "",err
	}
	return string(str),err
}

>>>>>>> feature/add-user-emal
func GetUserName() (string,error) {
	cmd := exec.Command("git", "config", "user.name")
	str, err := cmd.Output()
	if err != nil {
		return "",err
	}
	return string(str),err
}
