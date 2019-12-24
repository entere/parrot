/**
 * @Author: entere
 * @Description:
 * @File:  user_get.go
 * @Version: 1.0.0
 * @Date: 2019/12/19 21:33
 */

package user

import (
	"github.com/entere/parrot/basic/db"
	user "github.com/entere/parrot/user-srv/proto/user"
)

func (s service) GetUser(userID string) (ret *user.GetUserData, err error) {
	queryString := `SELECT user_id,avatar_url,mobile FROM users WHERE user_id = ?`
	o := db.GetDB()
	ret = &user.GetUserData{}
	err = o.QueryRow(queryString, userID).Scan(&ret.UserID, &ret.AvatarURL, &ret.Mobile)
	return
}
