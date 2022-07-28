package redis_te

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func ListRedis() redis.Conn {
	//redisv8.NewClient(&redisv8.Options{
	//	Addr:
	//})
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn redis.md failed,", err)
		return nil
	}
	return c
}

//Hash redis operation
func RedisHash(c redis.Conn) {
	//不区别大小写
	_, err := c.Do("hset", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.Do("HSet", "shui", "abc", 60)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	r2, err := redis.String(c.Do("hget", "shui", "abc"))
	if err != nil {
		fmt.Println("get shui failed,", err)
		return
	}
	fmt.Println(r, r2)
}

//List redis operation
func RedisList(c redis.Conn) {
	//左边插入
	c.Do("lpush", "fruits", "apple", 100)
	//右边插入
	c.Do("rpush", "fruits", "tangerine", 20)

	//左边输出
	reda1, _ := redis.String(c.Do("lpop", "fruits"))
	reda2, _ := redis.String(c.Do("rpop", "fruits"))
	fmt.Println(reda1, "....", reda2)
}

func RedHash(c redis.Conn) {
	_, e := c.Do("hset", "wang", "a1", "12", "a2", 14, "a3", 15)
	if e != nil {
		fmt.Println("set wang failed,", e)
		return
	}
	result, e := redis.StringMap(c.Do("hgetall", "wang"))
	if e != nil {
		fmt.Println("hget wang failed,", e)
		return
	}
	fmt.Println(result)
}
