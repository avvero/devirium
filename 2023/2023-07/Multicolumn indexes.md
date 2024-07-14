A multicolumn B-tree index can be used with query conditions that involve any subset of the index's columns, but the index is most efficient when there are constraints on the leading (leftmost) columns. 
Multicolumn indexes should be used sparingly. In most situations, an index on a single column is sufficient and saves space and time. 

Если мультиколоночный индекс используется, можно искать по первому набору полей.

https://www.postgresql.org/docs/current/indexes-multicolumn.html

#postgresql #index