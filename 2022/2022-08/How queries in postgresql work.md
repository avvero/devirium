>This node and some others do not store rows at all, but rather just deliver and forget them immediately. Other nodes, such as sorting, may potentially need to store vast amounts of data at a time. To deal with that, a work_mem memory chunk is allocated in backend memory. Its default size sits at a conservative 4MB limit; when the memory runs out, excess data is sent to a temporary file on-disk.

Resource - https://postgrespro.com/blog/pgsql/5969262

#postgresql #database #select