package test

import (
	"github.com/gomodule/redigo/redis"
	"testing"
)

// 连接redis
func TestConnectRedis(t *testing.T) {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	_, err = client.Do("set", "cong", "v11", "EX", 20)
	if err != nil {
		t.Fatal("set字符串失败，", err)
	}
	return
}
