# Maximum Profit (最大の利益)

## Problem Statement

In FX trading, you can make profit from exchange rate differences by exchanging currencies between different countries. For example, if you buy 1000 dollars when 1 dollar = 100 yen, and then sell when the price changes to 1 dollar = 108 yen, you can make a profit of (108 yen - 100 yen) × 1000 dollars = 8000 yen.

Given the price Rt of a certain currency at time t (t=0,1,2,...,n-1), find the maximum value of the price difference Rj - Ri (where j > i).

## Input

The first line contains an integer n.
The next n lines contain integers Rt (t=0,1,2,...,n-1) in order.

## Output

Output the maximum value on one line.

## Constraints

- 2 ≤ n ≤ 200,000
- 1 ≤ Rt ≤ 10^9

## Sample Input 1

```
6
5
3
1
3
4
3
```

## Sample Output 1

```
3
```

**Explanation:** The maximum profit is obtained by buying at price 1 (index 2) and selling at price 4 (index 4), resulting in a profit of 4 - 1 = 3.

## Sample Input 2

```
3
4
3
2
```

## Sample Output 2

```
-1
```

**Explanation:** Since prices are strictly decreasing, the best we can do is buy at price 3 and sell at price 2, resulting in a loss of 2 - 3 = -1.

## Note

Source: https://onlinejudge.u-aizu.ac.jp/problems/ALDS1_1_D
