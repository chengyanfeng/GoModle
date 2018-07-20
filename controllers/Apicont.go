package controllers

import (
	"github.com/astaxie/beego"
	."GoModle/models"
	"fmt"
	"GoModle/util"
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
	queryp:=util.P{}
	queryp["cname"]="aaaa"
	//添加
	util.D(util.Cheng).Add(stduenttree)
	//查询所有的cname为aaaa 的
	//db库下的test文档里的数据

	a:=util.D("test").Find(queryp).All()
	fmt.Print(a)
	util.D(util.Cheng).Explain()
	this.Ctx.WriteString(util.JsonEncode(stduenttree))





}