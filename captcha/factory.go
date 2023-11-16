package captcha

import (
	"log"
	"sync"
)

// CaptchaServiceFactory 验证码服务工厂
type CaptchaServiceFactory struct {
	serviceMap  map[int]Captcha
	serviceLock sync.RWMutex
}

func NewCaptchaServiceFactory() *CaptchaServiceFactory {
	factory := &CaptchaServiceFactory{
		serviceMap: make(map[int]Captcha),
	}
	return factory
}

//func (c *CaptchaServiceFactory) SetCache(cache store.Cache) {
//	c.serviceLock.Lock()
//	defer c.serviceLock.Unlock()
//	c.cache = cache
//}
//
//func (c *CaptchaServiceFactory) GetCache() store.Cache {
//	c.serviceLock.Lock()
//	defer c.serviceLock.Unlock()
//	return c.cache
//}

func (c *CaptchaServiceFactory) RegisterService(key int, service Captcha) {
	c.serviceLock.Lock()
	defer c.serviceLock.Unlock()
	c.serviceMap[key] = service
}

func (c *CaptchaServiceFactory) GetService(key int) Captcha {
	c.serviceLock.RLock()
	defer c.serviceLock.RUnlock()
	if _, ok := c.serviceMap[key]; !ok {
		log.Printf("未注册%d类型的验证码", key)
	}
	return c.serviceMap[key]
}
