package cacheInfo

// 通过缓存读取用户信息,也就是token中的非敏感信息
type UserCache struct {
	UserName string
}

// 文章的基础信息也应该用cache
