# Do not write tests for package main

Do not write tests for package main. Main is for integrating your code units (a.k.a packages) and usually would not contain unit-testable code. IF you have test-worthy code in main (for unit tests, that is; not for integration tests), consider moving it into a library package.

https://appliedgo.net/testmain/

#development #go #test
#draft