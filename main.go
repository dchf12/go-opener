package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd, err := exec.Command("ls").Output()
	fmt.Printf("ls:%s\n err:%v\n", cmd, err)
}
