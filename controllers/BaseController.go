package controllers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"mdbook/common"
	"mdbook/models"
	"mdbook/utils"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	Member          *models.Member    //用户
	Option          map[string]string //全局设置
	EnableAnonymous bool              //开启匿名访问
}

type CookieMemeber struct {
	MemberId int
	Account  string
	Time     time.Time
}

//每个子类Controller公用方法调用前，都执行一下Prepare方法
func (c *BaseController) Prepare() {
	c.Member = models.NewMember() //初始化
	c.EnableAnonymous = false
	//从session中获取用户信息
	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.MemberId > 0 {
		c.Member = &member
		fmt.Println("base/prepare session member ----------", c.Member)
	} else {
		//如果Cookie中存在登录信息，从cookie中获取用户信息
		if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
			var remember CookieMemeber
			err := utils.Decode(cookie, &remember)
			if err == nil {
				member, err := models.NewMember().Find(remember.MemberId)
				if err == nil {
					c.SetSessions(*member)
					c.Member = member
					fmt.Println("base/prepare cookie member ----------", c.Member)
				}
			}
		}
	}
	if c.Member.RoleName == "" {
		c.Member.RoleName = common.Role(c.Member.MemberId)
	}
	c.Data["Member"] = c.Member
	c.Data["BaseUrl"] = c.BaseUrl()
	c.Data["SITE_NAME"] = "MBOOK"
	//设置全局配置
	c.Option = make(map[string]string)
	c.Option["ENABLED_CAPTCHA"] = "false"
}

func (c *BaseController) JsonResult(code int, msg string, data ...interface{}) {
	res := make(map[string]interface{}, 3)
	res["errcode"] = code
	res["message"] = msg

	if len(data) > 0 && data[0] != nil {
		res["data"] = data[0]
	}

	ret, err := json.Marshal(res)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// gzip
	if strings.Contains(strings.ToLower(c.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		c.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		w := gzip.NewWriter(c.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(ret)
		w.Flush()
	} else {
		io.WriteString(c.Ctx.ResponseWriter, string(ret))
	}
	c.StopRun()
}

func (c *BaseController) SetSessions(member models.Member) {
	if member.MemberId <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(common.SessionName, member)
		c.SetSession("uid", member.MemberId)
	}
}

func (c *BaseController) SetCookies(member models.Member) error {
	var cm CookieMemeber
	cm.MemberId = member.MemberId
	cm.Account = member.Account
	cm.Time = time.Now()
	res, err := utils.Encode(cm)
	if err != nil {
		return err
	}
	c.SetSecureCookie(common.AppKey(), "login", res, 24*3600*1)
	return nil
}

func (c *BaseController) BaseUrl() string {
	host := beego.AppConfig.String("sitemap_host")
	if len(host) > 0 {
		if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
			return host
		}
		return c.Ctx.Input.Scheme() + "://" + host
	}
	return c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
}
