package controllers

import (
	"github.com/astaxie/beego"
	."testthree/models"
	"fmt"
	"testthree/util"
)
type TestController struct {
	beego.Controller
}

func (this *TestController) Getstudent(){
	stduenttree:=Studentth{}
	//credit:=Credit{}
	DB.Find(&stduenttree)
	fmt.Print(stduenttree)
	//DB.Model(stduenttree).Related(&credit, "Id")
	//返回json 格式的字符串
	util.D(util.Cheng).Add(stduenttree)
	this.Ctx.WriteString(util.JsonEncode(stduenttree))





}