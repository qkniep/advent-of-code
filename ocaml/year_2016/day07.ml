open Core
open Aocaml.Input
open Aocaml.Day_intf

type ip_part = Supernet of string | Hypernet of string
type ip = ip_part list

let is_hypernet part = match part with Hypernet _ -> true | _ -> false
let string_of_ip_part = function Supernet s | Hypernet s -> s

let ip_of_string str =
  String.split_on_chars ~on:[ '['; ']' ] str
  |> List.mapi ~f:(fun i str ->
         if i mod 2 = 0 then Supernet str else Hypernet str)

let has_abba part =
  let str = string_of_ip_part part in
  let seq = Sequence.unfold ~init:0 ~f:(fun i -> Some (i, i + 1)) in
  Sequence.take seq (String.length str - 4 + 1)
  |> Sequence.exists ~f:(fun i ->
         Char.equal str.[i] str.[i + 3]
         && Char.equal str.[i + 1] str.[i + 2]
         && not (Char.equal str.[i] str.[i + 1]))

let supports_tls ip =
  let has_supernet_abba =
    List.filter ~f:(Fn.non is_hypernet) ip |> List.exists ~f:has_abba
  in
  let has_hypernet_abba =
    List.filter ~f:is_hypernet ip |> List.exists ~f:has_abba
  in
  has_supernet_abba && not has_hypernet_abba

let all_aba part =
  let str = string_of_ip_part part in
  let seq = Sequence.unfold ~init:0 ~f:(fun i -> Some (i, i + 1)) in
  Sequence.take seq (String.length str - 3 + 1)
  |> Sequence.filter ~f:(fun i ->
         Char.equal str.[i] str.[i + 2] && not (Char.equal str.[i] str.[i + 1]))
  |> Sequence.map ~f:(fun i ->
         String.of_char_list [ str.[i]; str.[i + 1]; str.[i + 2] ])
  |> Sequence.to_list

let bab_of_aba aba =
  let a = Char.to_string aba.[0] in
  let b = Char.to_string aba.[1] in
  String.concat [ b; a; b ]

let has_corresponding_bab part aba =
  let bab = bab_of_aba aba in
  let str = string_of_ip_part part in
  String.is_substring str ~substring:bab

let supports_ssl ip =
  List.filter ~f:(Fn.non is_hypernet) ip
  |> List.concat_map ~f:all_aba
  |> List.exists ~f:(fun aba ->
         List.filter ~f:is_hypernet ip
         |> List.exists ~f:(fun part -> has_corresponding_bab part aba))

module Day07 : DAY = struct
  let name = "Internet Protocol Version 7"

  type input = ip list
  type output = int

  let parse_input raw = lines raw |> List.map ~f:ip_of_string
  let string_of_output = string_of_int
  let solve_part1 input = List.filter ~f:supports_tls input |> List.length
  let solve_part2 input = List.filter ~f:supports_ssl input |> List.length

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 7" =
  let open Aocaml.Test in
  let solutions = [ "105"; "258" ] in
  test_day (module Day07 : DAY) 2016 7 solutions
