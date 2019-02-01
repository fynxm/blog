package controllers

type AboutMeController struct {
	BaseController
}

func (this *AboutMeController) Get() {
	this.Preare()
	this.Data["wechat"] = "微信：sjzfynxm"
	this.Data["qq"] = "QQ: 64038944"
	this.Data["tel"] = "TEL: 15321410480"
	this.TplName = "aboultme.html"
}
