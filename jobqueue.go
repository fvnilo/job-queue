package jobqueue

type Listener interface {
	Listen() <-chan []byte
}

type Publisher interface {
	Publish(message []byte) error
}
