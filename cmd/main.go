package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/D1Y0RBEKORIFJONOV/working_with_git.git/git"
)

const PATH  = "/home/diyorbek/go/src/working_with_git/file.txt"
func PrintToFile(fileName string,data string) error {
	file,err  := os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,0666)
	if err != nil {
		return err
	}
	file.Write([]byte(data))

	cmd := exec.Command("git","add",".")
	cmd.Output()
	cmd = exec.Command("git","commit","-m","appendUsers")
	cmd.Output()
	return nil
}



func main()  {
	str,err := git.GetUserEmail()
	if err != nil {
		panic(err)
	}
	PrintToFile(PATH,str)

	str,err = git.GetUserName()
	if err != nil {
		panic(err)
	}
	PrintToFile(PATH,str)
	fmt.Print("wg")
} 