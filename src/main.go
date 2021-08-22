package common

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"googo.co/goo"
	"googo.co/utils"
)

const (
	UPLOAD_DIR = "static/"
)

type Upload struct {
}

func (this Upload) DoHandle(c *gin.Context) *goo.Result {
	fh, err := c.FormFile("file")
	if err != nil {
		return goo.Err(700, "上传失败："+err.Error(), err.Error())
	}
	f, err := fh.Open()
	if err != nil {
		return goo.Err(701, "上传失败："+err.Error(), err.Error())
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return goo.Err(702, "上传失败："+err.Error(), err.Error())
	}
	md5File := utils.MD5(bytes)
	fpath := md5File[0:2] + "/" + md5File[2:4] + "/"
	if err := os.MkdirAll(UPLOAD_DIR+fpath, 0755); err != nil {
		return goo.Err(703, "上传失败："+err.Error(), err.Error())
	}
	fname := fpath + md5File[8:24] + path.Ext(fh.Filename)
	fw, err := os.Create(UPLOAD_DIR + fname)
	if err != nil {
		return goo.Err(704, "上传失败："+err.Error(), err.Error())
	}
	defer fw.Close()
	if _, err := fw.Write(bytes); err != nil {
		return goo.Err(705, "上传失败："+err.Error(), err.Error())
	}
	return goo.Succ(gin.H{
		"url": "/file/" + fname,
	})
}
