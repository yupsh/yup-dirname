// Command yup-dirname is the CLI wrapper around github.com/gloo-foo/cmd-dirname.
package main

import (
	"strings"

	clix "github.com/gloo-foo/cli"
	command "github.com/gloo-foo/cmd-dirname"
)

// version is the build version. It defaults to "dev" for local builds and is
// overridden at release time via the linker: -ldflags "-X main.version=<v>".
var version = "dev"

const name = "dirname"

// synopsis is the multi-line --help usage block. urfave/cli indents the whole
// block three spaces, so the lines stay flush-left.
const synopsis = `dirname [NAME...]

Output each NAME with its last non-slash component and trailing slashes
removed; if NAME contains no /'s, output '.' (meaning the current
directory). With no NAME, read paths from standard input.`

// spec declares the dirname wrapper: a filter over path lines. Each NAME operand
// is a literal path fed as an input line, not a file to read; with no operand
// the paths are read from standard input.
var spec = clix.Spec{
	Name:     name,
	Summary:  "strip last component from file name",
	Synopsis: synopsis,
	Build:    build,
}

// build maps the invocation to dirname's pipeline: the NAME operands become the
// input lines (or standard input when none are given), fed through the dirname
// command.
func build(inv clix.Invocation) (clix.Source, clix.Command, error) {
	if inv.Args.NArg() == 0 {
		return clix.Stdin(inv.Stdin), command.Dirname(), nil
	}
	lines := strings.Join(inv.Args.Args().Slice(), "\n") + "\n"
	return clix.Stdin(strings.NewReader(lines)), command.Dirname(), nil
}

// runMain is an indirection seam so main's wiring is testable without spawning
// the process; a test swaps it and restores it.
var runMain = clix.Main

func main() { runMain(spec, version) }
