package auth

import (
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/basic/db"
	"github.com/micro/go-micro/util/log"
)

func (s *service) QueryUserByName(loginName string) (ret *auth.QueryUserData, err error) {
	log.Info(loginName)
	queryString := `SELECT user_id,login_name,login_password FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &auth.QueryUserData{}
	err = o.QueryRow(queryString, loginName).Scan(&ret.UserID, &ret.LoginName, &ret.LoginPwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
