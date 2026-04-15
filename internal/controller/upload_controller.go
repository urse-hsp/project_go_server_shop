package controller

import (
	"fmt"
	v1 "go-server/api/v1"
	"go-server/internal/service"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUploadController(handler *Handler, s service.UploadService) *uploadController {
	return &uploadController{
		Handler: handler,
		Service: s,
	}
}

type uploadController struct {
	*Handler
	Service service.UploadService // 依赖注入
}

// ================= 创建 =================

// @Summary **创建
// @Tags DEMO
// @Accept json
// @Produce json
// @Router /upload [post]

func (u *uploadController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"msg": "获取文件失败"})
		return
	}

	url, err := u.Service.Upload(c, file)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := c.Request.Host
	fullURL := fmt.Sprintf("%s://%s%s", scheme, host, url)

	data := map[string]string{
		"fileName": "",
		"tmp_path": fullURL,
	}
	v1.Success(c, data)
}

func (u *uploadController) Upload2(c *gin.Context) {
	// 1. 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"msg": "获取文件失败"})
		return
	}

	// 2. 创建目录（没有就创建）
	saveDir := "./uploads"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"msg": "创建目录失败"})
		return
	}

	// 3. 生成文件名（防止重复）
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// 4. 拼接路径
	savePath := filepath.Join(saveDir, filename)

	// 5. 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(500, gin.H{"msg": "保存失败"})
		return
	}

	// 6. 返回访问路径
	c.JSON(200, gin.H{
		"url": "/uploads/" + filename,
	})
}
