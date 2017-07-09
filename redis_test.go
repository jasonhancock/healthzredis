package healthzredis

import (
	"context"
	"net/http"
	"testing"

	"github.com/cheekybits/is"
	"github.com/garyburd/redigo/redis"
	"github.com/jasonhancock/healthz"
)

func TestCheckRedis(t *testing.T) {
	is := is.New(t)

	pool := &redis.Pool{
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

func ExampleCheckRedis() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
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

	checker := healthz.NewChecker()
	checker.AddCheck("redis", NewCheck(pool))

	http.ListenAndServe(":8080", checker)
}
