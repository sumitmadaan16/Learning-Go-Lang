package observer

import "fmt"

type Observer interface {
	Update(string)
}

type Subject interface {
	Register(observer Observer)
	NotifyAll(msg string)
}

// Now we implement Subject in a struct (NewsAgency)
type newsAgency struct {
	observers []Observer
}

// Register → adds observers to the list
func (a *newsAgency) Register(observer Observer) {
	a.observers = append(a.observers, observer)
}

// NotifyAll → loops through all observers and calls their Update.
func (a *newsAgency) NotifyAll(msg string) {
	for _, o := range a.observers {
		o.Update(msg)
	}
}

// creating and observer struct
type EmailClient struct{}

// Observers implement the Update method:
func (e EmailClient) Update(msg string) { fmt.Println("Email received:", msg) }

type SMSClient struct{}

func (s SMSClient) Update(msg string) { fmt.Println("SMS received:", msg) }

func ObserverPattern() {
	agency := &newsAgency{}
	agency.Register(EmailClient{})
	agency.Register(SMSClient{})
	agency.NotifyAll("Breaking news: Observer pattern with interface!")
}
