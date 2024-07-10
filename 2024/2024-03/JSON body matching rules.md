# JSON body matching rules

from [[Pact]] https://docs.pact.io/implementation_guides/jvm/matching

The following method is used to determine if two bodies match:


>1. If both the actual body and expected body are empty, the bodies match.
>2. If the actual body is non-empty, and the expected body empty, the bodies match.
>3. If the actual body is empty, and the expected body non-empty, the bodies don't match.
>4. Otherwise do a comparison on the contents of the bodies.

Интересно тут 2, я так делаю, указывая false в `JSONAssert.assertEquals(expected, actual, false)` [[Another assert json library 2]]

#test #json #pact #spring
#draft