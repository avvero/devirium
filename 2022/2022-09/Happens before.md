Java defines a happens-before relationship as follows.

> Two actions can be ordered by a _happens-before_ relationship. If one action _happens-before_ another, then the first is visible to and ordered before the second.

According to this, if there is a happens-before relationship between a write and read operation, the results of a write by one thread are guaranteed to be visible to a read by another thread. Therefore, we will be able to maintain the memory consistency if we are able to have the happens-before relationship between our actions.

Нашел тут - [](https://medium.com/@kasunpdh/handling-java-memory-consistency-with-happens-before-relationship-95ddc837ab13)

#java #happens_before #concurrency 