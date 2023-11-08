package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
	"httpserver/util"
	"log"
	"path/filepath"
)

type FileController struct {
	web.Controller
}

func (file *FileController) AlbumPage() {
	file.TplName = "album.html"
}

func (file *FileController) FileUpload() {
	f, h, err := file.GetFile("uploadfile")
	if err != nil {
		log.Fatal("getfile err ", err)
		return
	}
	defer f.Close()

	file_name := h.Filename
	//file_size := h.Size
	file_type := "other"

	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(file_name)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" || fileExt == "webp" {
		file_type = "img"
	}
	file_path := "./static/upload/" + file_name
	file_hash := util.FileHash(file_path)

	succ := models.QueryInsertFile(file_name, file_path, file_hash, file_type)
	if !succ {
		file.Data["json"] = map[string]interface{}{"code": 6, "message": "文件上传出错 || 已经存在"}
	} else {
		file.SaveToFile("uploadfile", "./static/upload/"+file_name)
		file.Data["json"] = map[string]interface{}{"code": 6, "message": "文件上传成功"}
	}

	file.ServeJSON()
}
