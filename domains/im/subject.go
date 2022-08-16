package im

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notify()
	notifyAll()
}
