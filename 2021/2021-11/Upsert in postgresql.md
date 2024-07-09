Update в postgresql это по факту вставка плюс удаление за счёт его механизма работы.

[We know](https://www.postgresql.org/docs/current/routine-vacuuming.html#VACUUM-FOR-SPACE-RECOVERY) that [[Postgresql]] does not update a table row in place. Rather, it writes a new version of the row (the [[Postgresql]] term for a row version is “tuple”) and leaves the old row version in place to serve concurrent read requests. `VACUUM` later removes these “dead tuples”.

https://www.cybertec-postgresql.com/en/is-update-the-same-as-delete-insert-in-postgresql

#postgresql #database #upsert
#draft