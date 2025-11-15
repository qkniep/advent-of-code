open Core

open Aocaml.Input
open Aocaml.Day_intf

type instruction = Left of int | Right of int
type step = TurnLeft | TurnRight | Forward
module Position = struct
  module T = struct
    type t = { x : int; y : int }

    let origin = { x = 0; y = 0 }

    let compare t1 t2 =
      match Int.compare t1.x t2.x with
      0 -> Int.compare t1.y t2.y | c -> c

    let sexp_of_t t : Sexp.t = List [ Atom (string_of_int t.x); Atom (string_of_int t.y) ]
  end

  include T
  include Comparator.Make (T)
end
type orientation = North | South | East | West

let steps_of_inst = function
  | Left s -> TurnLeft :: List.init s ~f:(fun _ -> Forward) 
  | Right s -> TurnRight :: List.init s ~f:(fun _ -> Forward) 

(* module Point2d = struct *)
(*   module T = struct *)
(*     type t = { x : int; y : int } *)
(*  *)
(*     let compare t1 t2 = *)
(*       let cmp_title = Int.compare t1.x t2.x in *)
(*       if cmp_title <> 0 then cmp_title else Int.compare t1.y t2.y *)
(*  *)
(*     let sexp_of_t t : Sexp.t = List [ Atom (string_of_int t.x); Atom (string_of_int t.y) ] *)
(*   end *)
(*  *)
(*   include T *)
(*   include Comparator.Make (T) *)
(* end *)

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

let move (start : Position.t) dir steps =
  match dir with
  | North -> { start with y = start.y + steps }
  | South -> { start with y = start.y - steps }
  | East -> { start with x = start.x + steps }
  | West -> { start with x = start.x - steps }

let apply_step (start : Position.t) dir step =
  let new_dir =
    match step with
    | TurnLeft -> turn_left dir
    | TurnRight -> turn_right dir
    | Forward -> dir
  in
  let new_pos =
    match step with
    | TurnLeft -> start
    | TurnRight -> start
    | Forward -> move start dir 1
  in
  (new_pos, new_dir)

let apply_inst (start : Position.t) dir inst =
  let new_dir =
    match inst with Left _ -> turn_left dir | Right _ -> turn_right dir
  in
  let new_pos =
    match inst with Left steps | Right steps -> move start new_dir steps
  in
  (new_pos, new_dir)

let distance (pos : Position.t) = abs pos.x + abs pos.y

module Day01 : DAY = struct
  let name = "No Time for a Taxicab"

  type input = instruction list

  let parse_input str =
    let input = words str in
    List.map
      ~f:(fun inst ->
        if String.length inst < 2 then
          invalid_arg @@ Printf.sprintf "invalid instruction: %s" inst;
        let steps =
          int_of_string (String.sub inst ~pos:1 ~len:(String.length inst - 2))
        in
        match inst.[0] with
        | 'L' -> Left steps
        | 'R' -> Right steps
        | c -> invalid_arg @@ Printf.sprintf "invalid direction: %c" c)
      input

  let solve_part1 input =
    let start_dir = North in
    let start_pos = Position.origin in
    let end_pos, _ =
      List.fold_left
        ~f:(fun (pos, dir) inst -> apply_inst pos dir inst)
        ~init:(start_pos, start_dir)
        input
    in
    distance end_pos

  let solve_part2 input =
    let steps = List.concat_map ~f:steps_of_inst input in
    let start_dir = North in
    let start_pos = Position.origin in
    let _, _, end_pos, _ =
      List.fold_left
        ~f:(fun (is_done, visited, pos, dir) step -> 
          if is_done then (true, visited, pos, dir) else
            let new_pos, new_dir = apply_step pos dir step in
            let is_done = match step with Forward -> Set.mem visited new_pos | _ -> false in
            let new_visited = Set.add visited new_pos in
            (is_done, new_visited, new_pos, new_dir))
        ~init:(false, Set.empty (module Position), start_pos, start_dir)
        steps
    in
    distance end_pos

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
