open Aocaml.Input
open Aocaml.Day_intf

type instruction = Left of int | Right of int
type position = { x : int; y : int }
type orientation = North | South | East | West

let turn_left = function
  | North -> West
  | South -> East
  | East -> North
  | West -> South

let turn_right = function
  | North -> East
  | South -> West
  | East -> South
  | West -> North

let move start dir steps =
  match dir with
  | North -> { start with y = start.y + steps }
  | South -> { start with y = start.y - steps }
  | East -> { start with x = start.x + steps }
  | West -> { start with x = start.x - steps }

let apply_inst start dir inst =
  let new_dir =
    match inst with Left _ -> turn_left dir | Right _ -> turn_right dir
  in
  let new_pos =
    match inst with Left steps | Right steps -> move start new_dir steps
  in
  (new_pos, new_dir)

let distance pos = abs pos.x + abs pos.y

module Day01 : DAY = struct
  let name = "No Time for a Taxicab"

  type input = instruction list

  let parse_input str =
    let input = words str in
    List.map
      (fun inst ->
        if String.length inst < 2 then
          failwith (Printf.sprintf "Invalid instruction: %s" inst);
        let steps =
          int_of_string (String.sub inst 1 (String.length inst - 2))
        in
        match inst.[0] with
        | 'L' -> Left steps
        | 'R' -> Right steps
        | c -> failwith (Printf.sprintf "Invalid direction: %c" c))
      input

  let solve_part1 input =
    let start_dir = North in
    let start_pos = { x = 0; y = 0 } in
    let end_pos, _ =
      List.fold_left
        (fun (pos, dir) inst -> apply_inst pos dir inst)
        (start_pos, start_dir) input
    in
    distance end_pos

  let solve_part2 input =
    let start_dir = North in
    let start_pos = { x = 0; y = 0 } in
    let end_pos, _ =
      List.fold_left
        (fun (pos, dir) inst -> apply_inst pos dir inst)
        (start_pos, start_dir) input
    in
    distance end_pos

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
