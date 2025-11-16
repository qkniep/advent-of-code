open Core
open Aocaml.Input
open Aocaml.Day_intf

type room = { words : string list; sector : int; checksum : string }

let room_of_string str =
  let parts = split ~on:'-' str in
  let last_part = List.last_exn parts in
  let sector, checksum =
    let parts =
      String.split
        (String.sub ~pos:0 ~len:(String.length last_part - 1) last_part)
        ~on:'['
    in
    match parts with
    | [ sector; checksum ] -> (sector, checksum)
    | _ -> failwith @@ Printf.sprintf "invalid last part: %s" last_part
  in
  let words =
    match List.drop_last parts with
    | Some words -> words
    | _ -> failwith "no words"
  in
  { words; sector = int_of_string sector; checksum }

let top5_letters counter =
  let count_list = Array.to_list counter in
  let sorted =
    List.mapi
      ~f:(fun i count -> (count, Char.of_int_exn (Char.to_int 'a' + i)))
      count_list
    |> List.sort ~compare:(fun (n1, c1) (n2, c2) ->
           match Int.descending n1 n2 with
           | 0 -> Char.compare c1 c2
           | cmp -> cmp)
  in
  List.take sorted 5 |> List.map ~f:snd |> String.of_char_list

let real_room room =
  let counter = Array.create ~len:26 0 in
  List.iter
    ~f:(fun word ->
      List.iter
        ~f:(fun c ->
          let pos = Char.to_int c - 97 in
          counter.(pos) <- counter.(pos) + 1)
        (String.to_list word))
    room.words;
  let top5 = top5_letters counter in
  String.equal top5 room.checksum

let shift_char shift c =
  Char.of_int_exn
    (((Char.to_int c - Char.to_int 'a' + shift) mod 26) + Char.to_int 'a')

let decrypt_name room =
  let shift = room.sector mod 26 in
  List.map
    ~f:(fun word -> String.map ~f:(fun c -> shift_char shift c) word)
    room.words
  |> String.concat ~sep:" "

module Day04 : DAY = struct
  let name = "Security Through Obscurity"

  type input = room list
  type output = int

  let parse_input raw = List.map ~f:room_of_string @@ lines raw
  let string_of_output = string_of_int

  let solve_part1 input =
    List.filter ~f:real_room input
    |> List.sum (module Int) ~f:(fun room -> room.sector)

  let solve_part2 input =
    let room =
      List.filter ~f:real_room input
      |> List.find_exn ~f:(fun room ->
             String.equal (decrypt_name room) "northpole object storage")
    in
    room.sector

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 4" =
  let open Aocaml.Test in
  let solutions = [ "173787"; "548" ] in
  test_day (module Day04 : DAY) 2016 4 solutions
