package auth

import (
	"database/sql"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/basic/db"
)

func (s *service) QueryUserByName(loginName string) (ret *auth.QueryUserData, err error) {
	queryString := `SELECT user_id,login_name,login_password FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &auth.QueryUserData{}
	err = o.QueryRow(queryString, loginName).Scan(&ret.UserID, &ret.LoginName, &ret.LoginPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			// log.Warn("[QueryUserByName] 查询数据不存在,loginName:"+loginName, zap.Any("err", err))
			return nil, err
		}
		// log.Error("[QueryUserByName] 查询数据失败,loginName:"+loginName, zap.Any("err", err))
		return
	}
	return

}
