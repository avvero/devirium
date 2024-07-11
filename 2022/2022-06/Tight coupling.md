Из [[Learning PostgreSQL]]
Ошибка дизайна когда используются общие понятие в домене, когда домен расширяется или меняется, то возникает проблема того, что общее понятие занято для частной задачи.

Tight coupling: In some cases, tight coupling leads to complex and difficult- to-change data structures. Since business requirements change with time, some requirements might become obsolete. Modeling generalization and specialization (for example a part-time student is a student) in a tightly coupled way may cause problems.

#postgresql #database #design