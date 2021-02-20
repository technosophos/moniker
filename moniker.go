package moniker

import (
	"math/rand"
	"strings"
	"time"
)

// New returns a generic namer using the default word lists.
func New() Namer {
	return &defaultNamer{
		Descriptor: Descriptors,
		Noun:       Animals,
		r:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type defaultNamer struct {
	Descriptor, Noun []string
	r                *rand.Rand
}

func (n *defaultNamer) Name() string {
	return n.NameSep(" ")
}

func (n *defaultNamer) NameSep(sep string) string {
	return n.NameWithOptions(sep, "", false)
}

func (n *defaultNamer) NameWithOptions(sep string, startingLetter string, alliterate bool) string {
	filteredDescriptor :=  n.Descriptor
	if startingLetter != "" { 
		filteredDescriptor = choose(n.Descriptor, startingLetter)
		if len(filteredDescriptor) <= 0 {
			return "No mathing startletter found"
		}
	}

	filteredName := n.Noun
	if alliterate {
		filteredName = choose(n.Noun, startingLetter)
		if len(filteredName) <= 0 {
			return "No mathing startletter found"
		}
	}

	a := filteredDescriptor[n.r.Intn(len(filteredDescriptor))]
	b := filteredName[n.r.Intn(len(filteredName))]

	return strings.Join([]string{a, b}, sep)
}

func hasPrefix(a string, startingLetter string) bool {
    return strings.HasPrefix(a, startingLetter)
}

func choose(ss []string, startingLetter string) (ret []string) {
    for _, s := range ss {
        if strings.HasPrefix(s, startingLetter) {
            ret = append(ret, s)
        }
    }
    return
}

// Namer describes anything capable of generating a name.
type Namer interface {
	// Name returns a generated name.
	Name() string
	// NameSep returns a generated name with words separated by the given string.
	NameSep(sep string) string
	// Name returns a generated name with a chosen separator starting letter and possible alliteration
	NameWithOptions(sep string, startingLetter string, alliterate bool) string
}
