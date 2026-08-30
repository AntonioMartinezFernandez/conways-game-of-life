# Conway's Game of Life

Conway's Game of Life is a cellular automaton devised by the British mathematician John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves.

**Rules**:

1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
1. Any live cell with two or three live neighbours lives on to the next generation.
1. Any live cell with more than three live neighbours dies, as if by overpopulation.
1. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

## Getting Started

1. Clone the repository to your local machine.
1. Install Go if you haven't already. You can download it from [the official Go website](https://golang.org/dl/).
1. Navigate to the project directory in your terminal.
1. Run `make run` to start the application.

## Pattern Taxonomy

### Still Lifes

Patterns that do not change between generations:

- **Block**: A 2x2 square of live cells
- **Beehive**: Hexagonal pattern of 6 cells
- **Loaf**: 4-cell pattern resembling a loaf of bread
- **Boat**: 5-cell pattern in a boat-like shape
- **Tub**: 5-cell pattern with one central cell surrounded by 4 cells

### Oscillators

Patterns that repeat after a fixed number of generations:

- **Blinker**: Period 2, alternates between horizontal and vertical line of 3 cells
- **Toad**: Period 2, alternates between two configurations of 6 cells
- **Beacon**: Period 2, two blocks oscillating
- **Pulsar**: Period 3, large symmetric pattern with 48 cells
- **Pentadecathlon**: Period 15, chain of cells

### Spaceships

Patterns that translate across the grid without changing shape:

- **Glider**: Small pattern that moves diagonally (period 4)
- **Lightweight Spaceship (LWSS)**: Moves horizontally (period 4)
- **Middleweight Spaceship (MWSS)**: Moves horizontally (period 4)
- **Heavyweight Spaceship (HWSS)**: Moves horizontally (period 4)

### Methuselahs

Patterns that take a long time to stabilize:

- **R-pentomino**: 5-cell pattern that takes 1103 generations to stabilize
- **Acorn**: 7-cell pattern that takes 5206 generations to stabilize
