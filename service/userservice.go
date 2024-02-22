package service

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// FindUserByNameAndPassoword
// @Sumary 用户登陆
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json {"code", "message"}
// @Router /user/findUserByNameAndPassoword [post]
func FindUserByNameAndPassoword(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	user := models.FindUserByName(name)

	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)

	if !flag {
		c.JSON(200, gin.H{
			"message": "密码不正确 ",
		})
		return
	}

	pwd := utils.MakePassowrd(password, user.Salt)
	data := models.FindUserByNameAndPassoword(name, pwd)

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Sumary 新增用户
// @Tags 用户模块
// @Success 200 {string} json {"code", "message"}
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}

	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册",
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	user.PassWord = utils.MakePassowrd(password, salt)
	models.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @Sumary 删除用户
// @Tags 用户模块
// @Success 200 {string} json {"code", "message"}
// @param id query string false "id"
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))

	user.ID = uint(id)
	models.DeleteUser(user)

	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Sumary 修改用户
// @Tags 用户模块
// @Success 200 {string} json {"code", "message"}
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))

	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
			"data":    err,
		})
		return
	}

	models.UpdateUser(user)

	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})
}
