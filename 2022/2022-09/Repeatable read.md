That is a simple task that issue two reads from table T, with a delay of 1 minute between them. 

-   under READ COMMITTED, the second SELECT may return _any_ data. A concurrent transaction may update the record, delete it, insert new records. The second select will always see the _new_data.
-   under REPEATABLE READ the second SELECT is guaranteed to display at least the rows that were returned from the first SELECT _unchanged_. New rows may be added by a concurrent transaction in that one minute, but the existing rows cannot be deleted nor changed.
-   under SERIALIZABLE reads the second select is guaranteed to see _exactly_ the same rows as the first. No row can change, nor deleted, nor new rows could be inserted by a concurrent transaction.

Repeatable read is a higher isolation level, that in addition to the guarantees of the read committed level, it also guarantees that any data read cannot change, if the transaction reads the same data again, it will find the previously read data in place, unchanged, and available to read.

The next isolation level, serializable, makes an even stronger guarantee: in addition to everything repeatable read guarantees, it also guarantees that no new data can be seen by a subsequent read.

#transaction #isolation #repeatable_read #database 