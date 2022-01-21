# conways-game-of-life-go

Conways Game of Life in Golang. 

To play the game run:

```
make build run
```

Rules:
- Any live cell with fewer than two live neighbours dies (referred to as underpopulation or exposure).
- Any live cell with more than three live neighbours dies (referred to as overpopulation or overcrowding).
- Any live cell with two or three live neighbours lives, unchanged, to the next generation.
- Any dead cell with exactly three live neighbours will come to life.