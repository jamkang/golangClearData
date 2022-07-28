package redis_te

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Te_Exists(c redis.Conn) {
	ret, err := c.Do("exists", "sdaf")
	fmt.Println("ret:", ret)
	fmt.Println("error_te:", err)
}
