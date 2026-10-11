package main

import (
	"os"
	"os/exec"
	"syscall"
)


func main()  {
	binary, lookErr:= exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	arg := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, arg, env)
	if execErr != nil {
		panic(execErr)
	}

}