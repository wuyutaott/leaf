package mysql

import (
	"fmt"
	"testing"
)

func TestMySQL(t *testing.T) {
	db, err := Open("root:@tcp(127.0.0.1:3306)/leaf?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Errorf(fmt.Sprintf("open err %v", err))
	}
	err = Close(db)
	if err != nil {
		t.Errorf(fmt.Sprintf("close err %v", err))
	}
}
