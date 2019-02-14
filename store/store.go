package store

import (
	"github.com/stretchr/testify/mock"
)

type DbStore interface {
	Insert(data interface{}) bool

	//for testing purpose
	On(methodName string, arguments ...interface{}) *mock.Call
	AssertExpectations(t mock.TestingT) bool
}
