# About

This is a simple program written in Go which is intended
to act as keyboard input to a dialog windows, text
file, etc. It can ready either from a text file or
from clipboard input.

## Initial Use Case

I wrote this because I got a hardware device that can simulate text input as a 
HID device and I needed a way to quickly add text macros. Unfortunately, the 
UI provided is sub-optimal for large blocks of text and using their SDK would
require more work compared to hacking this together.

## Hacking

This program requires [Robot Go](github.com/go-vgo/robotgo). Follows its 
instructions for getting it installed. You will not be able to cross-compile
due to the underlying C library dependencies.
