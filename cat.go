package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	cat, err := ioutil.ReadFile("cat.txt")
	if err != nil {
		panic(err)
	}
	cat2, err := ioutil.ReadFile("cat2.txt")
	for true {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Print(string(cat))
		time.Sleep(time.Second)
		cmd = exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Print(string(cat2))
		time.Sleep(time.Second)
	}
}
