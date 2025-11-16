open Core

open Aocaml.Input
open Aocaml.Day_intf

type direction = Up | Down | Left | Right

let direction_of_char = function
  | 'U' -> Up
  | 'D' -> Down
  | 'L' -> Left
  | 'R' -> Right
  | c -> invalid_arg @@ Printf.sprintf "invalid direction: %c" c

let directions_of_string str =
  str |> String.to_list |> List.map ~f:direction_of_char

let move (x, y) dir =
  let (nx, ny) = match dir with
  | Up -> (x, y - 1)
  | Down -> (x, y + 1)
  | Left -> (x - 1, y)
  | Right -> (x + 1, y)
  in
  if abs nx < 2 && abs ny < 2 then (nx, ny) else (x, y)

let move_2 (x, y) dir =
  let (nx, ny) = match dir with
  | Up -> (x, y - 1)
  | Down -> (x, y + 1)
  | Left -> (x - 1, y)
  | Right -> (x + 1, y)
  in
  if abs nx + abs ny <= 2 then (nx, ny) else (x, y)

let digit_of_position (x, y) =
  match (x, y) with
  | -1, -1 -> 1
  | 0, -1 -> 2
  | 1, -1 -> 3
  | -1, 0 -> 4
  | 0, 0 -> 5
  | 1, 0 -> 6
  | -1, 1 -> 7
  | 0, 1 -> 8
  | 1, 1 -> 9
  | _ -> invalid_arg "invalid position"

let digit_of_position_2 (x, y) =
  match (x, y) with
  | 0, -2 -> 1
  | -1, -1 -> 2
  | 0, -1 -> 3
  | 1, -1 -> 4
  | -2, 0 -> 5
  | -1, 0 -> 6
  | 0, 0 -> 7
  | 1, 0 -> 8
  | 2, 0 -> 9
  | -1, 1 -> 0xA
  | 0, 1 -> 0xB
  | 1, 1 -> 0xC
  | 0, 2 -> 0xD
  | _ -> invalid_arg "invalid position"

module Day02 : DAY = struct
  let name = "Bathroom Security"

  type input = direction list list
  type output = int

  let parse_input raw = List.map ~f:directions_of_string @@ lines raw

  let string_of_output hex_code = Printf.sprintf "%X" hex_code

  let solve_part1 input =
    let code, _ = List.fold_left
      ~f:(fun (code, pos) directions -> 
        let new_pos = List.fold_left ~f:move ~init:pos directions in
        let digit = digit_of_position new_pos in
        (code * 16 + digit, new_pos)
      )
      ~init:(0, (0, 0))
      input
    in
    code

  let solve_part2 input =
    let code, _ = List.fold_left
      ~f:(fun (code, pos) directions -> 
        let new_pos = List.fold_left ~f:move_2 ~init:pos directions in
        let digit = digit_of_position_2 new_pos in
        (code * 16 + digit, new_pos)
      )
      ~init:(0, (0, 0))
      input
    in
    code

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
