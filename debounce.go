package debounce

import (
	"log"
	"time"
)

type Event struct {
	Time    time.Time
	Payload interface{}
}

type Watcher struct {
	BounceTime   time.Duration
	Notification chan interface{}
	report       chan interface{}
	notified     Event
	reported     interface{}
}

func Channel(n chan interface{}, b time.Duration) chan interface{} {
	w := NewWatcher(n, b)
	return w.report
}

func NewWatcher(n chan interface{}, b time.Duration) *Watcher {
	emptyEvent := Event{
		Time:    time.Now(),
		Payload: nil,
	}
	w := Watcher{
		Notification: n,
		BounceTime:   b,
		report:       make(chan interface{}),
		notified:     emptyEvent,
		reported:     emptyEvent,
	}
	go w.run()
	return &w
}

func (w *Watcher) Watch() interface{} {
	return <-w.report
}

func (w *Watcher) run() {
	for {
		n := w.notified
		select {
		case payload := <-w.Notification:
			log.Println("Got notified with", payload)
			if n.Payload != payload {
				w.notified = Event{
					Time:    time.Now(),
					Payload: payload,
				}
			}
		default:
			r := w.reported
			payload := n.Payload
			if n.Payload != nil && r != payload {
				if time.Since(n.Time) > w.BounceTime {
					log.Println("Sending stable at", time.Now())
					w.reported = payload
					w.report <- payload
				}
			}
		}
	}
}
