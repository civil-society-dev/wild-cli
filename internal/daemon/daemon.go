package daemon

import (
	"fmt"
	"net/http"

	"wild-cli/internal/web"
)

func Start() error {
	fmt.Println("Initializing daemon...")
	
	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/api/status", web.StatusHandler)
	
	port := ":5065"
	fmt.Printf("Starting web server on http://localhost%s\n", port)
	
	return http.ListenAndServe(port, nil)
}