// Copyright (c) 2017-2024 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package shitgibbon

import (
	"fmt"
	"math/rand"
	"time"
)

var expletives = []string{
	"cock",
	"cunt",
	"dick",
	"douche",
	"fart",
	"fuck",
	"jizz",
	"piss",
	"shit",
	"slut",
	"splooge",
	"splunk",
	"vag",
}

var trochees = []string{
	"biscuit",
	"puffin",
	"gibbon",
	"trumpet",
	"waffle",
	"weasel",
	"womble",
}

type generator struct {
	rng     *rand.Rand
	useDash bool
}

// Generator defines a "shitgibbon" generator interface.
type Generator interface {
	// Returns a Boolean flag indicating whether or not the generator is using a dash between expletive and trochee.
	UseDash() bool
	// Sets the Boolean flag governing the use of a dash.
	SetUseDash(flag bool)
	// Make makes a "shitgibbon".
	Make() string
	// MakeAll makes all possible "shitgibbon" combinations.
	MakeAll() []string
}

// NewGenerator creates and returns a new Generator.
func NewGenerator() Generator {
	return &generator{
		rng:     rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		useDash: false,
	}
}

func (sg *generator) UseDash() bool {
	return sg.useDash
}

func (sg *generator) SetUseDash(flag bool) {
	sg.useDash = flag
}

func (sg *generator) Make() string {
	expletive := expletives[sg.rng.Intn(len(expletives))]
	trochee := trochees[sg.rng.Intn(len(trochees))]
	if sg.useDash {
		return fmt.Sprintf("%s-%s", expletive, trochee)
	}
	return expletive + trochee
}

func (sg *generator) MakeAll() []string {
	shitgibbons := make([]string, 0, len(expletives)*len(trochees))
	for _, trochee := range trochees {
		for _, expletive := range expletives {
			if sg.useDash {
				shitgibbons = append(shitgibbons, fmt.Sprintf("%s-%s", expletive, trochee))
			} else {
				shitgibbons = append(shitgibbons, expletive+trochee)
			}
		}
	}
	return shitgibbons
}
