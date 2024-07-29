# Port Scanner
**Demostration Video :** [Youtube Link](https://www.youtube.com/watch?v=TCB7zSeHnH8)
- A personal project to scan open ports on a given IP/Domain address using TCP connect scan.

## Overview
Port Scanner is a simple web application that takes an IP/Domain address, start port, and end port as input and scans for open ports using TCP connect scan. The application displays a list of open ports found during the scan.

### Features
- Scans open ports on a given IP/Domain address
- Takes start and end ports as input for scanning
- Uses TCP connect scan for port scanning
- Displays a list of open ports found during the scan

### Usage
- Run the application on a web server (e.g., http://localhost:8080)
- Enter the IP/Domain address, start port, and end port in the input fields
- Click the "Submit" button to initiate the port scan
- The application will display a list of open ports found during the scan
### Technical Details
- Backend:
  - Written in Go programming language (Golang)
  - Uses Go's templating engine for dynamic HTML rendering
  - Employs Go concurrency concepts for efficient port scanning:
    - Goroutines for concurrent port scanning
    - Channels for communication between goroutines
    - Sync.WaitGroup for synchronization and waiting for goroutines to  finish
- Frontend:
    - Built using HTML and Bootstrap for responsive design
    - Utilizes HTMX for dynamic client-side rendering and AJAX requests
    - Provides a user-friendly interface for inputting scan parameters and displaying results
### Future Enhancements
- Add support for other scanning techniques (e.g., SYN scan, UDP scan)
- Implement rate limiting to prevent abuse
  

### Acknowledgments
Inspired by port scanning tools like ncap.org

Thank You.
