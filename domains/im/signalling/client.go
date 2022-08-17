package signalling

type Client struct {
	observerList []Observer
	name         string
	// instock      bool
}

func (c *Client) Register(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *Client) Deregister(observer Observer) {
	c.observerList = removeFormObserverList(c.observerList, observer)
}

func removeFormObserverList(observerList []Observer, item Observer) []Observer {
	observerListLength := len(observerList)
	for i, value := range observerList {
		if value.getID() == item.getID() {
			observerList[observerListLength-1], observerList[i] =
				observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (c *Client) Notify(id string, message string) {
	for _, observer := range c.observerList {
		if observer.getID() == id {
			observer.update(c.name, message)
		}
	}
}

func (c *Client) NotifyAll(message string) {
	for _, observer := range c.observerList {
		observer.update(c.name, message)
	}
}

func (c *Client) GetName() string {
	return c.name
}
