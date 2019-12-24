package auth

import (
	"errors"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/basic/db"
)

func (s *service) LoginByName(req *auth.LoginByNameRequest) (ret *auth.LoginByNameData, err error) {
	queryString := `SELECT user_id,login_name,login_password FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &auth.LoginByNameData{}
	err = o.QueryRow(queryString, req.LoginName).Scan(&ret.UserID, &ret.LoginName, &ret.LoginPwd)
	if err != nil {
		return nil, err
	}

	if ret.LoginPwd != req.LoginPwd {
		return nil, errors.New("密码错误")
	}
	return ret, nil

}
