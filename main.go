package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

var (
	openPort string = ":8080"
	err      error
	service  bool
)

func init() {
	tempService := flag.Bool("service", false, "Install the appplication as a serivce.")
	flag.Parse()
	service = *tempService
}

func main() {
	// Install the application as a service, if the flag is set.
	if service {
		installService()
	}
	// Determine how to handle requests.
	http.Handle("/", http.HandlerFunc(handleRequest))
	err = http.ListenAndServe(openPort, nil)
	// Log the error if there is one.
	if err != nil {
		log.Println(err)
	}
}

func handleRequest(httpWriter http.ResponseWriter, r *http.Request) {
	// Set the header to status not found.
	httpWriter.WriteHeader(http.StatusNotFound)
	// Set the content type to application/json.
	httpWriter.Header().Set("Content-Type", "application/json")
	// Set the body to an error message.
	type errorMessage struct {
		Code    int
		Message string
	}
	errorMsg := errorMessage{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
	// Wrap the error in a error object.
	type jsonError struct {
		Error errorMessage
	}
	// The content of the error object.
	jsonReturn := jsonError{
		Error: errorMsg,
	}
	// Marshal the error message to JSON.
	errorJsonMessage, err := json.Marshal(jsonReturn)
	// Log the error if there is one.
	if err != nil {
		log.Println(err)
	}
	// Write the JSON error message.
	httpWriter.Write(errorJsonMessage)
}

// Install the application as a service.
func installService() {
	switch runtime.GOOS {
	case "linux":
		linuxService := `[Unit]
Description=Obfuscated Http Notice
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/obfuscated-http-notice
Restart=on-failure

[Install]
WantedBy=multi-user.target`
		// Create the service file.
		err = os.WriteFile("/etc/systemd/system/obfuscated-http-notice.service", []byte(linuxService), 0644)
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
		// Reload the systemd daemon.
		err = exec.Command("systemctl", "daemon-reload").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
		// Enable the service.
		err = exec.Command("systemctl", "enable", "obfuscated-http-notice.service").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
		// Start the service.
		err = exec.Command("systemctl", "start", "obfuscated-http-notice.service").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
	case "windows":
		// Create a windows service.
		err = exec.Command("sc", "create", "ObfuscatedHttpNotice").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
		// Set the service to start automatically.
		err = exec.Command("sc", "config", "ObfuscatedHttpNotice", "start=", "auto").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
	case "darwin":
		// Create a darwin service.
		err = exec.Command("launchctl", "load", "-w", "/Library/LaunchDaemons/com.obfuscated.http.notice.plist").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
		// Start the service.
		err = exec.Command("launchctl", "start", "com.obfuscated.http.notice").Run()
		// Log the error if there is one.
		if err != nil {
			log.Println(err)
		}
	default:
		log.Fatal(runtime.GOOS, " is not supported.")
	}
}
