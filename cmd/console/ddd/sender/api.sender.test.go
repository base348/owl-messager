package sender

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/etc"
	"github.com/lishimeng/owl-messager/sdk"
)

func testMailSender(ctx server.Context) {
	var resp sdk.Response
	var req sdk.MailRequest
	resp.Code = tool.RespCodeSuccess
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	// 请求代理
	resp, err = sdk.New(
		sdk.WithAuth(etc.Config.Console.AppKey, etc.Config.Console.Secret),
		sdk.WithHost(etc.Config.Console.Host)).
		SendMail(req)

	if err != nil {
		resp.Code = tool.RespCodeError
	}

	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}

	ctx.Json(resp)
}

// StockAlarm TemplateParam 库存存量预警（过高/过低）
type StockAlarm struct {
	Name       string  `json:"name"`
	Current    float32 `json:"current"`
	Safe       int     `json:"safe"`
	Suggestion int     `json:"suggestion"`
}

// DateAlarm TemplateParam 保质期预警
type DateAlarm struct {
	Name  string `json:"name"`
	Days  int    `json:"days"`
	Stock int    `json:"stock"`
}

type TemplateParam struct {
	List []interface{} `json:"list"`
}

func testImSender(ctx server.Context) {
	var resp sdk.Response
	var req sdk.ImRequest
	resp.Code = tool.RespCodeSuccess
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	//param := TemplateParam{}

	//param.List = []interface{`{"list":[{"name":"12313","days":2,"stock":5,"ctype":"辅材","expirationDate":"2025-07-15"}]}`}

	req.TemplateParam = `{"list":[{"name":"  12313","days":2,"stock":5,"ctype":"辅材","expirationDate":"2025-07-15"}]}`

	// 请求代理
	resp, err = sdk.New(
		sdk.WithAuth(etc.Config.Console.AppKey, etc.Config.Console.Secret),
		sdk.WithHost(etc.Config.Console.Host)).
		SendIm(req)

	if err != nil {
		resp.Code = tool.RespCodeError
	}

	if resp.Code == -1 {
		resp.Code = tool.RespCodeError
	}

	ctx.Json(resp)
}
