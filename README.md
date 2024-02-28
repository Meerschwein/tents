
**clingo and golang 1.21.6 are required**

[Repository](https://github.com/meerschwein/tents)

To build run: `go build`

To run the tests `go test ./...`

The asp solutions are found in ./pkg/asp/solution

example execution:

```raw
$ cat puzzlefile
5 10
.......... 0
......T... 2
..T...T... 0
.......... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

$ cat puzzlefile | ./tents
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

$ cat puzzlefile | ./tents -p
5 10
. . . . . . . . . . 0
. . A . . . T A . . 2
. . T . . . T . . . 0
. . . . . . A . . . 1
. . . . . . . . . . 0
0 0 1 0 0 0 1 1 0 0

$ cat puzzlefile | ./tents -o asp
lines(5).
columns(10).
free(1,1).
    ...

$ cat aspfile
lines(5).
columns(10).
free(1,1).
    ...

$ cat aspfile | ./tents -f asp
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

$ cat aspfile | ./tents -f asp -n
5 10
.......... 0
......T... 2
..T...T... 0
.......... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

$ cat puzzlefile | ./tents -s negation
5 10
.......... 0
..A...TA.. 2
..T...T... 0
......A... 1
.......... 0
0 0 1 0 0 0 1 1 0 0

$ ./tents -n -o asp puzzlefile | clingo /dev/stdin ./pkg/asp/solution/choice.asp
clingo version 5.6.2
Reading from /dev/stdin ...
Solving...
Answer: 1
 ...
SATISFIABLE

Models       : 1
Calls        : 1
Time         : 0.003s (Solving: 0.00s 1st Model: 0.00s Unsat: 0.00s)
CPU Time     : 0.003s

```
