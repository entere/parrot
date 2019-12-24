package auth

import (
	"github.com/entere/parrot/basic/db"
)

func (s *service) UpdatePassword(password string, userID string) (err error) {
	updateSQL := `UPDATE auth_passwords SET login_password  = ? WHERE user_id=?`
	o := db.GetDB()
	// 更新
	_, err = o.Exec(updateSQL, password, userID)
	if err != nil {
		return err
	}
	return nil
}
