# TDD, Where Did It All Go Wrong

https://www.youtube.com/watch?v=EZ05e7EMOLM&ab_channel=DevTernityConference

Simon VV
2 года назад (изменено)
Notes I made during this presentation (just a dump, might be useful to some):

- Test requirements, not low level
- Test public API. Given when then
- Test the exports from a module
- Focus on higher-level
- Test modules, not class
- Refactoring is needed to see what is implementation and what is exports from module
- Test behaviours
- Think about your code as an api
- Test the abstraction, not the implementation
- Test are isolated and with shared fixture (to run quickly)
- Red-green-refactor (go fast to working code)
- No new tests during refactoring
- Heavy coupling is the problem with all software
- Thin public api 
- Refactoring = changing internals
- Patterns in the refactoring
- If you're not really sure, write tests for implementation (delete the tests)
- Not classes, behaviours
- Don't isolate classes in testing
- Private methods (these are implementation details)

#tdd #code #testing #tests
#draft