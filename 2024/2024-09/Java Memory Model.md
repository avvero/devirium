Java Memory Model определяет правила взаимодействия потоков с памятью, устанавливая принципы видимости, порядка и атомарности операций. Он описывает, как изменения в памяти, сделанные одним потоком, становятся видимыми другим потокам.

### **Основные концепции JMM**

#### **a. Атомарность**

- **Описание:** Операция является атомарной, если она выполняется как единое неделимое действие.
- **Пример атомарных операций в Java:**
  - Чтение и запись примитивных типов данных размером до 32 бит (`int`, `float`, `boolean`, `char`, `short`, `byte`).
  - Операции с `long` и `double` **не гарантированно атомарны** в Java до версии 1.5, но начиная с Java 1.5, они стали атомарными.

#### **b. Видимость**

- **Описание:** Обеспечивает, что изменения, сделанные одним потоком, становятся видимыми другим потокам.
- **Проблема видимости:** Без правильной синхронизации один поток может не увидеть изменения, сделанные другим потоком.

#### **c. Упорядочение (Ordering)**

- **Описание:** Определяет порядок, в котором выполняются операции.
- **Переупорядочивание инструкций:** Компилятор и процессор могут менять порядок выполнения инструкций для оптимизации, что может привести к непредсказуемому поведению в многопоточной среде.

#### **d. Отношение "Happens-Before"**

- **Описание:** Модель, определяющая порядок видимости операций между потоками.
- **Примеры "happens-before":**
  - Блок `synchronized`: выход из синхронизированного блока одним потоком "happens-before" входа в тот же блок другим потоком.
  - Запись в переменную `volatile` "happens-before" последующего чтения этой переменной другим потоком.

### **Механизмы синхронизации в JMM**

#### **a. `volatile`**

- **Гарантии:**
  - Видимость изменений между потоками.
  - Запрет переупорядочивания операций с этой переменной.
- **Использование:** Для переменных, которые читаются и записываются несколькими потоками, и где операции над ними атомарны.

#### **b. `synchronized`**

- **Гарантии:**
  - Эксклюзивный доступ к блоку кода или методу.
  - Видимость и атомарность операций внутри синхронизированного блока.
- **Использование:** Когда требуется предотвратить состояние гонки и обеспечить атомарность составных операций.

#### **c. `final`**

- **Гарантии:** Объекты, на которые ссылаются `final` переменные, видны другим потокам после создания объекта.
- **Использование:** Для неизменяемых объектов и обеспечения безопасной публикации объектов.

### **4. Проблемы многопоточности и JMM**

#### **a. Состояние гонки (Race Condition)**

- **Описание:** Происходит, когда несколько потоков одновременно обращаются к общим данным, и результат зависит от порядка выполнения операций.
- **Решение:** Использование синхронизации для контроля доступа к общим ресурсам.

#### **b. Видимость изменений**

- **Описание:** Поток может не видеть изменения переменной, сделанные другим потоком.
- **Решение:** Использование `volatile` или синхронизации.

#### **c. Переупорядочивание инструкций**

- **Описание:** Компилятор или процессор может менять порядок выполнения инструкций, что может привести к неожиданному поведению.
- **Решение:** Использование `volatile`, `synchronized`, которые устанавливают барьеры памяти.

### **5. Примеры и пояснения**

#### **Двойная проверка синхронизации (Double-Checked Locking)**

- **Проблема:** В прошлом, из-за переупорядочивания инструкций, паттерн двойной проверки не работал корректно.
- **Решение:** Использование `volatile` с переменной-синглтоном.

```java
public class Singleton {
    private static volatile Singleton instance;

    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized(Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

### **6. Рекомендации для подготовки**

- **Изучите Java Memory Model (JSR-133):** Понимание официальной спецификации поможет глубже понять принципы работы.
- **Практикуйте написание многопоточных программ:** Реализуйте примеры с использованием `volatile`, `synchronized`, `Locks`.
- **Изучите проблемы многопоточности:** Понимайте причины и способы решения состояний гонки, проблем видимости и переупорядочивания.
- **Ознакомьтесь с высокоуровневыми конструкциями из `java.util.concurrent`:** `ReentrantLock`, `Semaphore`, `CountDownLatch`, `CyclicBarrier`, `ConcurrentHashMap` и др.

#java #memory