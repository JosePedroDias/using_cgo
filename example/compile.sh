#!/bin/bash

gcc -c -o example.o example.c
gcc -shared -o libexample.so example.o
cp libexample.so ..
