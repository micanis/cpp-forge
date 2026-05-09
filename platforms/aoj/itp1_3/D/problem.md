# How Many Divisors?

## Problem Statement

３つの整数 a、b、c を読み込み、a から b までの整数の中に、c の約数がいくつあるかを求めるプログラムを作成してください。

Read three integers a, b, c and count how many divisors of c exist in the range from a to b (inclusive).

## Input

a、b、c が１つの空白区切りで１行に与えられます。

One line containing three integers a, b, c separated by a single space.

## Output

約数の数を１行に出力してください。

Output the count of divisors on a single line.

## Constraints

- 1 ≤ a, b, c ≤ 10000
- a ≤ b

## Sample Input/Output

### Sample 1

**Input:**
```
5 14 80
```

**Output:**
```
3
```

**Explanation:** The divisors of 80 are: 1, 2, 4, 5, 8, 10, 16, 20, 40, 80. Among these, the divisors in the range [5, 14] are: 5, 8, 10. Therefore, the answer is 3.

---

**Source:** https://onlinejudge.u-aizu.ac.jp/problems/ITP1_3_D
