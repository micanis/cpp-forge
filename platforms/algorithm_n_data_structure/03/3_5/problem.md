# 安定なソート (Stable Sort)

## Problem Statement

トランプのカードを整列しましょう。ここでは、４つの絵柄(S, H, C, D)と９つの数字(1, 2, ..., 9)から構成される計 36 枚のカードを用います。例えば、ハートの 8 は"H8"、ダイヤの 1 は"D1"と表します。

バブルソート及び選択ソートのアルゴリズムを用いて、与えられた N 枚のカードをそれらの数字を基準に昇順に整列するプログラムを作成してください。アルゴリズムはそれぞれ以下に示す疑似コードに従うものとします。数列の要素は 0 オリジンで記述されています。

```
BubbleSort(C, N)
  for i = 0 to N-1
    for j = N-1 downto i+1
      if C[j].value < C[j-1].value
        C[j] と C[j-1] を交換

SelectionSort(C, N)
  for i = 0 to N-1
    minj = i
    for j = i to N-1
      if C[j].value < C[minj].value
        minj = j
    C[i] と C[minj] を交換
```

また、各アルゴリズムについて、与えられた入力に対して安定な出力を行っているか報告してください。ここでは、同じ数字を持つカードが複数ある場合それらが入力に出現する順序で出力されることを、「安定な出力」と呼ぶことにします。（※常に安定な出力を行うソートのアルゴリズムを安定なソートアルゴリズムと言います。）

## Input Format

1 行目にカードの枚数 N が与えられます。
2 行目に N 枚のカードが与えられます。各カードは絵柄と数字のペアを表す２文字であり、隣合うカードは１つの空白で区切られています。

## Output Format

1 行目に、バブルソートによって整列されたカードを順番に出力してください。隣合うカードは１つの空白で区切ってください。
2 行目に、この出力が安定か否か（Stable またはNot stable）を出力してください。
3 行目に、選択ソートによって整列されたカードを順番に出力してください。隣合うカードは１つの空白で区切ってください。
4 行目に、この出力が安定か否か（Stable またはNot stable）を出力してください。

## Constraints

- 1 ≤ N ≤ 36
- 時間制限 : 1 sec
- メモリ制限 : 131072 KB

## Sample Input/Output

### Example 1

**Input:**
```
5
H4 C9 S4 D2 C3
```

**Output:**
```
D2 C3 H4 S4 C9
Stable
D2 C3 S4 H4 C9
Not stable
```

**Explanation:**
- Bubble Sort: Produces the cards sorted by value (2, 3, 4, 4, 9). The two 4s appear in their original order (H4 before S4), making it stable.
- Selection Sort: Also produces sorted cards, but the two 4s are reversed (S4 before H4) compared to the original input, making it not stable.

### Example 2

**Input:**
```
2
S1 H1
```

**Output:**
```
S1 H1
Stable
S1 H1
Stable
```

**Explanation:**
- Both algorithms produce the same stable output because the two cards have the same value, and they maintain their original order.
