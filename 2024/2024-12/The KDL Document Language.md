https://kdl.dev

KDL is a small, document language with XML-like node semantics that looks like you're invoking a bunch of CLI commands. It's meant to be used both as a serialization format and a configuration language, much like JSON, YAML, or XML. It looks like this:

```json
package {
  name my-pkg
  version "1.2.3"

  dependencies {
    // Nodes can have standalone values as well as
    // key/value pairs.
    lodash "^3.2.1" optional=#true alias=underscore
  }

  scripts {
    // "Raw" and dedented multi-line strings are supported.
    message """
      hello
      world
      """
    build #"""
      echo "foo"
      node -c "console.log('hello, world!');"
      echo "foo" > some-file.txt
      """#
  }

  // `\` breaks up a single node across multiple lines.
  the-matrix 1 2 3 \
             4 5 6 \
             7 8 9

  // "Slashdash" comments operate at the node level,
  // with just `/-`.
  /-this-is-commented {
    this entire node {
      is gone
    }
  }
}
```

#language