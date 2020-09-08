package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"mdbook/common"
	"mdbook/models"
	"mdbook/utils"
	"regexp"
	"strings"
	"time"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) Register() {
	if c.Ctx.Input.IsPost() {
		//account=&password1=&password2=&nickname=&email=
		// parameters
		account := c.GetString("account")
		password1 := c.GetString("password1")
		password2 := c.GetString("password2")
		nickname := c.GetString("nickname")
		email := c.GetString("email")

		// validate
		if password1 != password2 {
			c.JsonResult(1, "登录密码与确认密码不一致")
		}
		if l := strings.Count(password1, ""); password1 == "" || l > 20 || l < 6 {
			c.JsonResult(1, "密码必须在6-20个字符之间")
		}
		if ok, err := regexp.MatchString(common.RegexpEmail, email); !ok || err != nil || email == "" {
			c.JsonResult(1, "邮箱格式错误")
		}
		if l := strings.Count(nickname, "") - 1; l < 2 || l > 20 {
			c.JsonResult(1, "用户昵称限制在2-20个字符")
		}

		// register
		member := models.NewMember()
		member.Account = account
		member.Password = password1
		member.Nickname = nickname
		member.Email = email
		member.Avatar = common.DefaultAvatar()
		member.Role = common.MemberGeneralRole
		member.CreateAt = 0
		member.Status = 1
		if err := member.Create(); err != nil {
			beego.Error(err)
			c.JsonResult(1, err.Error())
		}

		// login
		member.LastLoginTime = time.Now()
		err := member.Update()
		if err != nil {
			fmt.Println(err)
		}

		// cookie && session
		err = c.remember(member.MemberId)
		if err != nil {
			c.JsonResult(1, err.Error())
		}

		// 调用错误，使用session前记得开启
		//// setsessions
		//c.SetSessions(*member)
		//// setcookies
		//err = c.SetCookies(*member)
		//if err != nil {
		//	c.JsonResult(1, err.Error())
		//}

		c.JsonResult(0, "注册成功")
	} else {
		c.TplName = "account/register.html"
	}
}

func (c *AccountController) Login() {
	// 验证session和cookie
	if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
		var cm CookieMemeber
		if err := utils.Decode(cookie, &cm); err == nil {
			//_, err := models.NewMember().Find(cm.MemberId)
			err := c.remember(cm.MemberId)
			if err == nil {
				c.Redirect(beego.URLFor("HomeController.Index"), 302)
			} else {
				c.JsonResult(1, err.Error())
			}
		}
	}

	if c.Ctx.Input.IsPost() {
		account := c.GetString("account")
		password := c.GetString("password")
		member, err := models.NewMember().Login(account, password)
		fmt.Println(err)
		if err != nil {
			c.JsonResult(1, "登陆失败")
		}

		// login
		member.LastLoginTime = time.Now()
		err = member.Update()
		if err != nil {
			fmt.Println(err)
		}

		c.JsonResult(0, "登陆成功")
	} else {
		c.TplName = "account/login.html"
	}
}

func (c *AccountController) Logout() {
	c.SetSessions(models.Member{})
	c.SetSecureCookie(common.AppKey(), "login", "", -1)
	c.Redirect(beego.URLFor("AccountController.Login"), 302)
}

func (c *AccountController) remember(id int) error {
	member, err := models.NewMember().Find(id)
	if member.MemberId == 0 {
		return errors.New("用户不存在")
	}

	// TODO Handler crashed with error runtime error: invalid memory address or nil pointer dereference
	// setsessions
	c.SetSessions(*member)
	// setcookies
	err = c.SetCookies(*member)
	if err != nil {
		return err
	}

	return nil
}
