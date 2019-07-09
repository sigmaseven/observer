package observer

// Event - Generic type for passing observable events.
type Event interface{}

// Subject - This is the type which Observers observe.
type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers(Event)
}

// Observer - This is the type which receives events from Subjects.
type Observer interface {
	Notify(Event)
}

// SubjectComponent - This struct is meant to be embedded into subject
// structs.
type SubjectComponent struct {
	observers []Observer
}

// AddObserver - Dynamically adds an observer to a subject.
func (s SubjectComponent) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver - Dynamically removes an observer from a subject.
func (s SubjectComponent) RemoveObserver(observer Observer) {
	observers := []Observer{}

	for _, o := range s.observers {
		if o != observer {
			observers = append(observers, o)
		}
	}

	s.observers = observers
}

// NotifyObservers - Broadcasts an Event to all observers.
func (s SubjectComponent) NotifyObservers(event Event) {
	for _, observer := range s.observers {
		observer.Notify(event)
	}
}
