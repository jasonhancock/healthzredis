package healthzredis

import (
	"context"
	"testing"

	"github.com/cheekybits/is"
	"github.com/garyburd/redigo/redis"
)

func TestCheckRedis(t *testing.T) {
	is := is.New(t)

	pool := &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	c := NewCheck(pool)
	result := c.Check(context.Background())

	is.NoErr(result.Error)
}

// TODO: example
