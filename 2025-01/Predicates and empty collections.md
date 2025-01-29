The function with perhaps the largest mind-bending potential is all. Regardless of the predicate, it returns true when called on an empty collection:
```
fun main() {
    println(emptyList<Int>().all { it > 42 })
    // true
}
```
This might surprise you at first, but upon further investigation, you’ll find that this is a very reasonable return value. You can’t name an element that violates the predicate, so the predicate clearly has to be true for all elements in the collection—even if there are none! This concept is known as the [[Vacuous truth]] and, in most cases, actually ends up a good fit for conditionals that should also work with empty collections.

From [[Kotlin in Action]].

#kotlin