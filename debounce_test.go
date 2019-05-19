package debounce

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestChannelWithRandomNotifications(t *testing.T) {
	notifier := NewRandomNotifier(4 * time.Millisecond)
	c := Channel(notifier.Notification, 10*time.Millisecond)

	last := -1
	for i := 0; i < 5; i++ {
		// Eventually there will be 5 stable signals
		n := (<-c).(int)
		log.Println("Test got stable value", n)
		if last == n {
			t.Fatal("Successive signals with the same value")
		}
		last = n
	}
}

func TestWatcherWithRandomNotifications(t *testing.T) {
	notifier := NewRandomNotifier(4 * time.Millisecond)
	w := NewWatcher(notifier.Notification, 10*time.Millisecond)

	last := -1
	for i := 0; i < 5; i++ {
		// Eventually there will be 5 stable signals
		n := w.Watch().(int)
		log.Println("Test got stable value", n)
		if last == n {
			t.Fatal("Successive signals with the same value")
		}
		last = n
	}
}

type RandomNotifier struct {
	Notification chan interface{}
}

func NewRandomNotifier(sleep time.Duration) *RandomNotifier {
	rand.Seed(time.Now().UnixNano())
	c := make(chan interface{})
	go func(c chan interface{}) {
		for {
			time.Sleep(sleep)
			c <- interface{}(rand.Intn(3))
		}
	}(c)
	return &RandomNotifier{c}
}
