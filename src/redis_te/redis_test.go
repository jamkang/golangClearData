package redis_te

import (
	"testing"
)

func TestRedis(t *testing.T) {
	c := ListRedis()
	//RedisHash(c)
	//RedisList(c)
	//RedHash(c)
	Te_Exists(c)
	defer c.Close()
}
