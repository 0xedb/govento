package target

// EventListener is an event callback function
type EventListener func()

// EventListernerOptions pecifies characteristics 
// about the event listener
type EventListernerOptions struct {
	// Capture indicates that event will 
	// be dispatched to the registered listener 
	// before being dispatched to any 
	// EventTarget beneath it in hierarchy
	Capture bool

	// Once indicates that the listener should be
	// invoked at most once after being added.
	// If true, the listener would be 
	// automatically removed when invoked
	Once    bool


	//
	Passive bool
}

// EventTarget represents an object that can receive
// events and may have listeners
type EventTarget struct{}

// Interface represents the target object interface
type Interface interface {
	AddEventListener()
	RemoveEventListener()
	DispatchEvent()
}

// Init creates a new EventTarget
func (target EventTarget) Init() EventTarget {
	return EventTarget{}
}

//AddEventListener registers an event handler of a
// specific event type on the EventTarget
func (target EventTarget) AddEventListener(eventType string, listener EventListener) {

}
