package solution

import _ "embed"

var Solutions = map[string]string{
	// newlines so that line comments dont accidentally comment out the first line
	"choice":      "\n" + choice + "\n",
	"disjunction": "\n" + disjunction + "\n",
	"negation":    "\n" + negation + "\n",
}

//go:embed choice.asp
var choice string

//go:embed disjunction.asp
var disjunction string

//go:embed negation.asp
var negation string
