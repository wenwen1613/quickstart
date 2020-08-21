package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Upload() {
	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		log.Fatal("上传文件失败")
	}
	defer file.Close()
	c.SaveToFile("file", "static/upload"+fileHeader.Filename)

}
