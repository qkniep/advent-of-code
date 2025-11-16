open Core

let load_input_file year day =
  let path = Printf.sprintf "../data/%d/inputs/day%02d.txt" year day in
  In_channel.read_all path

let read_lines () =
  let rec loop acc =
    match In_channel.input_line_exn In_channel.stdin with
    | line -> loop (line :: acc)
    | exception End_of_file -> List.rev acc
  in
  loop []

let read_all () =
  let buf = Buffer.create 1024 in
  (try
     while true do
       Buffer.add_string buf (In_channel.input_line_exn In_channel.stdin);
       Buffer.add_char buf '\n'
     done
   with End_of_file -> ());
  Buffer.contents buf

let split s ~on = String.split_on_chars ~on:[ on ] s
let split_non_empty s ~on = split ~on s |> List.filter ~f:(Fn.non String.is_empty)
let lines s = split_non_empty ~on:'\n' s
let words s = split_non_empty ~on:' ' s
let ints s = s |> words |> List.map ~f:int_of_string
let comma_ints s = s |> split ~on:',' |> List.map ~f:int_of_string

let digits_of_string s =
  String.to_list s
  |> List.map ~f:(fun c -> Char.to_int c - Char.to_int '0')

let int_grid lines = List.map ~f:digits_of_string lines
