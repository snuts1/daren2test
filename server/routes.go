package server

import "net/http"

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	err := s.Templates.ExecuteTemplate(w, "placeholder.html", nil) // Assuming you have an index.html
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleWeather(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This page is for do Weather"))
}

func (s *Server) handlePayback(w http.ResponseWriter, r *http.Request) {
	err := s.Templates.ExecuteTemplate(w, "payback.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) registerRoutes() {
	s.Router.HandleFunc("/", s.handleIndex)
	s.Router.HandleFunc("/weather", s.handleWeather)
	s.Router.HandleFunc("/payback", s.handlePayback)
}
