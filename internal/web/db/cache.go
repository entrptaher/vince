package db

import (
	v1 "github.com/vinceanalytics/vince/gen/go/vince/v1"
)

func (db *Config) append(e *v1.Model) {
	hit(e)
	if cached, ok := db.cache.Get(uint64(e.Id)); ok {
		update(cached, e)
		db.db.Add(e)
		return
	}
	db.db.Add(e)
	db.cache.Add(uint64(e.Id), e)
}
