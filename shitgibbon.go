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

// Make makes a random "shitgibbon".
func Make() string {
	rand.Seed(time.Now().UTC().UnixNano())
	expletive := expletives[rand.Intn(len(expletives))]
	trochee := trochees[rand.Intn(len(trochees))]
	return fmt.Sprintf("%s%s", expletive, trochee)
}
