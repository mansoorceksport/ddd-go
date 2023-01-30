package redis

import "github.com/mansoorceksport/tavern/domain/common/idempotent"

type Redis struct {
}

func (r Redis) Check(k string) error {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Add(k string) error {
	//TODO implement me
	panic("implement me with ttl")
}

func NewRedis(conn string) idempotent.Idempotent {
	return &Redis{}
}
