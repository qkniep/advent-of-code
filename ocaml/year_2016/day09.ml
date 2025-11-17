open Core
open Aocaml.Day_intf

let rec decompressed_length input =
  if String.length input = 0 then 0
  else
    let (read, output) = if Char.equal input.[0] '(' then
      let i = String.index_exn input ')' in
      let str = String.sub input ~pos:1 ~len:(i - 1) in
      match String.split ~on:'x' str with
      | [ len; rep ] ->
          let len = int_of_string len in
          let rep = int_of_string rep in
          (i + len + 1, len * rep)
      | _ -> invalid_arg ("invalid marker: " ^ str)
    else (1, 1) in
    let new_str = (String.sub input ~pos:read ~len:(String.length input - read)) in
    output + decompressed_length new_str

let decompressed_length_v2 _input = 0

module Day09 : DAY = struct
  let name = "Explosives in Cyberspace"

  type input = string
  type output = int

  let parse_input raw = String.strip raw
  let string_of_output = string_of_int
  let solve_part1 input = decompressed_length input
  let solve_part2 input = decompressed_length_v2 input

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 9" =
  let open Aocaml.Test in
  let solutions = [ ""; "" ] in
  test_day (module Day09 : DAY) 2016 9 solutions
