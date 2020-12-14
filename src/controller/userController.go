// @File: userController
// @Date: 2020/12/6 20:51
// @Author: 安红豆
// @Description: 用户相关的控制器
package controller

import (
	"fmt"
	"github.com/Ormissia/ormissia_go/src/model"
	"github.com/Ormissia/ormissia_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (w *UserController) Register(ctx *gin.Context) {
	requestUser := model.User{}
	//从请求中获取用户参数
	err := ctx.Bind(&requestUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//获取参数
	username := requestUser.Username
	password := requestUser.Password
	email := requestUser.Email

	//用户名密码
	if len(username) == 0 || len(password) == 0 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, "输入错误", nil)
		return
	}
	//邮箱格式验证
	if !util.EmailRegexp(email) {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, "输入错误", nil)
		return
	}
	//判断用户是否存在
	existUser, _ := model.SelectUserInfoByUsername(username)
	if existUser.ID != 0 {
		util.Response(ctx, http.StatusUnprocessableEntity, 422, "该用户已存在", nil)
		return
	}

	//往数据库中插入该注册用户
	result := model.InsertUser(requestUser)
	if result != nil {
		//插入出错
	}
	//插入成功
	util.Success(ctx, nil, "注册成功")
	return
}

func (w *UserController) Login(ctx *gin.Context) {
	requestUser := model.User{}
	//Bind在请求过程中，如果参数错误会直接抛异常返回400状态
	err := ctx.Bind(&requestUser)
	//ShouldBind在请求过程中，对参数检测不做处理
	//ctx.ShouldBind(&requestUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.SelectUserInfoByUserId(requestUser.ID)
	fmt.Println(user)

	// 没有数据库的用下面这个方法：这里先写死账号和密码  有数据库的要从数据库中获取
	if requestUser.Username != "123" || requestUser.Password != "123" {
		ctx.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 1001
			"code":    1001,
			"message": "failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		// 登录失败返回code 1000
		"code":    1000,
		"message": "success",
	})
	return
}
