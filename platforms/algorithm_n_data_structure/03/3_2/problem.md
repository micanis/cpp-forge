# Insertion Sort (挿入ソート)

## Problem Statement

Insertion Sort (挿入ソート) is one of the most natural and intuitive sorting algorithms, similar to how one might sort a hand of playing cards. When arranging cards from a deck by holding them in one hand and sorting from left to right in ascending order, you take one card at a time and insert it into the appropriate position in the already-sorted sequence of cards.

The insertion sort algorithm works as follows:

```
# Sort N elements in array A in ascending order using insertion sort
def insertion_sort(N, A):
    for i in range(1, N):            # Start from index 1 (0-indexed)
        key = A[i]                   # Hold the element to be inserted
        j = i - 1                    # Start from one position before current
        while j >= 0 and A[j] > key: # Find position to insert key
            A[j + 1] = A[j]          # Move elements larger than key backward by 1
            j -= 1
        A[j + 1] = key               # Insert key into the vacant position
```

Create a program that sorts a sequence A containing N elements in ascending order using insertion sort. Follow the pseudocode above to implement the algorithm. To verify the algorithm's operation, output the array at each computation step (the arrangement immediately after input and immediately after processing each i).

## Input

The first line contains an integer N representing the length of the sequence. The second line contains N integers separated by spaces.

## Output

The output consists of N lines. Output the intermediate result of the insertion sort at each computation step on one line. Array elements should be separated by a single space. Note that a Presentation Error will occur if you include extra spaces after the last element or unnecessary newlines.

## Constraints

- 1 ≤ N ≤ 100
- 0 ≤ A elements ≤ 1,000

## Sample Input 1

```
6
5 2 4 6 1 3
```

## Sample Output 1

```
5 2 4 6 1 3
2 5 4 6 1 3
2 4 5 6 1 3
2 4 5 6 1 3
1 2 4 5 6 3
1 2 3 4 5 6
```

## Sample Input 2

```
3
1 2 3
```

## Sample Output 2

```
1 2 3
1 2 3
1 2 3
```
