package service

type HomepageCardProvider interface {
	Id() string
	Provide(userId uint64) interface{}
	StrongDepend() bool
}
