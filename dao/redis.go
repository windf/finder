package dao

import (
	"finder/model"
	"github.com/gomodule/redigo/redis"
)

func (d *Dao) GetRecordCache(id int64) (record *model.Record, err error) {
	c := d.cpool.Get()
	defer c.Close()

	redisKey := RKRecord(id)
	reply, err := redis.Bytes(c.Do("GET", redisKey))
	if err == nil {
		if err = json.Unmarshal(reply, &record); err == nil {
			return
		}
	}
	return
}

func (d *Dao) GetBatchRecordCache(id int64) (record *model.Record, err error) {
	c := d.cpool.Get()
	defer c.Close()

	redisKey := RKRecord(id)
	reply, err := redis.Bytes(c.Do("GET", redisKey))
	if err == nil {
		if err = json.Unmarshal(reply, &record); err == nil {
			return
		}
	}
	return
}
