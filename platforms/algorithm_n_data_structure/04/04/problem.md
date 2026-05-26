# Doubly Linked List

**時間制限**: 1 sec  
**メモリ制限**: 131,072 KB

## 問題文

以下の命令を受けつける双方向連結リストを実装してください。

- `insert x`: 連結リストの先頭にキー x を持つ要素を継ぎ足す。
- `delete x`: キー x を持つ最初の要素を連結リストから削除する。そのような要素が存在しない場合は何もしない。
- `deleteFirst`: リストの先頭の要素を削除する。
- `deleteLast`: リストの末尾の要素を削除する。

## 入力

入力は以下の形式で与えられます。

```
n
command1
command2
...
commandn
```

最初の行に命令数 n が与えられます。続く n 行に命令が与えられます。命令は上記4つの命令のいずれかです。キーは整数とします。

## 出力

全ての命令が終了した後の、連結リスト内のキーを順番に出力してください。連続するキーは1つの空白文字で区切って出力してください。

## 制約

- 命令数は 2,000,000 を超えない。
- delete 命令の回数は 20 を超えない。
- 0 ≤ キーの値 ≤ 10^9
- 命令の過程でリストの要素数は 10^6 を超えない。
- delete, deleteFirst, または deleteLast 命令が与えられるとき、リストには1つ以上の要素が存在する。

## 入力例1

```
7
insert 5
insert 2
insert 3
insert 1
delete 3
insert 6
delete 5
```

## 出力例1

```
6 1 2
```

## 入力例2

```
9
insert 5
insert 2
insert 3
insert 1
delete 3
insert 6
delete 5
deleteFirst
deleteLast
```

## 出力例2

```
1
```
