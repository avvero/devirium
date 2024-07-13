```java
    /**
     * Matches fieldName in string name="fieldName"
     */
    public static final Pattern FIELD_NAME_PATTERN = Pattern.compile("(?<=name=\").*?(?=\")");

    // trim all
    header = header.replace(" ", "");
    Matcher matcher = FIELD_NAME_PATTERN.matcher(header);
    return matcher.find() ? matcher.group() : null;
```

#development #java #regex