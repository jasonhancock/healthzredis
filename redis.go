package healthzredis

import (
	"context"

	"github.com/garyburd/redigo/redis"
	"github.com/jasonhancock/healthz"
)

// CheckRedis is the redis healthz check
type CheckRedis struct {
	pool *redis.Pool
}

// NewCheck creates a new CheckRedis.
func NewCheck(pool *redis.Pool) CheckRedis {
	return CheckRedis{
		pool: pool,
	}
}

// Check performs the check by issuing a PING command to the Redis server
func (c CheckRedis) Check(ctx context.Context) *healthz.Response {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")

	return &healthz.Response{
		Error: err,
	}
}
