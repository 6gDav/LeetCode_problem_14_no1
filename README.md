# <a href="https://leetcode.com/problems/longest-common-prefix/description/">1. Two Sum</a>

## 📝 Description

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

## 🧠 How I solved the problem 

First of all, I implemented an early return statement if the given list is empty, it returns an empty string.
I wrote a method that returns the length of the shortest word in the list to optimize performance. I created a loop that goes only up to that length.
Within this loop, I used a variable that stores the currently indexed character of the first element. After this, I iterated through the given list and compared every single element's character at the current index with the previously created variable. If they don't match, then the code will return the result string built so far.
If the inner loop finishes successfully, it concatenates the current character to a result string.

## ➗ Complexity

* **Time complexity**: *O(n x m)* – I use to two type of loop.
* **Space complexity**: *O(1)* – No extra data structures are used.


## 📊 Benchmark

I made it in release mode for more accurate results:
```bash
go run .
```

Hardware: *Apple Mac Mini M4*

### 🤏 Small Input Test

* **Execution Time**: *52.083µs*
* **Memory Delta**:   *0 bytes*
* **Current Memory**: *12667144 bytes*

### 😖 Stress Test (Large Input)

* **Execution Time**: *7.833µs*
* **Memory Delta**:   *16 bytes*
* **Current Memory**: *12994824 bytes*