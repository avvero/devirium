Попробовал и вот что получилось - фигня, семантика стала сложнее. Нужно каждую строку завернуть и между ними поставить равно, хотя json большие

```groovy
package pw.avvero.spring.sandbox.jsonassert

import com.fasterxml.jackson.databind.ObjectMapper
import org.skyscreamer.jsonassert.JSONAssert
import spock.lang.Specification

import static pw.avvero.spring.sandbox.jsonassert.AssertableJson.json

class JsonAssertTests extends Specification {

    def "Compare two jsons with JSONAssert"() {
        expect:
        JSONAssert.assertEquals(expected, actual, false)
    }

    def "Compare two jsons with JSONAssert"() {
        expect:
        json(expected) == json(actual)
    }

    def "Compare two jsons with JSONAssert"() {
        expect:
        JSONAssert.assertEquals("""{
          "model": "gpt-4",
          "messages": [{
            "role": "user", 
            "content": "Hi"
          }]
        }""", """{
          "model": "gpt-4",
          "messages": [{
            "role": "user", 
            "content": "Hi!"
          }]
        }""", false)
    }

    def "Compare two jsons with JSONAssert"() {
        expect:
        json("""{
          "model": "gpt-4",
          "messages": [{
            "role": "user", 
            "content": "Hi"
          }]
        }""") == json("""{
          "model": "gpt-4",
          "messages": [{
            "role": "user", 
            "content": "Hi!"
          }]
        }""")
    }
}
```

```java
package pw.avvero.spring.sandbox.jsonassert;

import lombok.RequiredArgsConstructor;
import lombok.SneakyThrows;
import org.skyscreamer.jsonassert.JSONAssert;

import java.util.Objects;

@RequiredArgsConstructor
public class AssertableJson {

    private final String value;

    public static AssertableJson json(String value) {
        return new AssertableJson(value);
    }

    @SneakyThrows
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) {
            throw new UnsupportedOperationException("Incorrect type, expected " + getClass());
        }
        AssertableJson that = (AssertableJson) o;
        JSONAssert.assertEquals(this.value, that.value, false);
        return true;
    }

    @Override
    public int hashCode() {
        return Objects.hash(value);
    }
}
```

#json #assert #idea
#draft