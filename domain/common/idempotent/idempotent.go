package idempotent

type Idempotent interface {
	Check(k string) error
	Add(k string) error
}
