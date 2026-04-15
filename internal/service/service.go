package service

import (
	"go-server/internal/bootstrap"
	"go-server/pkg/jwt"
	"go-server/pkg/log"
	"go-server/pkg/sid"
	"strconv"
	"strings"
)

// Service 服务
// 小写命名[私有]只能在 service 包内部用
// 负责业务逻辑处理，(调用 Repository/dao 进行数据访问)，调用其他工具包进行辅助功能（如日志、JWT、Sid 等）
type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     bootstrap.Transaction
}

func NewService(
	tm bootstrap.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}

// 手动判空赋值
func AssignIfNotNil[T any](dst *T, src *T) {
	if src != nil {
		*dst = *src
	}
}

// 通用：string → []int
func ParseToIntSlice(str string) ([]int, error) {
	if str == "" {
		return []int{}, nil
	}

	parts := strings.Split(str, ",")
	result := make([]int, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		result = append(result, num)
	}

	return result, nil
}

// 通用：string → []uint（更适合ID）
func ParseToUintSlice(str string) ([]uint, error) {
	if str == "" {
		return []uint{}, nil
	}

	parts := strings.Split(str, ",")
	result := make([]uint, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		num, err := strconv.ParseUint(p, 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, uint(num))
	}

	return result, nil
}
