Guard Clause pattern.

In computer programming, a guard is a boolean expression that must evaluate to true if the program execution is to continue in the branch in question. Regardless of which programming language is used, guard code or a guard clause is a check of integrity preconditions used to avoid errors during execution

https://medium.com/lemon-code/guard-clauses-3bc0cd96a2d3

```javascript
public void run( Timer timer ) {
  if( !timer.isEnabled() )
    return;
  
  if( !timer.valid() )
    throw new InvalidTimerException();
  timer.run();
}
```

#development #design 
#draft