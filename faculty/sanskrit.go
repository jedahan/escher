// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly, unless you have a better idea.

package faculty

import (
	"github.com/gocircuit/escher/tree"
)

// Time—>SentenceFunctional
// Time zero is most recent, time one is earlier, and so on.
type Sentence tree.Tree

func (s Sentence) Grow(time int, valve tree.Name, value tree.Meaning) Sentence {
	return Sentence(
		tree.Tree(s).Grow(
			time,
			SentenceFunctional{
				"Valve": tree.Branch{valve}, 
				"Value": tree.Branch{value}, 
				"Time": tree.Branch{time},
			},
		).Collapse(),
	)
}

func (s Sentence) At(time int) SentenceFunctional {
	return tree.Tree(s).At(time).(SentenceFunctional)
}

// "Valve"—>tree.Name, "Value"—>tree.Meaning, "Time"—>int
type SentenceFunctional tree.Tree

func (sf SentenceFunctional) Valve() tree.Name {
	return tree.Tree(sf).AtName("Valve")
}

func (sf SentenceFunctional) Value() tree.Meaning {
	return tree.Tree(sf).At("Value")
}

func (sf SentenceFunctional) Time() int {
	return tree.Tree(sf).AtInt("Time")
}

// Valve—>MemoryFunctional
type Memory tree.Tree

func (m Memory) Grow(valve tree.Name, age int, value tree.Meaning) Memory {
	return Memory(
		tree.Tree(m).Grow(
			valve, 
			MemoryFunctional{
				"Valve": tree.Branch{valve}, 
				"Value": tree.Branch{value}, 
				"Age": tree.Branch{age},
			},
		).Collapse(),
	)
}

func (m Memory) At(valve tree.Name) MemoryFunctional {
	return tree.Tree(m).At(valve).(MemoryFunctional)
}

// "Valve"—>tree.Name, "Value"—>tree.Meaning, "Age"—>int
type MemoryFunctional tree.Tree

func (mf MemoryFunctional) Valve() tree.Name {
	return tree.Tree(mf).AtName("Valve")
}

func (mf MemoryFunctional) Value() tree.Meaning {
	return tree.Tree(mf).At("Value")
}

func (mf MemoryFunctional) Age() int {
	return tree.Tree(mf).AtInt("Age")
}

// ShortCognize is the cognition interface provided by the Mind's Eye (short-term memory) mechanism.
// The short-term memory is what allows people to process a linguistic sentence with all its structure.
type ShortCognize func(Sentence)
