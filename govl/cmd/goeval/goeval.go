package main

import (
	"log"

	"github.com/ecollado/go-regex-eval/internal/grammar"
)

func main() {
	log.Print("hello")

	grammar.Evaluate("${1} < 10")
	/* 	grammar.Evaluate("${10} < 10")
	   	grammar.Evaluate("${tmeout} < 10")
	   	grammar.Evaluate("len($1) > 1")
	   	grammar.Evaluate("len() > 1")
	   	grammar.Evaluate("'len'() > 1")
	   	grammar.Evaluate("len(10) > 1")
	   	grammar.Evaluate("len(item1,item2,item3) > 1")
	   	grammar.Evaluate("contains($1,localhost) > 1")
	   	grammar.Evaluate("contains(remote.host.com, [remote2.host.com,localhost,remote3.localhost.com, localhost])")
	   	grammar.Evaluate("if contains($1,localhost) : $2 else $3")
	   	grammar.Evaluate("if len($0) < 3 : $2 else $3")
	   	grammar.Evaluate("$1 < 5 and $1 > 0")
	   	grammar.Evaluate("$1 < 5 or $1 > 10")
	   	grammar.Evaluate("$1 < 5 and ($2 > 10 or $3 < 20)") */

}
