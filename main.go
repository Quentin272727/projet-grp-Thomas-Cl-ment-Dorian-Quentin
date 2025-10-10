package main

import (
	"os"
	"os/exec"
	"power4web/src/server"
	"runtime"
)

func main() {
	ClearTerminal()
	println("Starting HTTP Server...")
	server.StartHttpServer()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
