open Core
open Aocaml.Day_intf

type token =
  | Literal of char
  | Marker of { len : int; rep : int; rep_str : string }

(** Returns the token and the number of characters read. *)
let parse_token input =
  if Char.equal input.[0] '(' then
    let i = String.index_exn input ')' in
    let str = String.sub input ~pos:1 ~len:(i - 1) in
    match String.split ~on:'x' str with
    | [ len; rep ] ->
        let len = int_of_string len in
        let rep = int_of_string rep in
        let rep_str = String.sub input ~pos:(i + 1) ~len in
        (Marker { len; rep; rep_str }, i + len + 1)
    | _ -> invalid_arg ("invalid marker: " ^ str)
  else (Literal input.[0], 1)

let rec decompressed_length input =
  if String.is_empty input then 0
  else
    let token, read = parse_token input in
    let output =
      match token with Literal _ -> read | Marker { len; rep; _ } -> len * rep
    in
    output + decompressed_length (String.drop_prefix input read)

let rec decompressed_length_v2 input =
  if String.is_empty input then 0
  else
    let token, read = parse_token input in
    let output =
      match token with
      | Literal _ -> read
      | Marker { rep; rep_str; _ } -> rep * decompressed_length_v2 rep_str
    in
    output + decompressed_length_v2 (String.drop_prefix input read)

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
  let solutions = [ "115118"; "11107527530" ] in
  test_day (module Day09 : DAY) 2016 9 solutions
