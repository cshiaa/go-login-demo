package controller

import (
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/source/storage"
)

var Storage = storage.FileStorage{}

func UploadFile(c *gin.Context) {

	// file, err := c.FormFile("file")
	// if err != nil {
	// 	c.String(500, "上传图片出错")
	// }
    file, err := c.FormFile("file")
    if err != nil {
		global.RY_LOG.Error("上传文件失败")
        return
    }
	// c.JSON(200, gin.H{"message": file.Header.Context})
	// c.SaveUploadedFile(file, file.Filename)
	// c.String(http.StatusOK, file.Filename)
	err = Storage.UploadFile(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success"})
}

func GetFile(c *gin.Context) {
	// 读取yaml文件内容
	yamlContent, err := ioutil.ReadFile("./upload/config.yaml")
	if err != nil {
		c.String(http.StatusInternalServerError, "无法读取文件")
		return
	}

	// 设置响应头，以便前端可以正确解析YAML文件
	c.Header("Content-Type", "multipart/form-data")
	c.Header("Content-Disposition", "form-data; name=file; filename=config.yaml")

	// 将文件内容作为响应体返回给前端
	c.Writer.Write(yamlContent)
}