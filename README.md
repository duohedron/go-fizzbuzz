# Fizz Buzz - For smart Go programmers

[![Build Status](https://travis-ci.org/duohedron/go-fizzbuzz.svg?branch=master)](https://travis-ci.org/duohedron/go-fizzbuzz)
[![Go Report Card](https://goreportcard.com/badge/github.com/duohedron/go-fizzbuzz)](https://goreportcard.com/report/github.com/duohedron/go-fizzbuzz)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

So you're in a job interview, and the interviewer asks you to write a 'fizz buzz generator'. What would a good developer do? 
A novice developer writes 80 lines of code, an experienced developer would write 20, but you? You're smart. You install it using `go get`, and write three lines of code.

## Installation

```
go get github.com/duohedron/go-fizzbuzz
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/duohedron/go-fizzbuzz"
)

func main() {
	for s := range fizzbuzz.FizzBuzz(20) {
		fmt.Println(s)
	}
}
```