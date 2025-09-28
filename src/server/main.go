package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"power4web/src/server/handlers"
	"runtime"
)

func main() {
	ClearTerminal()
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("src/public"))))
	http.HandleFunc("/play", handlers.GameHandler)
	http.HandleFunc("/placeLine", handlers.GameLinePlayHandler)
	http.HandleFunc("/newGame", handlers.GameNewPartyHandler)
	http.HandleFunc("/changeSlot", handlers.ChangeSlotHandler)
	http.HandleFunc("/selectPlayer", handlers.SelectPlayerHandler)
	http.HandleFunc("/", handlers.LandingPageHandler)

	fmt.Println("Server is listening on 127.0.0.1:3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
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
