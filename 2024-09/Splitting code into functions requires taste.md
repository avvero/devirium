From [[TigerStyle]].

There are many ways to cut a wall of code into chunks of 70 lines, but only a few splits will feel right. Some rules of thumb:
- Good function shape is often the inverse of an hourglass: a few parameters, a simple return type, and a lot of meaty logic between the braces.
- [[Centralize control flow]].
- [[Centralize state manipulation]].

#design #clean #code