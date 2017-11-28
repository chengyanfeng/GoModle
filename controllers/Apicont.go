package controllers

import (
	"github.com/astaxie/beego"
	."testthree/models"
	"fmt"
)
type TestController struct {
	beego.Controller
}

func (this *TestController) Getstudent(){
	stduenttree:=&Studentth{}
	credit:=Credit{}
	DB.Find(stduenttree)
	fmt.Print(stduenttree)
	DB.Model(&stduenttree).Related(&credit, "Id")
	this.Ctx.WriteString("aaa")


}