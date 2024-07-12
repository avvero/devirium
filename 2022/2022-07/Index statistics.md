From the [[Learning PostgreSQL]]

The first view is pg_stat_all_indexes, which gives the statistics about the index usage
The function pg_index_size can be used with the pg_size_pretty function to get the index size in a human-readable form, as follows:
```bash
car_portal=# SELECT pg_size_pretty(pg_indexes_size ('car_portal_app.account_pkey'));
```

#postgresql #index