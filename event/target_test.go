package event

import (
	"fmt"
	"reflect"
	"testing"
)

func hello(evt Event) {
	fmt.Println("hello", evt.Type)
}

// TestTarget tests EventTarget
func TestTarget(t *testing.T) {
	target := NewEventTarget()

	if !reflect.DeepEqual(target, Target{}) {
		t.FailNow()
	}

	target.AddEventListener("click", hello, ListernerOptions{})
	target.DispatchEvent(NewEvent("click", InitOptions{}))
}
