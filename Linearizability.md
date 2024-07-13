From [[Designing Data Intensive Applications]]

Informally: every operation takes effect atomically sometime after it started and before it finished. All operations behave as if executed on a single copy of the data. Not to be confused with [[Serializability]].

>This is the idea behind linearizability (also known as atomic consistency, strong consistency, immediate consistency, or external consistency). The exact definition of linearizability is quite subtle, and we will explore it in the rest of this section. But the basic idea is to make a system appear as if there were only one copy of the data, and all operations on it are atomic. With this guarantee, even though there may be multiple replicas in reality, the application does not need to worry about them. In a linearizable system, as soon as one client successfully completes a write, all clients reading from the database must be able to see the value just written. 

#term