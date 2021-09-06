# Algorithms

## Divide And Conquer:
### Explain
divide a big problem into many small problems, then fix every small problems, finally combine the results.
### Examples
1. Binary search in a sorted array;
2. Calculate Fibonacci numbers by calculate matrics (1 1, 1 0) of power(n) (f(n+1) fn, fn, f(n-1));
3. calculate X power(n)
4. Merge Sort:  
	a. divide into 2 parts, 
	b. sort each of them, 
	c. then merge;
5. Quick Sort: 
	a. select a random pivot index, 
	b. swap pivot and start, 
	c. compare start+x to pivot, keep index of element smaller then pivot, 
	d. swap start+x to the smaller area, swap pivot and smaller boundary, 
	e. now we found the pivot location in array; 
	f. then quick sort left and right sub-array.
	
6. Find x th min / max element in array: 
	a. take advantage of quick sort - find pivot location, 
	b. if that is x, then we found that, otherwise found in left or right part with quick sort again;
	
## Dynamic Programming

### Define
a problem can be calculated from it sub-problems (from one or many of them): store a result set of its sub-problems, then calculte the results base on these sub-problems, so we dont need to calculate again and again on those sub-problems.

### Examples
1. Fibonacci from 0/1 to n;
2. Cut rod problem: 钢条切割问题；
3. 上台阶问题：n 个台阶，一次只有1-x步走法，有多少种走法。

## Hash Algorithms
