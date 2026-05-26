# C - C Stands for Center

**Time Limit:** 2 sec  
**Memory Limit:** 1024 MiB  
**Points:** 300

## Problem Statement

You are given a string S consisting of uppercase English letters.

Find the number of substrings (consecutive subsequences) of S that satisfy all of the following conditions:

1. The substring consists of an odd number of characters.
2. The central character is 'C'. More precisely, if the extracted substring consists of l characters, the ((l+1)/2)-th character is 'C'.

Note: Even if two substrings are identical as strings, if they are extracted from different positions, they are counted separately.

## Constraints

- S is a string of uppercase English letters with length between 1 and 5×10^5

## Input Format

```
S
```

## Output Format

Output the answer.

## Sample Input 1

```
ABCCA
```

## Sample Output 1

```
5
```

**Explanation:**

For S = ABCCA, the substrings satisfying the conditions are:
- Characters 1-5: ABCCA
- Characters 2-4: BCC
- Character 3: C
- Characters 3-5: CCA
- Character 4: C

Total: 5

## Sample Input 2

```
XYZ
```

## Sample Output 2

```
0
```

## Sample Input 3

```
SMBCPROGRAMMINGCONTEST
```

## Sample Output 3

```
11
```
