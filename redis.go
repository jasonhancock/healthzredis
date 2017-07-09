package healthzredis

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/jasonhancock/healthz"
)

type CheckRedis struct {
	pool *redis.Pool
}

func NewCheck(pool *redis.Pool) CheckRedis {
	return CheckRedis{
		pool: pool,
	}
}

func (c CheckRedis) Check(ctx context.Context) *healthz.Response {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")

	return &healthz.Response{
		Error: err,
	}
}
