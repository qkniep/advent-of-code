open Core
open Aocaml.Input
open Aocaml.Day_intf

let most_common_letters input =
  let len = String.length (List.hd_exn input) in
  let counts = Array.init (26 * len) ~f:(fun _ -> 0) in
  let most_common = Bytes.init len ~f:(fun _ -> 'a') in
  List.iter input ~f:(fun word ->
      String.to_sequence word
      |> Sequence.iteri ~f:(fun i c ->
             let mc_char = Bytes.get most_common i in
             let mc_idx = (26 * i) + (Char.to_int mc_char - Char.to_int 'a') in
             let mc_count = Array.get counts mc_idx in
             let c_idx = (26 * i) + (Char.to_int c - Char.to_int 'a') in
             let c_count = Array.get counts c_idx in
             Array.set counts c_idx (c_count + 1);
             if c_count + 1 > mc_count then Bytes.set most_common i c));
  most_common

let count_letters input =
  let len = String.length (List.hd_exn input) in
  let counts = Array.init (26 * len) ~f:(fun _ -> 0) in
  List.iter input ~f:(fun word ->
      String.to_sequence word
      |> Sequence.iteri ~f:(fun i c ->
             let c_idx = (26 * i) + (Char.to_int c - Char.to_int 'a') in
             let c_count = Array.get counts c_idx in
             Array.set counts c_idx (c_count + 1)));
  counts

let least_common_letters input =
  count_letters input |> Array.chunks_of ~length:26
  |> Array.map ~f:(fun arr ->
         Array.to_sequence arr
         |> Sequence.mapi ~f:(fun i x -> (i, x))
         |> Sequence.min_elt ~compare:(fun (_, x) (_, y) -> Int.compare x y)
         |> Option.value_exn |> fst)
  |> Array.map ~f:(fun i -> Char.of_int_exn (Char.to_int 'a' + i))

module Day06 : DAY = struct
  let name = "Signals and Noise"

  type input = string list
  type output = string

  let parse_input raw = lines raw
  let string_of_output output = output

  let solve_part1 input = most_common_letters input |> Bytes.to_string

  let solve_part2 input =
    least_common_letters input
    |> Array.map ~f:Char.to_string
    |> String.concat_array

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 6" =
  let open Aocaml.Test in
  let solutions = [ "tzstqsua"; "myregdnr" ] in
  test_day (module Day06 : DAY) 2016 6 solutions
