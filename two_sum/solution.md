# Solution:

1. Create a map with index type and integer value with `container` name.

2. Do a loop on the `nums` variable.

3. Subtract each value from the `nums` loop with the `target`.

4. The results of the reduction above are checked in the map, whether there is an index in the map that is the same as the result of the reduction.

* 4.a. If it doesn't exist then, enter the result of the reduction above into the map to become the index and the value is the index of the nums loop.

* 4.b. If there is, return the value of the map above and also the index of the current loop when the subtraction result matches the index of the map container

The solution above has complexity **O<sup>(n)</sup>**. Because just do one checking in looping.