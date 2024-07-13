Не рабочий, но сохраню
```java
class Solution {
    public void rotate(int[] nums, int k) {
        if (nums.length < 2) return;
        k = k % nums.length;
        // 3
        // [1,2,3,4,5,6,7] /7
        // [5,2,3,4,1,6,7] /7
        // [5,6,3,4,1,2,7] /7
        // [5,6,7,4,1,2,3] /7
        // [5,6,7,1,4,2,3] /7
        // [5,6,7,1,2,4,3] /7
        // [5,6,7,1,2,3,4] /7
        // 4
        // [1,2,3,4,5,6,7] /7
        // [4,2,3,1,5,6,7] /7
        // [4,5,3,1,2,6,7] /7
        // [4,5,6,7,2,3,1] /7
        // [4,5,6,7,1,3,2] /7
        // [4,5,6,7,1,2,3] /7
        // 5
        // [1,2,3,4,5,6,7] /7
        // [3,2,1,4,5,6,7] /7
        // [3,4,1,2,5,6,7] /7
        // [3,4,5,2,1,6,7] /7
        // [3,4,5,6,1,2,7] /7
        // [3,4,5,6,7,2,1] /7
        
        // [-1,-100,3,99]  2
        // [99,-1,-100,3]  2
        // [3, 99,-1,-100] 2
        
        // [1,2,3,4,5,6] 6
        // [6,2,3,4,5,1] 
        // [6,2,3,4,5,1] 
        
        for (int i = 0; i < nums.length; i++) {
            int p = nums.length - k + i;
            if (p < nums.length) {
                int b = nums[p];
                nums[p] = nums[i];
                nums[i] = b;
            }
        }
        //if (k == 1) return; // buffer is enough
        int ks = nums.length % (nums.length - k);
        for (int i = nums.length - ks; i < nums.length; i++) {
            int p = i - 1;
            if (p < nums.length) {
                int b = nums[i];
                nums[i] = nums[p];
                nums[p] = b;
            }
        }
    }
}
```

#algorithm #leetcode