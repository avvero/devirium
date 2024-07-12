Из [[Learning PostgreSQL]]

>In the PostgreSQL renaming conventions, the suffixes for unique and normal indexes are _key and _idx respectively.
_fkey для внешних ключей

```sql 
ALTER TABLE employee ADD CONSTRAINT supervisor_id_fkey FOREIGN KEY
(supervisor_id) REFERENCES employee(employee_id);
```

#postgresql #index 