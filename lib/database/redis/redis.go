package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

func New(c *Config) *Redis {

	pool := &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))
			if err != nil {
				panic(err)
			}
			if c.Password != "" {
				if _, err := conn.Do("AUTH", c.Password); err != nil {
					conn.Close()
					panic(err)
				}
			}

			if _, err := conn.Do("SELECT", c.DB); err != nil {
				conn.Close()
				panic(err)
			}
			return conn, nil
		},
	}

	return &Redis{
		pool: pool,
	}
}
