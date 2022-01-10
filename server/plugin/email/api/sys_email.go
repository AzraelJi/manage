package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"zga-client-manage/global"
	"zga-client-manage/model/common/response"
	email_response "zga-client-manage/plugin/email/model/response"
	"zga-client-manage/plugin/email/service"
)

type EmailApi struct{}

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	if err := service.ServiceGroupApp.EmailTest(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

// @Summary 发送邮件
// @Security ApiKeyAuth
// @Author sun.ji
// @Description 发送邮件
// @Tags System
// @Param x-token	header	string false "token 31a165baebe6dec616b1f8f3207b4273"
// @Param body body	email_response.Email true "JSON数据"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"successful operation"}"
// @Router	/email/sendEmail [POST]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	_ = c.ShouldBindJSON(&email)
	if err := service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
