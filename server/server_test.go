package server

import (
	"testing"

	"github.com/utkarsh17ife/goMockPractice/store/mockDb"
)

func TestDoSomeThing(t *testing.T) {

	srv := &Server{}
	srv.db = &mockDb.MockDb{}

	srv.db.On("Insert", "some data").Return(true)

	srv.DoSomeThing()

	srv.db.AssertExpectations(t)

}
