```sql
SELECT now() - query_start as s, pid, backend_start, xact_start, query_start, state, query
FROM pg_stat_activity
WHERE state = 'idle in transaction'
order by s desc;
```
```sql
SELECT 
  now() - query_start as s, 
    pid,
    usename,
    application_name,
    client_addr,
    backend_start,
    state,
    wait_event,
    query,
    xact_start,
    query_start,
    backend_xid,
    backend_xmin,
    txid_current_if_assigned() AS current_tx
FROM pg_stat_activity
WHERE application_name = 'r2dbc-postgresql'
ORDER BY backend_start;
```

```sql
SELECT 
    pid, 
    usename,
    client_addr,
    state,
    query_start,
    state_change,
    wait_event,
    backend_type,
    query
FROM pg_stat_activity
ORDER BY query_start;
```

#postgresql #sql