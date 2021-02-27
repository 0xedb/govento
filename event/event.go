package event

import "time"

const (
	// PhaseNone has no event being processed
	PhaseNone uint8 = iota

	// PhaseCapturing has event being propagated through the target's ancestor objects
	PhaseCapturing

	// PhaseTarget has event arrived at the event's target.
	PhaseTarget

	// PhaseBubbling has event propagating back up through the target's ancestors
	// in reverse order, starting with the parent
	PhaseBubbling
)

// InitOptions is a dictionary of event options
type InitOptions struct {
	// Bubbles indicates whether the event bubbles up through
	// it's parent hierarchy  or not
	Bubbles bool

	// Cancelable indicates whether the event can be canceled,
	// and therefore prevented as if the event never happened
	Cancelable bool

	// Composed indicates whether or not the event can bubble
	// across boundary
	Composed bool
}

// Event represents an event
type Event struct {
	InitOptions

	CancelBubble     bool
	CurrentTarget    Target
	DefaultPrevented bool
	EventPhase       uint8
	ReturnValue      bool
	Target           Target
	TimeStamp        time.Time
	Type             string
	IsTrusted        bool
}

// NewEvent creates a new Event
func NewEvent(eventType string, options InitOptions) Event {
	evt := Event{}
	evt.Type = eventType
	evt.Bubbles = options.Bubbles
	evt.Cancelable = options.Cancelable
	evt.Composed = options.Composed

	return evt
}

// ComposedPath returns the event’s path which is an array
// of the objects on which listeners will be invoked
func (event Event) ComposedPath() []Target {
	return []Target{}
}

// PreventDefault  tells the user agent that if the event does not get explicitly
//  handled, its default action should not be taken as it normally would be
func (event Event) PreventDefault() {}

// StopImmediatePropagation  prevents other listeners of the same
// event from being called
func (event Event) StopImmediatePropagation() {}

// StopPropagation prevents further propagation of the current
// event in the capturing and bubbling phases
func (event Event) StopPropagation() {}
