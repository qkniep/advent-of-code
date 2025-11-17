open Core
open Aocaml.Input
open Aocaml.Day_intf

let font = [|
  ".##.. ###.. .##.. ###.. ####. ####. .###. #..#. .###. ..##. #..#. #.... #...# #...# .###. ###.. .##.. ###.. .#### ##### #...# #...# #...# #...# #...# ####.";
  "#..#. #..#. #..#. #..#. #.... #.... #...# #..#. ..#.. ...#. #.#.. #.... ##.## ##..# #...# #..#. #..#. #..#. #.... ..#.. #...# #...# #.#.# .#.#. #...# ...#.";
  "#..#. ###.. #.... #..#. ###.. ###.. #.... ####. ..#.. ...#. ##... #.... #.#.# #.#.# #...# #..#. #..#. #..#. .###. ..#.. #...# .#.#. #.#.# ..#.. .#.#. ..#..";
  "####. #..#. #.... #..#. #.... #.... #.### #..#. ..#.. ...#. #.#.. #.... #.#.# #.#.# #...# ###.. #..#. ###.. ....# ..#.. #...# .#.#. ##### .#.#. ..#.. .#...";
  "#..#. #..#. #..#. #..#. #.... #.... #...# #..#. ..#.. #..#. #.#.. #.... #...# #..## #...# #.... #.##. #.#.. ....# ..#.. #...# ..#.. ##.## #...# ..#.. #....";
  "#..#. ###.. .##.. ###.. ####. #.... .###. #..#. .###. .##.. #..#. ####. #...# #...# .###. #.... .##.# #..#. ####. ..#.. .###. ..#.. #...# #...# ..#.. ####.";
|]

type cmd =
  | Rect of { w : int; h : int }
  | RotateRow of { row : int; steps : int }
  | RotateColumn of { col : int; steps : int }

let parse_rect dim =
  match String.split ~on:'x' dim with
  | [ w; h ] -> Rect { w = int_of_string w; h = int_of_string h }
  | _ -> invalid_arg ("invalid rect: " ^ dim)

let parse_rotate_row idx steps =
  match String.split ~on:'=' idx with
  | [ "y"; row ] ->
      RotateRow { row = int_of_string row; steps = int_of_string steps }
  | _ -> invalid_arg ("invalid rotate row: " ^ idx)

let parse_rotate_col idx steps =
  match String.split ~on:'=' idx with
  | [ "x"; col ] ->
      RotateColumn { col = int_of_string col; steps = int_of_string steps }
  | _ -> invalid_arg ("invalid rotate col: " ^ idx)

let cmd_of_string str =
  match String.split ~on:' ' str with
  | [ "rect"; dim ] -> parse_rect dim
  | [ "rotate"; "row"; idx; "by"; steps ] -> parse_rotate_row idx steps
  | [ "rotate"; "column"; idx; "by"; steps ] -> parse_rotate_col idx steps
  | _ -> invalid_arg ("invalid cmd: " ^ str)

type screen = bool array array

let make_screen w h = Array.make_matrix ~dimx:h ~dimy:w false

let draw_rect screen w h =
  for x = 0 to w - 1 do
    for y = 0 to h - 1 do
      screen.(y).(x) <- true
    done
  done

let rotate_row screen y steps =
  let row = Array.copy screen.(y) in
  let len = Array.length row in
  for i = 0 to len - 1 do
    let x = (i + steps) % len in
    screen.(y).(x) <- row.(i)
  done

let rotate_col screen x steps =
  let len = Array.length screen in
  let col = Array.init len ~f:(fun i -> screen.(i).(x)) in
  for i = 0 to len - 1 do
    let y = (i + steps) % len in
    screen.(y).(x) <- col.(i)
  done

let run_cmd screen cmd =
  match cmd with
  | Rect { w; h } -> draw_rect screen w h
  | RotateRow { row; steps } -> rotate_row screen row steps
  | RotateColumn { col; steps } -> rotate_col screen col steps

let active_pixels screen =
  Array.fold screen ~init:0 ~f:(fun acc row ->
      Array.fold row ~init:acc ~f:(fun acc b -> if b then acc + 1 else acc))

let print_screen screen =
  Array.iter screen ~f:(fun row ->
      Array.iter row ~f:(fun b ->
          if b then print_string "#" else print_string ".");
      print_endline "")

let check_letter screen pos letter =
  let xo = pos * 5 in
  let font_xo = letter * 6 in
  let mismatch =
    List.exists (List.range 0 6) ~f:(fun y ->
        List.exists (List.range 0 5) ~f:(fun x ->
            let s = screen.(y).(xo + x) in
            let f = Char.equal font.(y).[font_xo + x] '#' in
            not (Bool.( = ) s f)))
  in
  not mismatch

let detect_letter screen pos =
  let letters = Sequence.unfold ~init:0 ~f:(fun i -> Some (i, i + 1)) in
  Sequence.take letters 26
  |> Sequence.find_map ~f:(fun letter ->
         if check_letter screen pos letter then
           Some (char_of_int (letter + Char.to_int 'a'))
         else None)
  |> Option.value_exn

module Day08 : DAY = struct
  let name = "Two-Factor Authentication"

  type input = cmd list
  type output = string

  let parse_input raw = lines raw |> List.map ~f:cmd_of_string
  let string_of_output out = out

  let width = 50
  let height = 6

  let solve_part1 input =
    let screen = make_screen width height in
    List.iter ~f:(run_cmd screen) input;
    active_pixels screen |> string_of_int

  let solve_part2 input =
    let screen = make_screen width height in
    List.iter ~f:(run_cmd screen) input;
    let positions = Sequence.unfold ~init:0 ~f:(fun i -> Some (i, i + 1)) in
    Sequence.take positions 10
    |> Sequence.map ~f:(fun i -> detect_letter screen i)
    |> String.of_sequence

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 8" =
  let open Aocaml.Test in
  let solutions = [ "115"; "efeykfrfij" ] in
  test_day (module Day08 : DAY) 2016 8 solutions
