While Postgres has the ability to create multi-column indexes, it’s important to understand when it makes sense to do so. The Postgres query planner has the ability to combine and use multiple single-column indexes in a multi-column query by performing a bitmap index scan. In general, you can create an index on every column that covers query conditions and in most cases Postgres will use it, so make sure to benchmark and justify the creation of a multi-column index before you create one. As always, indexes come with a cost, and multi-column indexes can only optimize the queries that reference the columns in the index in the same order, while multiple single column indexes provide performance improvements to a larger number of queries.

[](https://devcenter.heroku.com/articles/postgresql-indexes)

#postgresql #index
#draft