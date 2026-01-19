1. Правим compose: добавляем shared_preload_libraries и (опционально) параметры pg_stat_statements
```yaml
command:
  - "-c"
  - "max_connections=256"
  - "-c"
  - "shared_preload_libraries=pg_stat_statements"
  - "-c"
  - "pg_stat_statements.max=10000"
  - "-c"
  - "pg_stat_statements.track=all"
  - "-c"
  - "pg_stat_statements.save=on"
```

2. Запускаем
3. Выполняем
```sql
create extension if not exists pg_stat_statements;
```
4. Проверяем
```sql
select * from pg_stat_statements;
```

## Запросы

1. Топ по суммарному времени
```sql
select
  calls,
  round(total_exec_time::numeric) as total_ms,
  round(mean_exec_time::numeric, 2) as mean_ms,
  rows,
  query
from pg_stat_statements
order by total_exec_time desc
limit 10;
```

2. Самые медленные по latency
```sql
select
  calls,
  round(mean_exec_time::numeric, 2) as mean_ms,
  round(max_exec_time::numeric, 2) as max_ms,
  query
from pg_stat_statements
where calls > 10
order by mean_exec_time desc
limit 10;
```

3. Частые, но дорогие
```sql
select
  calls,
  round(total_exec_time::numeric) as total_ms,
  round((total_exec_time / calls)::numeric, 2) as avg_ms,
  query
from pg_stat_statements
where calls > 1000
order by total_exec_time desc
limit 10;
```

#postgresql