Из [[Learning PostgreSQL]]

Ошибка дизайна когда много атрибутов и они не заполняются данными.

> Null saturation: By nature, some applications have sparse data, such as medical applications. Imagine a relation called diagnostics which has hundreds of attributes for symptoms like fever, headache, sneezing, and so on. Most of them are not valid for certain diagnostics, but they are valid in general. This could be modeled by utilizing complex data types like JSON, or by using vertical modeling like entity-attribute-value (EAV).

#postgresql #design #database 