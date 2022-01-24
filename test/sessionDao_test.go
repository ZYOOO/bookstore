package test

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "213456",
		Username:  "ss",
		UserID:    5,
	}
	dao.AddSession(sess)
}

func TestDeleteSession(t *testing.T) {
	dao.DeleteSession("213456")
}

func TestGetSession(t *testing.T) {
	sess, err := dao.GetSession("b16f65b8-da56-4792-7200-af8fdb456606")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sess)
}
