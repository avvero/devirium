In more simple terms:

-   **Pure deterministic function**: The output is based entirely, and only, on the input values and nothing else: there is no other (hidden) input or state that it relies on to generate its output. There are no side-effects or other output.
-   **Impure deterministic function**: As with a deterministic function that is a pure function: the output is based entirely, and only, on the input values and nothing else: there is no other (hidden) input or state that it relies on to generate its output - **however** there is other output (side-effects).
-   **Idempotency**: The practical definition is that you can safely call the same function multiple times without fear of negative side-effects. More formally: there are no _changes of state_ between subsequent identical calls

[Нашел тут](https://stackoverflow.com/questions/40296211/what-is-the-difference-between-an-idempotent-and-a-deterministic-function)

|              | Pure deterministic                    | Impure deterministic                  | Pure Nondeterministic                | Impure Nondeterministic              | Idempotent                                  |
| ------------ | ------------------------------------- | ------------------------------------- | ------------------------------------ | ------------------------------------ | ------------------------------------------- |
| Input        | Only parameter arguments (incl. this) | Only parameter arguments (incl. this) | Parameter arguments and hidden state | Parameter arguments and hidden state | Any                                         |
| Output       | Only return value                     | Return value or side-effects          | Only return value                    | Return value or side-effects         | Any                                         |
| Side-effects | None                                  | Yes                                   | None                                 | Yes                                  | After 1st call: Maybe. After 2nd call: None |
| SQL Example  | UCASE                                 | CREATE TABLE                          | GETDATE                              | DROP TABLE                           |
| C# Example   | String.IndexOf                        | DateTime.Now                          | Directory.Create(String)             |                                      |

**Чем меньше побочных эффектов имеет функция, тем лучше.**

Когда функция детерминированная и не имеет побочных эффектов, мы называем её "**чистой**" функцией. Чистые функции:

-   проще читать
-   проще отлаживать
-   проще тестировать
-   не зависят от порядка, в котором они вызываются
-   просто запустить параллельно (одновременно)

Чистые функции независимы от времени. Недетерминизм и побочные эффекты добавляют понятие времени. Если функция зависит от чего-то, что может случиться, а может не случиться и меняет что-то за пределами своих границ, то она неожиданно становится зависимой от времени.

#code #design #function #deterministic #idempotency