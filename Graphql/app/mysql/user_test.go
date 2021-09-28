package mysql

import (
	"fmt"
	"testing"
)

func TestselectUserByID(t *testing.T) {

}
func TestName(t *testing.T) {
	dao := CreateUserDao()
	res, err := dao.selectUserByID(1)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, re := range res {
		fmt.Printf("%v", re)
	}
}
