open Core

open Aocaml.Input
open Aocaml.Day_intf

let valid_triangle a b c = a + b > c && a + c > b && b + c > a

let transform (input : int list list) =
  List.chunks_of ~length:3 input
  |> List.map ~f:(fun chunks ->
         match chunks with
         | [ r1; r2; r3 ] -> (
             match (r1, r2, r3) with
             | [ r1a; r1b; r1c ], [ r2a; r2b; r2c ], [ r3a; r3b; r3c ] ->
                 [ [ r1a; r2a; r3a ]; [ r1b; r2b; r3b ]; [ r1c; r2c; r3c ] ]
             | _ -> invalid_arg "columns are not equal 3")
         | _ -> invalid_arg "rows are not multiple of 3")
  |> List.concat

module Day03 : DAY = struct
  let name = "Squares With Three Sides"

  type input = int list list
  type output = int

  let parse_input raw = List.map ~f:ints @@ lines raw

  let string_of_output = string_of_int

  let solve_part1 input =
    List.fold_left
      ~f:(fun valid sides ->
        match sides with
        | [ a; b; c ] -> valid + if valid_triangle a b c then 1 else 0
        | _ -> invalid_arg "invalid number of sides")
      ~init:0 input

  let solve_part2 input =
    let input = transform input in
    solve_part1 input

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
