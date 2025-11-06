#!/usr/bin/env bash
# Usage: ./scripts/run.sh rust 2015 day01 part1

lang=$1 year=$2 day=$3 part=$4

case $lang in
  go)      (cd go && go run ./$year/$day/*.go < ../data/$year/inputs/$day.txt) ;;
  ocaml)   (cd ocaml && dune exec ${project}/$1.exe) ;;
  rust)    (cd rust && cargo run -- "$@") ;;
  zig)     (cd zig && zig build run-${project}-$1) ;;
  *) echo "Unknown language: $lang" ;;
esac
