package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
	"time"
)

type IFace interface {
	// 获取对应的限流器的健值对名称
	Key(c *gin.Context) string
	// 获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// 新增多个令牌桶
	AddBucket(rules ...BucketRule) IFace
}


type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 存储令牌桶与键值对的映射关系
type BucketRule struct {
	Key          string        // 自定义健值对
	FillInterval time.Duration // 间隔多久时间存放N个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}

type MethodLimiter struct {
	*Limiter
}

//通过RequestURI 获取当前访问的路径
func (m MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
}


func (m MethodLimiter) AddBucket(rules ...BucketRule) IFace {
	for _, rule := range rules {
		if _, ok := m.limiterBuckets[rule.Key]; !ok {
			m.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(
				rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return m
}


func NewMethodLimiter() IFace {
	return MethodLimiter{Limiter: &Limiter{limiterBuckets: map[string]*ratelimit.Bucket{}}}
}
