Оптимизация в JIT 

for (int x = 0; i < 100; i++){  
do_something(i);  
}

Без оптимизации 400 инструкций (100*4):
add data (i) - (doing actual work)  
add i +1(increment)  
cmp (compare)  
jl (jump to)

С оптимизациями 175
add data (0)  
add data (1)  
add data (2)  
add data (3)  
add i +4 (increment)  
cmp (compare)  
jl (jump to)

Статья - [](https://medium.com/@Styp/observing-java-19-jvm-optimization-with-jmh-hsdis-perfasm-part-i-e80c4907e2f9)

#java #jit
#draft