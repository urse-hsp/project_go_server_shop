package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadService interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (string, error)
}

func NewUploadService(
	service *Service,
) UploadService {
	return &uploadService{
		Service: service,
	}
}

type uploadService struct {
	*Service
}

// ================= 创建 =================

func (s *uploadService) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// 校验
	if ext != ".jpg" && ext != ".png" {
		return "", fmt.Errorf("格式不支持")
	}

	// 目录
	date := time.Now().Format("20060102")
	dir := filepath.Join("storage", "uploads", date)
	os.MkdirAll(dir, os.ModePerm)

	// 文件名
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	savePath := filepath.Join(dir, filename)

	// ⚠️ 注意：这里不能用 c.SaveUploadedFile（controller 才有）
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/storage/uploads/%s/%s", date, filename), nil
}
