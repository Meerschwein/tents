% A cell on the board
cell(X,Y) :- line(X), column(Y).

% Horizontal and vertical adjacency
hvadj(X1,Y1,X2,Y2) :- hvadj(X2,Y2,X1,Y1).

hvadj(X1,Y,X2,Y) :- cell(X1,Y), cell(X2,Y), X2 = X1+1.
hvadj(X,Y1,X,Y2) :- cell(X,Y1), cell(X,Y2), Y2 = Y1+1.

% Horizontal, vertical and diagonal adjacency
adj(X1,Y1,X2,Y2) :- hvadj(X1,Y1,X2,Y2).
adj(X1,Y1,X2,Y2) :- adj(X2,Y2,X1,Y1).

adj(X1,Y1,X2,Y2) :- cell(X1,Y1), cell(X2,Y2), X2 = X1+1, Y2 = Y1+1.
adj(X1,Y1,X2,Y2) :- cell(X1,Y1), cell(X2,Y2), X2 = X1+1, Y2 = Y1-1.


% Place tents on each free cell
tent(X,Y)  :- free(X,Y), not -tent(X,Y).
-tent(X,Y) :- free(X,Y), not tent(X,Y).


% Assign each tree a tent
treeTentPair(RX,RY,EX,EY) :- tree(RX,RY), tent(EX,EY), hvadj(RX,RY,EX,EY),
    not -treeTentPair(RX,RY,EX,EY).
-treeTentPair(RX,RY,EX,EY) :- tree(RX,RY), tent(EX,EY), hvadj(RX,RY,EX,EY),
    not treeTentPair(RX,RY,EX,EY).

:- tree(RX,RY), 1 != #count { EX,EY : treeTentPair(RX,RY,EX,EY) }.
:- tent(EX,EY), 1 != #count { RX,RY : treeTentPair(RX,RY,EX,EY) }.


% Limit the number of tents in each row and column
:- rowsum(X,N), N != #count { Y : tent(X,Y) }.
:- colsum(Y,N), N != #count { X : tent(X,Y) }.


% Two tents cannot be adjacent
:- tent(X1,Y1), tent(X2,Y2), adj(X1,Y1,X2,Y2).