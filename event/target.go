package event

import (
	"fmt"
	"reflect"
)

// Listener is an event callback function
type Listener func(Event)

// listenermap is a map with listener as key and
// a map of listerneroptions as value
type listenermap map[*Listener]map[ListernerOptions]ListernerOptions

// eventlisteners is a map of event and their listeners
var eventlisteners map[string]listenermap

// CaptureOptions sets the options for Capture & UseCapture
type CaptureOptions struct {
	// Capture indicates that event will
	// be dispatched to the registered listener
	// before being dispatched to any
	// EventTarget beneath it in hierarchy
	Capture bool

	// UseCapture indicates whether events of this type will be
	// dispatched to the registered listener before being
	// dispatched to any EventTarget beneath it in the hierarchy
	UseCapture bool
}

// ListernerOptions pecifies characteristics
// about the event listener
type ListernerOptions struct {
	CaptureOptions

	// Once indicates that the listener should be
	// invoked at most once after being added.
	// If true, the listener would be
	// automatically removed when invoked
	Once bool

	// Passive indicates that the function specified by listener
	// will never call preventDefault().
	// If a passive listener does call preventDefault(),
	//the user agent will do nothing other than generate a warning
	Passive bool
}

// CompareListenerOptions compares two listener options
func CompareListenerOptions(first, second ListernerOptions) bool {
	return reflect.DeepEqual(first, second)
}

// Target represents an object that can receive
// events and may have listeners
type Target struct{}

// Interface represents the target object interface
type Interface interface {
	// AddEventListener registers an event handler
	AddEventListener(string, Listener, ListernerOptions)
	// RemoveEventListener removes an event handler
	RemoveEventListener(string, Target, CaptureOptions)
	// DispatchEvent dispatches an event to a target
	DispatchEvent(Event, Target) (bool, error)
}

// NewEventTarget creates a new EventTarget
func NewEventTarget() Target {
	return Target{}
}

// AddEventListener registers an event handler of a
// specific event type on the EventTarget
func (target Target) AddEventListener(event string, listener Listener, options ListernerOptions) {


	if event == "" || listener == nil {
		return
	}

	// check if event already exists for the listener
	present, ok := eventlisteners[event]

	callback := present[&listener]

	if ok && CompareListenerOptions(options, callback[options]) {
		return
	}

	 v := callback[options] 
	 fmt.Println(v, "v")
	// present[&listener] 
	// eventlisteners[event] = []interface{}{listener, options}
}

// RemoveEventListener removes an event listener from the EventTarget
func (target Target) RemoveEventListener(event string, listeners Target, options CaptureOptions) {

}

// DispatchEvent dispatches an event to the EventTarget
func (target Target) DispatchEvent(event Event) (bool, error) {
	if event.Type == "" {
		return false, ErrorUnspecifiedEventType
	}

	// The return value is false if event is cancelable and at least one of the event
	// handlers which received event called Event.preventDefault(). Otherwise it returns true.

	// get listeners and call them all with the event
	listeners := eventlisteners[event.Type]
	fmt.Println(listeners)
	for key := range listeners {
		fmt.Println(*key)
		(*key)(event)
	}
	fmt.Println("herezz")

	return true, nil
}
