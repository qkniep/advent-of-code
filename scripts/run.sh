#!/usr/bin/env bash
# Usage: ./scripts/run.sh rust 2015 day01 part1

case $lang in
  go)      (cd go/$project && go run ./$(echo "$@" | tr ' ' '/')) ;;
  ocaml)   (cd ocaml && dune exec ${project}/$1.exe) ;;
  rust)    (cd rust/$project && cargo run -- "$@") ;;
  zig)     (cd zig && zig build run-${project}-$1) ;;
  *) echo "Unknown language: $lang" ;;
esac
