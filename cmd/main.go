package main

import (
	"os"
	"os/exec"

	"github.com/D1Y0RBEKORIFJONOV/working_with_git.git/git"
)


func PrintToFile(fileName string,data string) error {
	err := os.WriteFile(fileName,[]byte(data),0344)
	if err != nil {
		return err
	}

	cmd := exec.Command("git","add",".")
	cmd.Output()
	cmd = exec.Command("git","commit","-m","appendUsers")
	cmd.Output()
	return nil
}





func main()  {
	str,err := git.GetUserName()
	if err != nil {
		panic(err)
	}
	PrintToFile("/home/diyorbek/go/src/working_with_git/file.txt",str)
}