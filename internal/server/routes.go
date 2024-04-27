package server

func (s *FiberServer) RegisterRoutes() {
  // Sends a hello world message
	s.App.Get("/", s.HelloWorldHandler)
}
