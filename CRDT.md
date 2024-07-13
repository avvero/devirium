From [[Jzhao]]

CRDT stands for conflict-free/commutative/convergent replicated data type. CRDTs are a family of data structures that are designed to be replicated across multiple computers without needing to worry about conflicts when people write data to the same place. If you’ve ever had to deal with a nasty git merge conflict, you know how painful these can be to resolve.

[Building a BFT JSON CRDT](https://jzhao.xyz/posts/bft-json-crdt#when-should-we-use-strong-eventual-consistency-over-linearizability)#

>I’ll note here that the term ‘conflict-free’ is a little misleading. It’s not that conflict doesn’t ever occur, but rather that CRDTs are always able to determine how to resolve the conflict up front (without user intervention)

-
#algorithm #distributed #crdt