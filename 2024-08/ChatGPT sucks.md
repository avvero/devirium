Если спросить chatgpt проверить на корректность ваше решение проблемы https://leetcode.com/problems/maximum-subsequence-score, то он уверенно сообщит, что нужно делать так
```java
class Solution {
    public long maxScore(int[] nums1, int[] nums2, int k) {
        Pair[] pairs = new Pair[nums1.length];
        for (int i = 0; i < nums1.length; i++) {
            pairs[i] = new Pair(nums1[i], nums2[i]);
        }
        Arrays.sort(pairs, (a, b) -> b.n2 - a.n2);  // Sort pairs by n2 in descending order
        long max = 0;
        PriorityQueue<Integer> queue = new PriorityQueue<>(); // Min-heap
        long queueSum = 0;
        for (int i = 0; i < pairs.length; i++) {
            queue.add(pairs[i].n1);
            queueSum += pairs[i].n1;
            if (queue.size() > k) {
                queueSum -= queue.poll();
            }
            if (queue.size() == k) {
                max = Math.max(max, queueSum * pairs[i].n2);
            }
        }
        return max;
    }
    private static class Pair {
        int n1;
        int n2;
        public Pair(int n1, int n2) {
            this.n1 = n1;
            this.n2 = n2;
        }
    }
}
```
, что может вызвать ваше недоумение, ведь блок
```java
queue.add(pairs[i].n1);
queueSum += pairs[i].n1;
if (queue.size() > k) {
    queueSum -= queue.poll();
}
if (queue.size() == k) {
    max = Math.max(max, queueSum * pairs[i].n2);
}
```
не верный, ведь в требованиях указано, что мы должны посчитать сумму левых значений (n1) K пар и умножить на минимальное правое (n2) значение из K пар.
```
It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).
```
, но в предложенном решении добавленное `queue.add(pairs[i].n1)` тут же может быть удалено с `queue.poll()` и тогда требование нарушится. 

ChatGPT винить сложно, ведь почти все решения предложенные на LeetCode не корректны, но все успешно проходят тесты. А он на них и учился.

#ai