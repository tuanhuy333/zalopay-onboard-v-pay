package kafka

import "testing"

func TestMessageString(t *testing.T) {
	t.Run("return message string", func(t *testing.T) {
		m := Message{
			Topic:   "test",
			Key:     []byte("test"),
			Value:   []byte("test"),
			Headers: map[string]string{"test": "test"},
		}
		if m.String() != "topic=test, key=test, value=test, headers=map[test:test]" {
			t.Fatalf("message string != %v", m.String())
		}
	})
}
