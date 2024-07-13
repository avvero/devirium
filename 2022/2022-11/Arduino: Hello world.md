```c
int val;  // define variable val
void setup() {
  Serial.begin(9600);  // set the baud rate at 9600 .
}
void loop() {
  // read the instruction or character from PC to Arduino, and assign them to Val.
  val = Serial.read();
  // determine if the instruction or character received is “R”.
  if (val == 'R') {
    Serial.println("Hello World!");  // display“Hello World！”string.
  }
}  //
```

#arduino