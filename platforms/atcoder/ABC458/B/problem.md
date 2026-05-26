# B - Count Adjacent Cells

## Problem Statement

There is an H x W grid. We denote the cell at row i (from the top) and column j (from the left) as cell (i, j).

Two cells (x1, y1) and (x2, y2) are **edge-adjacent** if |x1 - x2| + |y1 - y2| = 1.

For every cell, find the number of cells that are edge-adjacent to it.

## Constraints

- 1 ≤ H, W ≤ 50
- All input values are integers

## Input Format

```
H W
```

## Output Format

```
x1,1 x1,2 ⋯ x1,W
x2,1 x2,2 ⋯ x2,W
⋮
xH,1 xH,2 ⋯ xH,W
```

Where xi,j represents the number of cells that are edge-adjacent to cell (i, j).

## Sample Input/Output

### Sample 1

**Input:**
```
4 5
```

**Output:**
```
2 3 3 3 2
3 4 4 4 3
3 4 4 4 3
2 3 3 3 2
```

Explanation:
- Cell (1, 5) is edge-adjacent to cells (1, 4) and (2, 5), so 2 cells.
- Cell (2, 3) is edge-adjacent to cells (1, 3), (2, 2), (2, 4), and (3, 3), so 4 cells.
- Cell (4, 2) is edge-adjacent to cells (3, 2), (4, 1), and (4, 3), so 3 cells.

### Sample 2

**Input:**
```
1 1
```

**Output:**
```
0
```

Explanation:
- Cell (1, 1) has no edge-adjacent cells.

### Sample 3

**Input:**
```
12 8
```

**Output:**
```
2 3 3 3 3 3 3 2
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
3 4 4 4 4 4 4 3
2 3 3 3 3 3 3 2
```
