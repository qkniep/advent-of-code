open Aocaml.Input
open Aocaml.Day_intf

module Day02 : DAY = struct
  let name = "Bathroom Security"

  let read_input () = Input.read_lines ()

  let solve_part1 (input : string list) = List.length input

  let solve_part2 input = List.fold_left (fun acc line -> acc + String.length line) 0 input

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
