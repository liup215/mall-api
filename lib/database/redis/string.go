package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *Redis) GET(key string) (string, error) {
	c := r.pool.Get()
	defer c.Close()

	return redis.String(c.Do("GET", key))

}

func (r *Redis) SET(key, value string, expire int) error {
	c := r.pool.Get()
	defer c.Close()

	_, err := c.Do("SET", key, value)
	if err != nil {
		return err
	}

	if expire > 0 {

		_, err = c.Do("EXPIRE", key, expire)
	}
	return err
}
