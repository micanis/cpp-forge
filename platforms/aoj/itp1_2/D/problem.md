# Circle in a Rectangle

## Problem Statement

長方形の中に円が含まれるかを判定するプログラムを作成してください。

Determine whether a circle is contained within a rectangle.

A rectangle has its lower-left vertex at the origin and its upper-right vertex at coordinates (W, H). A circle is given by its center coordinates (x, y) and radius r.

You need to determine if the circle is completely contained within the rectangle.

## Input

Five integers W, H, x, y, r are given on a single line separated by spaces.

- W, H: dimensions of the rectangle
- x, y: center coordinates of the circle
- r: radius of the circle

## Output

Output "Yes" if the circle is completely contained within the rectangle, and "No" if any part of the circle extends outside the rectangle.

## Constraints

- −100 ≤ x, y ≤ 100
- 0 < W, H, r ≤ 100

## Sample Input/Output

### Sample 1
**Input:**
```
5 4 2 2 1
```

**Output:**
```
Yes
```

### Sample 2
**Input:**
```
5 4 2 4 1
```

**Output:**
```
No
```
