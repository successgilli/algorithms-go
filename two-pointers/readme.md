### Two pointers algorithm


## Strobogrammatic number

This is a number that appears the same after rotated 180º.

Examples includes 69, 88, 808, 101 etc.

### Technique

The idea here is to first identify a map of numbers that appear same as another number between 0 and 9 when rotated 180º (upside down).

They include
0 -> 0, 1 -> 1, 6 -> 9, 8 -> 8, 9 -> 6.

- Initialize two pointers at the start and end of the input
- For each number in an input, check if strobogrammatic with the other pointer.
- If not, return false since a strobogrammatic number must contain all individual constituent numbers strobogrammatic.
- replace number with the mapped rotation.
- return true at the end to show entire input is strobogrammatic. The number rotated can also be returned.
