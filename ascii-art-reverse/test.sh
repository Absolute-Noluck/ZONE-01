#!/bin/bash

declare -A validations

validations["test/example00.txt"]="Hello World"
validations["test/example01.txt"]="123"
validations["test/example02.txt"]="#=\["
validations["test/example03.txt"]="(somthing&234)"
validations["test/example04.txt"]="abcdefghijklmnopqrstuvwxyz"
validations["test/example05.txt"]="\!\" #$%&'()*+,-./"
validations["test/example06.txt"]=":;<=>?@"
validations["test/example07.txt"]="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
validations["test/example08.txt"]="queLQUE"
validations["test/example09.txt"]="quel   123"
validations["test/example10.txt"]="&'-"
validations["test/example11.txt"]="lowerUPPER   1209"




for f in test/*.txt; do
    validation=${validations[$f]}
    if [ -z "$validation" ]; then
        echo "no validation for $f"
    else
        out=$(go run . --reverse=$f)
        [ "${out%$'\n'}" = "$validation" ] && echo "test ok for  $f" || echo "test not ok for $f"
        echo "filecontent: $out"
        
    fi
done