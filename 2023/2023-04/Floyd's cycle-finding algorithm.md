x mod L = Z

<p align="center">
  <img src="Floyd's cycle-finding algorithm.png" />
</p>

Взято из видео https://www.youtube.com/watch?v=9YTjXqqJEFE&t=758s

Использовал для решения https://leetcode.com/problems/find-the-duplicate-number
```java
public int findDuplicate(int[] nums) {
    int slow = 0;
    int fast = 0;
    do {
        slow = nums[slow];
        fast = nums[nums[fast]];
    } while (slow != fast);
    fast = 0;
    while (slow != fast) {
        slow = nums[slow];
        fast = nums[fast];
    }
    return fast;
}
```

#algorithm #loop #joma #leetcode