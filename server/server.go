package server

import (
	"github.com/utkarsh17ife/goMockPractice/store"
	"github.com/utkarsh17ife/goMockPractice/store/realDb"
)

type Server struct {
	db store.DbStore //store interface
}

func NewServer() *Server{
	srv := &Server{}
	srv.db = &realDb.RealDb{}
	return srv
}

func (a *Server) DoSomeThing() {
	_ = a.db.Insert("some data")
}
