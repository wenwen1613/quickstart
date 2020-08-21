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
		log.Print("上传文件失败:", err)
		return
	}
	defer file.Close()
	c.SaveToFile("file", "static/upload"+fileHeader.Filename)

}
