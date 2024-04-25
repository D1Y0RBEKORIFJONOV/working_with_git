package main

import (
	"os"
	"os/exec"

	"github.com/D1Y0RBEKORIFJONOV/working_with_git.git/git"
)

const PATHS = "/home/diyorbek/go/src/working_with_git/file.txt"

func PrintToFileData(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	file.Write([]byte(data))

	cmd := exec.Command("git", "add", ".")
	cmd.Output()
	cmd = exec.Command("git", "commit", "-m", "appendUsers")
	cmd.Output()
	return nil
}



func main()  {
	strData,err := git.GetUserEmail()
	if err != nil {
		panic(err)
	}
	PrintToFileData(PATHS,strData)

	strData,err = git.GetsUserName()
	if err != nil {
		panic(err)
	}
	PrintToFileData(PATHS, strData)

}
