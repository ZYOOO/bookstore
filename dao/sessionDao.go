package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddSession 添加session
func AddSession(sess *model.Session) error {
	sql := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sql, sess.SessionID, sess.Username, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sql := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sql, sessID)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessID string) (*model.Session, error) {
	sql := "select session_id,username,user_id from sessions where session_id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(sessID)
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.Username, &sess.UserID)
	return sess, nil
}
