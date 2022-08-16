package im

type Client struct {
	observerList []Observer
	name         string
	instock      bool
}

func (c *Client) register(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *Client) deregister(observer Observer) {
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

func (c *Client) notify(id string, message string) {
	for _, observer := range c.observerList {
		if observer.getID() == id {
			observer.update(c.name, message)
		}
	}
}

func (c *Client) notifyAll(message string) {
	for _, observer := range c.observerList {
		observer.update(c.name, message)
	}
}

func (c *Client) getName() string {
	return c.name
}
