
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


% Place tents on the board
{ tent(X,Y) : cell(X,Y) }.

% Allow tents only on empty spaces
:- tent(X,Y), not free(X,Y).


% Assign each tent exactly one hv-adjacent tree
1 { treeTentPair(TR_X,TR_Y,TE_X,TE_Y) : tree(TR_X,TR_Y), hvadj(TR_X,TR_Y,TE_X,TE_Y) } 1 :- tent(TE_X,TE_Y).

% Each tree must have a tent paired with it
:- tree(X,Y), not treeTentPair(X,Y,_,_).


% Limit the number of tents in each row and column
N { tent(X,Y) : column(Y) } N :- rowsum(X,N).
N { tent(X,Y) : line(X)   } N :- colsum(Y,N).


% Two tents cannot be adjacent
:- tent(X1,Y1), tent(X2,Y2), adj(X1,Y1,X2,Y2).