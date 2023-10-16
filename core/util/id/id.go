package id

import (
	"github.com/google/uuid"
)

// idUtil 提供了UUID、Snowflake、数据中心ID、机器ID、NanoId等生成和获取相关功能
type idUtil struct {
}

// NewIdUtil 创建一个idUtil实例
func IdUtil() *idUtil {
	return &idUtil{}
}

// RandomUUID 生成随机UUID
func (iu *idUtil) RandomUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

// SimpleUUID 生成简化的UUID（去掉横线）
func (iu *idUtil) SimpleUUID() string {
	u := iu.RandomUUID()
	return u[:8] + u[9:13] + u[14:18] + u[19:23] + u[24:]
}

// FastUUID 生成随机UUID，使用性能更好的ThreadLocalRandom生成UUID
func (iu *idUtil) FastUUID() string {
	u := uuid.New()
	return u.String()
}

// FastSimpleUUID 生成简化的UUID（去掉横线），使用性能更好的ThreadLocalRandom生成UUID
func (iu *idUtil) FastSimpleUUID() string {
	u := iu.FastUUID()
	return u[:8] + u[9:13] + u[14:18] + u[19:23] + u[24:]
}
