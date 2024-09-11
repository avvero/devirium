involves placing constraints and assertions throughout your code to explicitly define invalid states and conditions. By doing this, you ensure that the code fails fast and early, preventing unintended behaviors from propagating through the system. 

From [[TigerStyle]]:
> Assert all function arguments and return values, pre/postconditions and invariants

Example
```Zig
const std = @import("std");
const assert = std.debug.assert;

fn calculateArea(width: i32, height: i32) i32 {
    assert(width > 0);
    assert(height > 0);
    return width * height;
}
```

One of the most compelling advocates for negative space programming is [[Joran Dirk Greef]].

Negative space programming is part of [[TigerStyle]].

For further reading - https://double-trouble.dev/post/negativ-space-programming

#design #clean #code