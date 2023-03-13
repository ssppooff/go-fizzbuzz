# go-fizzbuzz
## Rules
From [Using FizzBuzz to Find Developers who Grok Coding](https://imranontech.com/2007/01/24/using-fizzbuzz-to-find-developers-who-grok-coding/) by Imran Ghory, January 24, 2007, 
> In this game a group of children sit around in a group and say each number in sequence, except if the number is a multiple of three (in which case they say “Fizz”) or five (when they say “Buzz”). If a number is a multiple of both three and five they have to say “Fizz-Buzz”.

>> Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”.

## Versions
- [v1.3.0](https://github.com/ssppooff/go-fizzbuzz/releases/tag/v1.3.0): Like v1.0.0, but using `switch` instead of `else if`
- [v1.2.0](https://github.com/ssppooff/go-fizzbuzz/releases/tag/v1.2.0): Use string concatenation, print that string otherwise print the number.
- [v1.1.0](https://github.com/ssppooff/go-fizzbuzz/releases/tag/v1.1.0): Print `Fizz` and/or `Buzz` without newline, use boolean flag to print only newline or number.
- [v1.0.0](https://github.com/ssppooff/go-fizzbuzz/releases/tag/v1.0.0): Use third condition `i%3 == 0 && i%5 == 0`.
