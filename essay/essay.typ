#set page(header: grid(
  columns: (1fr,) * 2,
  [Introduction to Artificial Intelligence I],
  align(right, "Merlin Volkmer"),
))

#v(2cm)
#align(center, text(size: 2em, [Tents and Trees]))
#v(2cm)

#set text(lang: "en", hyphenate: false)
#set par(justify: true)

#let aspchef = link(
  "https://asp-chef.alviano.net/#eJzVV9uSqjgU/SVA7TM8zIOiYGyII3LNm0AraFCrlOby9bMCamOf7pqpmlMzNQ9WuGVn77XXWolv9eIcHSdyopEXsnd3ljOJo2xYmB+/naVXj3tPE+O4DgyMfIXrkSuek+kX87yqfy/mlZhXkul4Z63H3aiNKzKdibEkWZkxf7SPFJsHCi/iuSeR/VmKjx43tcWB+fSdGd42GnjS7V0e+lXD1mUWK+l7Iqsn5ldlb06zMdRBoKglCxZpYvD36Ghly6NdJ757IUd6igbJduOr29jQC+aXX7wbHePcO2x8r0h09cACO91gzd4aNfP1Q6SMLm+eeo7yBPf2GXVIcU1eWK5fYuXbeCcxF7n145Vh4DXMU7PQvyJenC1zlkZzeluLnmPDzZb8+oPk/ED2pyx0OA/9Wc1y98qmZEjXksT249J0do01ja90OsmZTxTajEfLaSgLnKPBhMe5LqH2QsRwDb1GnruVTP3b+/adqOGtXjSJoZZ9vMNgUkbzwye8Pud6w2XdwyFPz6KHG8NLMdb9mLjvej+YtBiGii4xzCVc3Ya53jAnlAL5twyxeYh+B4pXhAp64ttbjEViAFPwOJkv5I4THrAUfVfrtkZ5MllJrpj/ngR2id82ni84M/g+DGwuvglkdQ8OXcQ620C6oPYza7VBKkuxhsx3a9O3U+a4wBX4rqUhNRap6euZZZBB6Ixr5sSyVRPU2uNdJuqgHPnsN9qEijqjudXmEs/tM3ASHPgBHKVoMO7jmMbzSYvhR13g+hO3gIdvv0d5NerjuTH4ATy6cYz0Y+IZP27mq2y5J6WlIebRu7AA2sgWF3CUx3XqIw7WO+yWmbTr4XKy1mVl6iKWx1/XozjK6Tuur9DWyUNO4MOZTE9XoqNvineJML5OZxcrq+4xi8DAvCP6MKBytE4tseZHzPGLuZ7c87C68XAhswp6cHd/rMcVfOf2vHzKzTF4wdaXytTGtalNntdr82zXW0S5zd+06qnGSMwNrC5GYJ0xfxsolEdGtQ0eeVxE3Nv8cnev1+QV1hiKfhaJL2ct3vuZ6E26CSgXOrr3D3Fa7oUBbXkbD+w69HmRgA+9Hj387jNHO5+E3mSVi5rwzQF5wtvAqePqEeOJvznYq7gNVdzS9EkDnoK/8RD8HYUKaUzHLa2pVdOGHSxF30N3X/B3ksZHysmMOl7Wvf/w1EuGOt5jra91vYyNlpNPfvzk27nwAbVovaTHz40v9oBU8JrHWT+mzVvNe2rZ+ShFfq7gLXoFrSiuQsG3v4uRyIkFqAteGA3QM6D7V/6wgVf+7A88pY2+t5okM31rSJXwSg07s9ZySvd2bvohnNWtrH3cLJ1FttT+t/7Qau210yI0edN+9qEFoX1oYUcM8Y3XdN9iXJe7Tp9jFXt+YR665ybOF0SbwPsrOfJH+P7hJ0KD0+758K776f272zyh0RNwe2j0tZl9mVenUezLOTzmp/w7D2j9Srv50cfaF6KXmYdf53cM+5N785pyt1JUxJeLm299i0+gkQK/No/Wa+B733tdqVrrYXHzIpH3DZ/haVXfc+3WxbwHJgLjVSM949RqdYTrRaslq9Wdfsb+2D93dOdAXU1jcFzwDXzCWefTOeYX6+9ZQ4sDdUhDHUs2/ZUCrQiPwl4Kj2p06Go1WBqzEW0O0NBOot9oKMI5ItEmq4/8UOd/dKbp+4dY168XD56LnKBBcWZBjE4L4lmYc6kd0d944F3FdZyrg/s1g55bf75hibxaTrc1trwS8e68WyStnxkpj3m/Z23vowCGZs4v2JNWooec7g+j0GEprSVpCb8yHfTBn11DH3uDhrOOQkpqrIY4/8gs77wyRs2oDx41FDEcYFSHym7n+OJsVv0z7B/",
  "https://asp-chef.alviano.net",
)

= Introduction

"Tents and Trees" is a logic puzzle that requires players to place tents in a
grid next to trees while adhering to specific rules.

