open Aocaml.Input
open Aocaml.Day_intf

module Day02 : DAY = struct
  let name = "Bathroom Security"

  type input = string list

  let parse_input str = lines str

  let solve_part1 input = List.length input

  let solve_part2 input =
    List.fold_left (fun acc line -> acc + String.length line) 0 input

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
