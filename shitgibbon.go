package shitgibbon

import (
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

var rng = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

// Make makes a random "shitgibbon".
func Make() string {
	expletive := expletives[rng.Intn(len(expletives))]
	trochee := trochees[rng.Intn(len(trochees))]
	return expletive + trochee
}

// MakeAll makes all possible "shitgibbon" combinations.
func MakeAll() []string {
	shitgibbons := make([]string, 0, len(expletives)*len(trochees))
	for _, trochee := range trochees {
		for _, expletive := range expletives {
			shitgibbons = append(shitgibbons, expletive+trochee)
		}
	}
	return shitgibbons
}
