# fibobuzz

  * fibonize - generate fibonacci sequence (to stdout)
  * increment - start counting from 1 (to stdout)
  * fizzbuzz - (reads stdin) fizz multiples of three, buzz multiples of five, fizzbuzz multiples of both three and five.
  * fizzbuzz also checks for prime

### usage

  1. compile using 'make'
  2. run fizzbuzz using 'make fizzbuzz'
  3. run fibobuzz using 'make run'

### example

```

# reads standard input until EOF
echo 42 | ./bin/fizzbuzz
echo 7 | ./bin/fizzbuzz
echo 300000 | ./bin/fizzbuzz
echo 300000000000000000000000000000000000000000 | ./bin/fizzbuzz
printf '30\n31' | ./bin/fizzbuzz
# try even larger numbers!



```



### information

### author

aerth <aerth@riseup.net>

### known bugs

  * non-numeric input (such as "42 hello 42") will result in a FizzBuzz

```
 echo 7 bug 7 | ./bin/fizzbuzz
 echo 8 bug 7 | ./bin/fizzbuzz
 echo 9 bug 7 | ./bin/fizzbuzz
 echo 10 bug 7 | ./bin/fizzbuzz
```


