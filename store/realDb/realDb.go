package realDb

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type RealDb struct {
}

func (r *RealDb) Insert(data interface{}) bool {

	fmt.Println("Inserted data in real db")
	return true
}

//for testing
func (r *RealDb) On(methodName string, arguments ...interface{}) *mock.Call {
	return &mock.Call{}
}
func (r *RealDb) AssertExpectations(t mock.TestingT) bool {
	return true
}
