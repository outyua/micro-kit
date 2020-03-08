package service

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	"micro-kit/apps/srv/admin/api/auth"
)

type Auth struct{}

func (a *Auth) Login(c context.Context, req *auth.LoginReq, res *auth.LoginRes) error {
	logger.Error("auth req::::", req.Username, req.Password)
	res = &auth.LoginRes{
		IsExist:   true,
		Available: true,
		Token:     &auth.Token{Token: "000000000000"},
	}
	return nil
}

func (a *Auth) Logout(c context.Context, t *auth.Token, r *auth.OptResult) error {

	logger.Error("token", t.Token)

	r.Ok = true
	return nil
}

func (a *Auth) Info(c context.Context, t *auth.Token, r *auth.AdminInfo) error {
	logger.Error("token", t.Token)

	r = &auth.AdminInfo{
		Uid:           10086,
		Name:          "outyua",
		Avatar:        "https://media.ifanrusercontent.com/tavatar/28/d3/28d39bd75da9ef227188eae4fd73777cc8b30bfc.jpg",
		LastLoginTime: 1583682476,
		RegisterTime:  1583682476,
		Roles:         []string{"admin", "editor", "guest"},
		Status:        1,
	}
	return nil
}
