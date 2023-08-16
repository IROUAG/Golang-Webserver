
# Go Web Server

A lightweight web server written in Go. The server serves static files and handles specific routes for processing form data and displaying a welcome message.

## Features
- Serves static files from the `./static` directory.
- Handles the `/hello` route to display a "Hello!" message.
- Handles the `/form` route to process form data and display the form values.

## Installation and Usage
1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the directory containing the Go files.
4. Run the command `go run main.go` to start the server.
5. Visit `http://localhost:8080/` in your browser to access the server.
6. You can access the `/hello` and `/form` routes by visiting `http://localhost:8080/hello` and `http://localhost:8080/form` respectively.

## Customization
- You can modify the paths for static files and routes as needed.
- Extend the server by adding more routes and functionalities as per your requirements.

## License
This project is open-source and available to everyone. Feel free to use, modify, and distribute as you see fit.
