open Aocaml.Input
open Aocaml.Day_intf

module Day01 : DAY = struct
  let year = 2016
  let day = 1
  let name = "No Time for a Taxicab"

  let read_input () = lines (load_input_file year day)

  let solve_part1 (input : string list) = List.length input

  let solve_part2 input = List.fold_left (fun acc line -> acc + String.length line) 0 input

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
