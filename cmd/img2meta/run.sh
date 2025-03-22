#!/bin/sh

input=~/Downloads/PNG_transparency_demonstration_1.png

cat "${input}" |
	./img2meta |
	jq .
