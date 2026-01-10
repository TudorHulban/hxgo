package dsl

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGCRun(t *testing.T) {
	el := el("div", Text("hi!"))

	// Wait a bit (simulate real usage)
	runtime.GC()
	runtime.Gosched()

	require.Equal(t,
		"<div>hi!</div>",
		string(Render(el)),

		string(Render(el)),
	)
}

func TestGCActual(t *testing.T) {
	// 1. Create node in inner scope
	var nodePtr *Node

	func() {
		node := el("div",
			el("span", Text("nested")), // Inner node data
		)
		nodePtr = &node // Take address, but inner scope ends
	}()

	// 2. Force GC strongly
	for range 10 {
		runtime.GC()
		runtime.Gosched()

		// Allocate garbage to push out old heap
		_ = make([]byte, 1024*1024)
	}

	// 3. Try to render through dangling pointer
	result := Render(*nodePtr) // UAF risk!

	fmt.Println(
		string(result),
	)
}

// This mimics real usage pattern:
// 1. Build UI in factory function
// 2. Return it
// 3. Render later
func TestStackFrameIssue(t *testing.T) {
	buildUI := func() Node {
		return el("div",
			el("p", Text("Hi")),    // inner el() stack frame
			el("p", Text("World")), // another inner frame
		)
	}

	// Store for later
	ui := buildUI() // All inner stack frames gone!

	// Simulate "later" - force memory reuse
	runtime.GC()

	require.Equal(t,
		"<div><p>Hi</p><p>World</p></div>",
		string(Render(ui)),
	)
}
