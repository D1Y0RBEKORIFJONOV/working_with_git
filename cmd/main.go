package main

import (
	"os"
	"os/exec"

	"github.com/D1Y0RBEKORIFJONOV/working_with_git.git/git"
)


func PrintToFile(fileName string,data string) error {
	err := os.WriteFile(fileName,[]byte(data),777)
	if err != nil {
		return err
	}

	cmd := exec.Command("git","add",".")
	cmd.Output()
	
	return nil
}





func main()  {
}