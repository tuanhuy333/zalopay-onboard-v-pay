package kafka

import "fmt"

type Message struct {
	Topic      string
	Key, Value []byte
	Headers    map[string]string
}

func (m Message) String() string {
	return fmt.Sprintf("topic=%s, key=%s, value=%s, headers=%v", m.Topic, m.Key, m.Value, m.Headers)
}
