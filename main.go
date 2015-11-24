package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/reboot", serveTemplate)
	http.ListenAndServe(":5188", nil)
}

func reboot() {
	cmd := "bash"
	fmt.Println("reboot-factory")
	args := []string{"bin/factory-reboot"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PRESSED")
	reboot()
	http.FileServer(http.Dir("."))
}
