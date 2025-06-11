package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Wild CLI - Web UI</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background-color: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #333; border-bottom: 2px solid #007acc; padding-bottom: 10px; }
        .status { background: #e8f5e8; padding: 15px; border-radius: 5px; margin: 20px 0; }
        button { background: #007acc; color: white; border: none; padding: 10px 20px; border-radius: 5px; cursor: pointer; }
        button:hover { background: #005a99; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Wild CLI - Web Interface</h1>
        <div class="status">
            <h3>Status: Running</h3>
            <p>Wild daemon is active and running.</p>
        </div>
        <button onclick="checkStatus()">Check Status</button>
        <div id="result"></div>
    </div>
    <script>
        function checkStatus() {
            fetch('/api/status')
                .then(response => response.json())
                .then(data => {
                    document.getElementById('result').innerHTML = 
                        '<h3>API Response:</h3><pre>' + JSON.stringify(data, null, 2) + '</pre>';
                })
                .catch(error => {
                    document.getElementById('result').innerHTML = 
                        '<h3>Error:</h3><p>' + error + '</p>';
                });
        }
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":    "running",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "0.1.0",
		"uptime":    "running",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}