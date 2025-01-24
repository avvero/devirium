If the code in question has any of the following qualities, prefer a property over a function:
- Doesn’t throw exceptions
- Is cheap to calculate (or cached on the first run)
- Returns the same result across multiple invocations if the object state hasn’t changed
 
Otherwise, consider using a function instead.

#kotlin