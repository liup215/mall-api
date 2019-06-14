package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *Redis) HGETALL(key string, vl interface{}) error {
	c := r.pool.Get()
	defer c.Close()

	v, err := redis.Values(c.Do("HGETALL", key))
	if err != nil {
		return err
	}

	if err := redis.ScanStruct(v, &vl); err != nil {
		return err
	}

	return nil

}

func (r *Redis) HMSET(key string, v interface{}) error {
	c := r.pool.Get()
	defer c.Close()
	_, err := c.Do("HMSET", redis.Args{}.Add("key").AddFlat(&v)...)
	return err

}
