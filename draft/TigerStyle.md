Style of coding with philosophy:
- On Simplicity And Elegance ?
- zero technical debt policy - "do it right the first time"
- Safety
- - Do not use recursion
- - Use only a minimum of excellent abstractions
- - Use explicitly-sized types like u32
- - Restrict the length of function bodies to reduce the probability of poorly structured code. We enforce a hard limit of 70 lines per function.
- - [[Splitting code into functions requires taste]]
- - - [[Centralize control flow]]
- - - [[Centralize state manipulation]]


Details - https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md.

Includes [[Negative Space Programming]].

#design #clean #code #draft