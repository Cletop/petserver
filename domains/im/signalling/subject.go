package signalling

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	Notify()
	NotifyAll()
	GetName() string
}
