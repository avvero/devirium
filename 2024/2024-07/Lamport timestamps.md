From [[Jzhao]], relates to [[Happens before]]

Track logical time rather than actual wall time, meaning we count number of events that have occurred rather than seconds elapsed. This timestamp is just a simple counter.

All nodes start their counter at 0
- Every time we perform an operation locally, we increment our counter by one
- Every time we broadcast a message to our peers, we attach this counter
- Every time we receive a message, we set our timer to `max(self.seq, incoming.seq) + 1`
If `a.seq > b.seq` then event a must have happened after event b. However, if `a.seq == b.seq`, we cannot be sure which event came first. This means that Lamport timestamps only give us a partial ordering, which means two identical sequence numbers might not correspond to the same unique event. 

#algorithm