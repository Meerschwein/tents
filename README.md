
**clingo and golang 1.21.6 are required**

[Repository](https://github.com/meerschwein/tents)

To build run: `go build`

To run the tests `go test ./...`

The asp solutions are found in ./pkg/asp/solution

example execution:

```raw
# puzzle from stdin
$ cat puzzlefile | ./tents
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

# print spaces for easier verification
# specify puzzlefile using a positional argument
$ ./tents -p puzzlefile
5 10
. . . . . . . . . . 0
. . A . . . T A . . 2
. . T . . . T . . . 0
. . . . . . A . . . 1
. . . . . . . . . . 0
0 0 1 0 0 0 1 1 0 0

# solve the puzzle and output aspfacts
$ ./tents -o asp puzzlefile
lines(5).
columns(10).
free(1,1).
    ...

# take aspfacts as input and solve the puzzle
$ ./tents -f asp aspfile
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

# take aspfacts but don't solve the puzzle
# so convert it to puzzle notation
$ ./tents -f asp -n aspfile
5 10
.......... 0
......T... 2
..T...T... 0
.......... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

# use the negation solver
$ ./tents -s negation puzzlefile
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

# convert puzzle to facts and use clingo to solve the puzzle directly
# and then use tents to display the result
$ ./tents -n -o asp puzzlefile | clingo --outf 1 /dev/stdin ./pkg/asp/solution/choice.asp | sed -n '/ANSWER/{n;p;}' | sed 's/ /\n/g' | ./tents -f asp -n
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0
```
