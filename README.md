# Cloud Image Server

A Go-based web application that serves random images from a local directory, built for cloud computing deployment.

## Features

- **Random Image Selection**: Displays a random set of images on each page load
- **Base64 Encoding**: Images are encoded and embedded directly in HTML
- **Responsive Design**: Bootstrap-powered UI that works on all devices

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, Bootstrap 5.3
- **Templates**: Go HTML templates
- **Styling**: Custom CSS with Bootstrap

## Lab Purpose

This project is part of a Cloud Computing laboratory exercise, demonstrating web application deployment in cloud environments.

## Getting Started

To run the server:

```bash
go run cmd/main.go <port>

# Examples
go run cmd/main.go 8080    # Run on port 8080
go run cmd/main.go 3000    # Run on port 3000
go run cmd/main.go         # Default port 8000
```

## Developers

- Jose David Amaya
- Daniel Felipe Correa

Cloud Computing 2025-2
