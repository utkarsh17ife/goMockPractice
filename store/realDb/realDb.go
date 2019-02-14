package realDb

import "fmt"

type RealDb struct {
}

func (r *RealDb) Insert(data interface{}) bool {

	fmt.Println("Inserted data in real db")
	return true
}
