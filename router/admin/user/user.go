package user

import (
	"echo/common/logger"
	"echo/models/resp"
	"echo/services"
	"echo/utils/google"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"github.com/skip2/go-qrcode"
	"go.uber.org/zap"
)

type Svc struct{}

func Register(eg *echo.Group) {
	s := &Svc{}
	g := eg.Group("/user")
	{
		g.POST("/reset_google_code", s.ResetGoogleCode)
	}
}

func (s *Svc) ResetGoogleCode(c echo.Context) error {
	var params struct {
		Id int64 `json:"id"`
	}
	if err := c.Bind(&params); err != nil {
		return resp.Fail(c, "Parameter error")
	}
	//uid := utils.UserId(c)
	uid := int64(1)
	nickName, googleCode, err := services.User.ResetGoogleCode(uid, params.Id)
	if err != nil {
		logger.AdminLog.Error("[ADMIN] "+fmt.Sprintf("add or modify user GoogleCode --> Fail [%v]", params), zap.Error(err))
		return resp.Fail(c, err.Error())
	}
	url := google.GetGoogleCodeUrl(nickName, googleCode, "金盾")
	base64_img, _ := generateQRCodeBase64(url)
	m := map[string]string{}
	m["url"] = url
	m["google_code"] = googleCode
	m["base64_img"] = base64_img
	return resp.OK(c, m)
}

// 生成二维码图片转base64
func generateQRCodeBase64(content string) (string, error) {
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(png)
	return base64String, nil
}
