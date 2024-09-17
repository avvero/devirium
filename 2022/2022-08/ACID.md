## Atomicity 

Транзакция рассматривается как неделимая единица работы, которая либо выполняется полностью, либо не выполняется вовсе.

## Consistency

Транзакция должна переводить базу данных из одного согласованного состояния в другое, не нарушая при этом целостности данных.

Из [[Designing Data Intensive Applications]]: «The idea of ACID consistency is that you have certain statements about your data ([[Invariant]]) that must always be true—for example, in an accounting system, credits and debits across all accounts must always be balanced. If a transaction starts with a database that is valid according to these invariants, and any writes during the transaction preserve the validity, then you can be sure that the invariants are always satisfied.»

## Isolation

Одновременные транзакции не должны влиять друг на друга. Результат выполнения транзакций должен быть таким же, как если бы они выполнялись последовательно.

## Durability

После успешного завершения транзакции ее результаты должны быть надежно сохранены, даже в случае сбоев, например, потери питания или системных ошибок.

#acid #database #practice