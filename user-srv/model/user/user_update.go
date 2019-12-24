package user

import (
	"github.com/entere/parrot/basic/db"
	user "github.com/entere/parrot/user-srv/proto/user"
)

func (s service) UpdateUser(req *user.UpdateUserRequest) (err error) {
	updateSQL := `UPDATE users SET avatar_url  = ?, mobile = ? WHERE user_id=?`
	o := db.GetDB()
	// 更新
	_, err = o.Exec(updateSQL, req.AvatarURL, req.Mobile, req.UserID)
	if err != nil {
		return
	}
	return
}
