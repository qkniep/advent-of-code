open Core

open Aocaml.Input
open Aocaml.Day_intf

type room = { words : string list; sector : int; checksum : string }

let room_of_string str =
  let parts = split ~on:'-' str in
  let last_part = List.last_exn parts in
  let sector, checksum =
    let parts = String.split (String.sub ~pos:0 ~len:(String.length last_part - 1) last_part) ~on:'[' in
    match parts with
    | [ sector; checksum ] -> sector, checksum
    | _ -> failwith @@ Printf.sprintf "invalid last part: %s" last_part
  in
  let words = match List.drop_last parts with | Some words -> words | _ -> failwith "no words" in
  { words = words; sector = int_of_string sector; checksum }

let top5_letters counter =
  let count_list = Array.to_list counter in
  let sorted = List.mapi ~f:(fun i count -> count, Char.of_int_exn (Char.to_int 'a' + i)) count_list
    |> List.sort ~compare:(fun (n1, c1) (n2, c2) ->
        match Int.descending n1 n2 with
        | 0 -> Char.compare c1 c2
        | cmp -> cmp) in
  List.take sorted 5
    |> List.map ~f:snd
    |> String.of_char_list

let real_room room =
  let counter = Array.create ~len:26 0 in
  List.iter
    ~f:(fun word ->
      List.iter
      ~f:(fun c -> Array.set counter (Char.to_int c - 97) (Array.get counter (Char.to_int c - 97) + 1))
      (String.to_list word)
    )
    room.words;
  let top5 = top5_letters counter in
  String.equal top5 room.checksum

module Day04 : DAY = struct
  let name = "Security Through Obscurity"

  type input = room list
  type output = int

  let parse_input raw = List.map ~f:room_of_string @@ lines raw

  let string_of_output = string_of_int

  let solve_part1 input =
    List.filter ~f:real_room input
    |> List.sum (module Int) ~f:(fun room -> room.sector)

  let solve_part2 _input = 0

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end
