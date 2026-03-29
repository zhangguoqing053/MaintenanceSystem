package service

import (
	"MaintenanceSystem/repository"
	"MaintenanceSystem/pkg"
	"fmt"

)

type UserService struct {
	repo  *repository.UserDB
}

func InitServiceDB(repo *repository.UserDB) *UserService{
	return  &UserService{repo:repo}
}


//用户登录
func (s *UserService)Login(uname, pwd string)(string,error){
    u,err :=  s.repo.GetByUsername(uname)
	if err != nil {
		return  "",err
	}
	if !pkg.CheckPassword(u.Password,pwd) {
        return "",fmt.Errorf("账号密码错误")
	}
	return pkg.GenerateToken(u.ID,uname)
}