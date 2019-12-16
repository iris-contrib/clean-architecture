package controller

import (
	"github.com/iris-contrib/clean-architecture/domain/model"
	"github.com/iris-contrib/clean-architecture/usecase/interactor"

	"github.com/kataras/iris/v12"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(iris.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx iris.Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	_, err = ctx.JSON(u)
	return err
}
