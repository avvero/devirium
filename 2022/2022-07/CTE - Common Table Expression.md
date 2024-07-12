CTE (Common Table Expression) is a temporary result set in SQL that you can reference within a SELECT, INSERT, UPDATE, or DELETE statement. CTEs are useful for improving the readability and maintainability of complex queries. They can also be used to break down a complex query into simpler parts.

```sql
WITH SOME AS (SELECT 1),
     SOME2 AS (SELECT 2)
SELECT * FROM SOME, SOME2;
```

#postgresql #database #cte