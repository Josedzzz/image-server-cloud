package model

// Server configuration
type Config struct {
	Port     string
	Hostname string
	Theme    string
}

// Data passed to the template
type PageData struct {
	Title    string
	Hostname string
	Theme    string
	Images   []ImageData
}

// Image with base64 encoding
type ImageData struct {
	Name string
	Data string
}
