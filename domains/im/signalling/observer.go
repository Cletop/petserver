package signalling

type Observer interface {
	update(string, string)
	getID() string
}
