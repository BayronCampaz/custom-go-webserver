package webserver

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (server *Server) Listen() error {
	http.Handle("/", server.router)
	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := server.router.rules[method]
	if !exist {
		server.router.rules[method] = make(map[string]http.HandlerFunc)
	}
	server.router.rules[method][path] = handler
}

func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}
