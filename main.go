package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/manki/flickgo"
)

func clientAuth(client *flickgo.Client, frob string) (string, error) {
	u := client.AuthDesktopURL(flickgo.DeletePerm, frob)

	var err error
	browser := "xdg-open"
	args := []string{u}
	if runtime.GOOS == "windows" {
		browser = "rundll32.exe"
		args = []string{"url.dll,FileProtocolHandler", u}
	} else if runtime.GOOS == "darwin" {
		browser = "open"
		args = []string{u}
	} else if runtime.GOOS == "plan9" {
		browser = "plumb"
	}
	fmt.Println("Open this URL and enter PIN.")
	fmt.Println(u)
	browser, err = exec.LookPath(browser)
	if err == nil {
		cmd := exec.Command(browser, args...)
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			return "", fmt.Errorf("failed to start command: %v", err)
		}
	}

	fmt.Print("Hit Enter")
	var b [1]byte
	os.Stdin.Read(b[:])

	token, _, err := client.GetToken(frob)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}
	return token, nil
}

func main() {
	client := flickgo.New("9e076e33fec7988912ece2abe5c5990a", "7f445faa0c241cf1", http.DefaultClient)
	frob, err := client.GetFrob()
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("GOUPLOADR_TOKEN")
	if token == "" {
		token, err = clientAuth(client, frob)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("GOUPLOADR_TOKEN=" + token)
	}
	client.AuthToken = token

	for _, arg := range os.Args[1:] {
		var name, file string
		sa := strings.SplitN(arg, "#", 2)
		file = sa[0]
		if len(sa) == 1 {
			name = filepath.Base(file)
		} else {
			name = sa[1]
		}
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		tid, err := client.Upload(name, b, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Uploaded: %v", tid)
	}
}
