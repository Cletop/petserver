package im

type Observer interface {
	update(string, string)
	getID() string
}
