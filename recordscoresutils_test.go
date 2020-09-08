package main

func InitTest() *Server {
	s := Init()
	s.SkipLog = true
	s.SkipIssue = true
	return s
}
