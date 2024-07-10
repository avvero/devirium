# Longest increasing subsequence

> Лучшее решение - nlogn Patience sorting, работает за счет того, что гарантирует минимальный набор "стопок", через которые можно составить последовательность, ответ будет равен числу "стопок".

Мое решение с n^2:

```java
public int lengthOfLIS(int[] nums) {
    int result = 1;
    int[] weight = new int[nums.length];
    for (int i = nums.length - 1; i >= 0; i--) {
        weight[i] = 1;
        for (int j = i + 1; j < nums.length; j++) {
            if (nums[i] < nums[j]) {
                weight[i] = Math.max(weight[i], weight[j] + 1);
                result = Math.max(result, weight[i]);
            }
        }
    }
    return result;
}
```

Есть решение с nlogn с использование подхода [Patience sorting](https://en.wikipedia.org/wiki/Patience_sorting). 
> The number of piles is the length of the longest subsequence. For more info see [Princeton lecture](https://www.cs.princeton.edu/courses/archive/spring13/cos423/lectures/LongestIncreasingSubsequence.pdf) [[Theory of Algorithms]]

Нашел хорошее объяснение тут - https://www.youtube.com/watch?v=22s1xxRvy28&t=61s. 

Лучшее решение в leetcode
```java
int[] tails = new int[nums.length];
int size = 0;
for (int x : nums) {
    int i = 0, j = size;
    while (i != j) {
        int m = (i + j) / 2;
        if (tails[m] < x)
            i = m + 1;
        else
            j = m;
    }
    tails[i] = x;
    if (i == size) ++size;
}
return size;
```

Лучшее решение в leetcode2
```java
public int lengthOfLIS(int[] nums) {            
    int[] dp = new int[nums.length];
    int len = 0;

    for(int x : nums) {
        int i = Arrays.binarySearch(dp, 0, len, x);
        if(i < 0) i = -(i + 1);
        dp[i] = x;
        if(i == len) len++;
    }

    return len;
}
```

Одно из предложений к комментах
```java
public int lengthOfLIS(int[] nums) {
    List<Integer> piles = new ArrayList<>(nums.length);
    for (int num : nums) {
        int pile = Collections.binarySearch(piles, num);
        if (pile < 0) pile = ~pile;
        if (pile == piles.size()) {
            piles.add(num);
        } else {
            piles.set(pile, num);
        }
    }
    return piles.size();
}
```

#algorithm
#draft