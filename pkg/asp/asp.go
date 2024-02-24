package asp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Meerschwein/tents/pkg/util"
)

type Predicate struct {
	Name string
	Args []int
}

func (p Predicate) String() string {
	args := []string{}
	for _, arg := range p.Args {
		args = append(args, fmt.Sprint(arg))
	}
	return fmt.Sprintf("%s(%s).", p.Name, strings.Join(args, ","))
}

func NewPredicate(name string, args ...int) Predicate {
	return Predicate{name, args}
}

// ParsePredicate parses a string of othe form `name(arg1,arg2,...,argn).`
func ParsePredicate(s string) (Predicate, error) {
	// s = strings.TrimSpace(s)

	lpar := strings.Index(s, "(")
	rpar := strings.Index(s, ")")

	if lpar == -1 || rpar == -1 {
		return Predicate{}, fmt.Errorf("invalid predicate: %s", s)
	}

	name := s[:lpar]

	args := []int{}
	for _, arg := range strings.Split(s[lpar+1:rpar], ",") {
		i, err := strconv.Atoi(arg)
		if err != nil {
			return Predicate{}, err
		}
		args = append(args, i)
	}

	return NewPredicate(name, args...), nil
}

// strings containing only whitespace are ignored
func ParsePredicates(input []string) ([]Predicate, error) {
	preds := []Predicate{}
	for _, line := range input {
		if util.IsBlank(line) {
			continue
		}
		pred, err := ParsePredicate(line)
		if err != nil {
			return nil, err
		}
		preds = append(preds, pred)
	}
	return preds, nil
}