- Tents cannot be adjacent to each other, not even diagonally.
- Each tree must have a unique tent horizontally or vertically adjacent to it.
- The numbers along the grid's edges indicate how many tents must be placed in
  each row or column.

This project uses Golang for parsing, validation, and testing, and an ASP
program and Clingo to solve the puzzle.

= Implementation

The puzzle is presented as a two-dimensional grid of cells, where each cell can
be a tree, a tent, or empty (pkg/tents/types.go). It can be serialized to and
from a string and asp facts (pkg/tents/puzzle.go and pkg/tents/asp.go).

The asp solvers employed to solve the puzzle are located in pkg/asp/solution and
are compiled into the binary.

This project uses clingo to solve the puzzle. The puzzle is converted into ASP
facts and solved using Clingo (pkg/clingo/). The resulting solution is then
parsed back into a puzzle.

== Dependencies

The necessary binaries for this project are clingo for puzzle solving and golang
1.21.6 for building.

The project also requires the following dependencies:

- `github.com/alexflint/go-arg`: cli parsing
- `github.com/stretchr/testify`: test assertions
- `golang.org/x/tools`: txtarchive for testdata

These direct dependencies are included for testing and cli parsing, and are
vendored with the project to ensure reproducibility.

== Testing

The tests are located in pkg/tents/testdata and are written in the txtarchive
#footnote[#link("https://pkg.go.dev/golang.org/x/tools/txtar#hdr-Txtar_format")]
format. Each file must contain a puzzle section. If a json section exists, the
puzzle will be serialized and checked against the json. If a solution section
exists, the puzzle will be solved and checked against the solution using all
three solvers. The solvers will then be compared to check if they are
equivalent. If an ASP section exists, the puzzle will be converted to ASP facts
and checked against the ASP section.

== Command line usage

The command line interface is straightforward and should be easy to understand.
It can be used to convert between the puzzle and ASP formats and to solve the
puzzle using different ASP solutions.

```raw
Usage: tents [--informat INFORMAT] [--outformat OUTFORMAT] [--solution SOLUTION] [--nosolve] [--printspaces] [FILE]

Positional arguments:
  FILE                   stdin if not given

Options:
  --informat INFORMAT, -f INFORMAT
                         puzzle | asp [default: puzzle]
  --outformat OUTFORMAT, -o OUTFORMAT
                         puzzle | asp [default: puzzle]
  --solution SOLUTION, -s SOLUTION
                         choices | disjunction | negation [default: choices]
  --nosolve, -n          don't solve the puzzle [default: false]
  --printspaces, -p      print puzzle output with spaces between cells [default: false]
  --help, -h             display this help and exit
```

The input and output format can be specified using `-f` and `-o`. The accepted
formats are the puzzle format from Moodle and the ASP facts from the ASPChef
link on Moodle #footnote(aspchef). To convert between formats without solving
the puzzle, use `-n`. For example, `./tents -o asp -n puzzlefile` will convert
the puzzle to ASP facts without solving it. Use `-s` to specify the ASP program
to solve the puzzle.

= Ethical considerations

A potential military application for a solver that reflects the logic of the
Tents and Trees puzzle could be in the strategic placement of defensive units in
relation to critical assets or locations, similar to the trees in the puzzle. In
this scenario, the critical assets (analogous to trees in the puzzle) could
represent essential infrastructure, command centres, or other crucial resources
that require protection. The defensive units, which are similar to tents, may
consist of various types of forces, including ground troops, anti-aircraft
systems, surveillance units, or unmanned drones, that are responsible for
protecting these assets.

The primary ethical challenge pertains to the autonomous operation of systems
that can use lethal force. It is important to ensure that human oversight is
maintained to prevent any unintended consequences. While such systems can
optimize defense postures with precision and efficiency, it raises concerns
about the autonomy of decision-making processes and the need for human
oversight. Any operational deployment of these systems must strictly adhere to
international humanitarian laws, emphasizing the principles of distinction and
proportionality.

To ensure accountability, it is essential that the deployment of autonomous
defensive units adheres to ethical and legal standards and is subject to human
oversight. Decisions, particularly those with life-or-death consequences, must
be subject to human judgment. Questions regarding responsibility for the actions
of AI systems, particularly in cases resulting in unintended civilian casualties
or conflict escalation, highlight the necessity for well-defined protocols that
maintain human agency in critical decision-making loops.

#pagebreak()

The use of artificial intelligence in defence settings may trigger an arms race
in military AI technologies, which could exacerbate global military tensions and
destabilise already volatile regions. Urgent international cooperation is needed
to establish norms and regulations that govern the development and deployment of
military AI. This is to ensure that advancements in defence technologies do not
compromise global peace and security. The prospect of escalating militarisation
underscores the significance of this.

AI systems used in military contexts must adhere to international humanitarian
law. This requires the systems to differentiate between combatants and
non-combatants and evaluate the necessity and proportionality of force used in
complex conflict situations. The deployment of AI in military applications must
be ethical and committed to protecting civilian lives while upholding the
integrity of combatant engagement. Current AI systems face a challenge in
navigating the nuances of battlefield ethics.