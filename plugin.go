package main

import (
	"errors"
	"github.com/zbysir/drone-qiniu/lib/qiniu"
)

type Plugin struct {
	AccessKey string
	SercetKey string
	Zone      string
	Bucket    string
	Prefix    string
	Dir       string
	Parallel  int
}

func (p Plugin) Exec() error {
	q := qiniu.NewQiniu(p.AccessKey, p.SercetKey)
	u := q.Uploader()
	zone, exist := qiniu.GetZoneByName(p.Zone)
	if !exist {
		return errors.New("bad zone")
	}

	return u.UploadDir(zone, p.Bucket, p.Prefix, p.Dir, p.Parallel)
}
