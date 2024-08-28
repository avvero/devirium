![](TransactionalOutbox.png)

The solution is for the service that sends the message to first store the message in the database as part of the transaction that updates the business entities. A separate process then sends the messages to the message broker.

#pattern #transaction