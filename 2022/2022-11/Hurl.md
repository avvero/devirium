# Hurl

Hurl is a command line tool that runs HTTP requests defined in a simple plain text format.

```
# Get home:
GET https://example.org

HTTP/1.1 200
[Captures]
csrf_token: xpath "string(//meta[@name='_csrf_token']/@content)"

# Do login!
POST https://example.org/login?user=toto&password=1234
X-CSRF-TOKEN: {{csrf_token}}

HTTP/1.1 302
```

Site - https://hurl.dev/

#tool #cli #rest #http #request
#draft