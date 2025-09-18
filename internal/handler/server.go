package handler

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"image-server/internal/model"
	"image-server/internal/service"
)

type Server struct {
	port     string
	imageSvc service.ImageService
	template *template.Template
}

func NewServer(port string) *Server {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Warning: Error loading template: %v", err)
	}

	return &Server{
		port:     port,
		imageSvc: service.NewImageService(),
		template: tmpl,
	}
}

func (s *Server) Start() error {
	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	// Configuration
	config := model.Config{
		Port:     s.port,
		Hostname: hostname,
		Theme:    "Animals",
	}

	// Setup routes
	http.HandleFunc("/", s.indexHandler(config))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start server - bind to all interfaces
	addr := "0.0.0.0:" + s.port
	log.Printf("Server starting on http://%s", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) indexHandler(config model.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		images, err := s.imageSvc.LoadRandomImages("static/images", 4)
		if err != nil {
			http.Error(w, "Error loading images: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := model.PageData{
			Title:    "Image Server",
			Hostname: config.Hostname,
			Theme:    config.Theme,
			Images:   images,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := s.template.Execute(w, data); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Error generating page", http.StatusInternalServerError)
		}
	}
}
