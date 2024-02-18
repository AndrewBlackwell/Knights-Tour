# The Knight's Tour Problem
A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square exactly once. This type of problem is an instance of a more general Hamiltonian Path problem in graph theory. If a knight's tour ends in a space that is one move away from the starting position, it is considered a closed tour. This program was written with open tours in mind, but may be modified fairly easily to account for closed tours.

### Board:
The chessboard is represented as an 8x8 grid (since the standard chessboard is 8x8 squares). Each cell on the grid can either have the step number of the knight's tour or be marked as unvisited.

### Valid Moves: 
The knight's possible moves are represented as two arrays, dx and dy. These arrays define the relative x and y coordinates the knight can move to from its current position. Since a knight moves in an L-shape, there are eight possible moves from any given position.

### Algorithm:
The algorithm starts from the initial position provided by the user and marks it as the starting point (step 0).
It then recursively tries all eight possible moves of the knight from the current position.
For each move, it checks if the move is valid (i.e., within the bounds of the board and not already visited).
If a move is valid, it marks that position with the current step number and recursively attempts to solve the tour from that new position.
If the tour cannot be completed from a given position (i.e., no further valid moves are available), the algorithm backtracks by resetting the position as unvisited and tries the next possible move.

### Output:
The tour is complete when the knight successfully visits every square on the board exactly once.
The algorithm records each move of the knight as a string in chess notation (like "A1 ---> B3") and adds it to a list of moves.
Once the tour is complete, the program prints out each move in the order it was made.
If no tour is possible from the starting position, the program reports that no solution exists.

### Space for further optimization: 
While this implementation uses a straightforward backtracking method, it does not include advanced optimizations like Warnsdorff's rule, which can significantly speed up the solution finding process by choosing the next move based on heuristic criteria.
