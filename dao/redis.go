package dao

import (
	"finder/model"
	"github.com/gomodule/redigo/redis"
)

func (d *Dao) GetRecordCache(id int64) (record *model.Record, err error) {
	c := d.Cpool.Get()
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

func (d *Dao) GetCommentCache(id int64) (comment *model.Comment, err error) {
	c := d.Cpool.Get()
	defer c.Close()

	redisKey := RKComment(id)
	reply, err := redis.Bytes(c.Do("GET", redisKey))
	if err == nil {
		if err = json.Unmarshal(reply, &comment); err == nil {
			return
		}
	}
	return
}
