package dao

func (d *Dao) GetCache(key string) (string, error) {
	return d.redigo.GET(key)
}

func (d *Dao) SetCache(key, value string, expire int) error {
	return d.redigo.SET(key, value, expire)
}
