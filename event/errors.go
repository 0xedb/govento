package event

import "errors"

// ErrorUnspecifiedEventType occurs when the event's type was not specified
// by initializing the event before the method was called,
// or if the event's type is null or an empty string.
var ErrorUnspecifiedEventType = errors.New("Unspecified Event Type Error occured")
