# D - Chalkboard Median

**Time Limit:** 2 sec  
**Memory Limit:** 1024 MiB  
**Points:** 400

## Problem Statement

An integer X is initially written on a chalkboard.

You are given Q queries to process in order. For the i-th query (1≤i≤Q):

- Two integers A_i and B_i are given.
- Write two new integers A_i and B_i on the chalkboard.
- Then, output the median of all 2i+1 integers currently written on the chalkboard.

## Constraints

- 1 ≤ X ≤ 10^9
- 1 ≤ Q ≤ 2×10^5
- 1 ≤ A_i, B_i ≤ 10^9
- All input values are integers

## Input Format

```
X
Q
A_1 B_1
A_2 B_2
⋮
A_Q B_Q
```

## Output Format

Output Q lines. The i-th line should output the answer for the i-th query.

## Examples

### Example 1

**Input:**
```
5
3
2 3
1 2
8 9
```

**Output:**
```
3
2
3
```

**Explanation:**
- After the 1st query: The integers on the chalkboard are 5, 2, 3. The median is 3.
- After the 2nd query: The integers on the chalkboard are 5, 2, 3, 1, 2. The median is 2.
- After the 3rd query: The integers on the chalkboard are 5, 2, 3, 1, 2, 8, 9. The median is 3.

### Example 2

**Input:**
```
1
4
2 3
4 5
6 7
8 9
```

**Output:**
```
2
3
4
5
```

### Example 3

**Input:**
```
278117031
7
167642909 517897721
148434323 567739597
319926999 481642530
659199879 252516557
49913403 798318034
89701408 892537201
199166668 742285869
```

**Output:**
```
278117031
278117031
319926999
319926999
319926999
319926999
319926999
```
