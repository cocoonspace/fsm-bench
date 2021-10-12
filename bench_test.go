package fsm_bench

import (
	"testing"

	csfsm "github.com/cocoonspace/fsm"
	llfsm "github.com/looplab/fsm"
)

const (
	StateFoo csfsm.State = iota
	StateBar
)

const (
	EventFoo csfsm.Event = iota
	EventBar
)

func BenchmarkCocoonSpaceFSM(b *testing.B) {
	f := csfsm.New(StateFoo)
	f.Transition(csfsm.On(EventFoo), csfsm.Src(StateFoo), csfsm.Dst(StateBar))
	f.Transition(csfsm.On(EventBar), csfsm.Src(StateBar), csfsm.Dst(StateFoo))
	for i := 0; i < b.N; i++ {
		f.Event(EventFoo)
		f.Event(EventBar)
	}
}

func BenchmarkLooplabFSM(b *testing.B) {
	f := llfsm.NewFSM(
		"foo",
		llfsm.Events{
			{Name: "foo", Src: []string{"foo"}, Dst: "bar"},
			{Name: "bar", Src: []string{"bar"}, Dst: "foo"},
		},
		llfsm.Callbacks{},
	)
	for i := 0; i < b.N; i++ {
		f.Event("foo")
		f.Event("bar")
	}
}
