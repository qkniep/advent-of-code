open Core
open Aocaml.Day_intf

let has_n_zeroes hex n =
  String.sub hex ~pos:0 ~len:n
  |> String.for_all ~f:(fun c -> Char.equal c '0')

let hash_counter_hex input counter =
    let input = String.concat [ input; string_of_int counter ] in
    Md5.to_hex (Md5.digest_string input)

let next_key input start =
  Sequence.unfold ~init:start ~f:(fun i -> Some (i, i+1))
  |> Sequence.find_map ~f:(fun i ->
    let hex = hash_counter_hex input i in
    if has_n_zeroes hex 5 then Some (String.get hex 5, i+1) else None)
  |> Option.value_exn

let password_sequence input =
  Sequence.unfold ~init:0 ~f:(fun cnt ->
    let ch, next_cnt = next_key input cnt in
    Some (ch, next_cnt)
  )

module Day05 : DAY = struct
  let name = "How About a Nice Game of Chess?"

  type input = string
  type output = string

  let parse_input raw = String.strip raw
  let string_of_output output = output

  let solve_part1 input =
    Sequence.take (password_sequence input) 8
    |> Sequence.to_list
    |> String.of_char_list

  let solve_part2 _input = ""

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "hex" =
  let input = "abc" in
  let i = 5017308 in
  let input = String.concat [ input; string_of_int i ] in
  let hex = Md5.to_hex (Md5.digest_string input) in
  Printf.printf "HEX TEST: %s\n" hex;
  let (c, new_cnt) = next_key "abc" 0 in
  assert (Char.equal c '1');
  assert (new_cnt = 3231929);
  let (c, new_cnt) = next_key "abc" 3231929 in
  assert (Char.equal c '1');
  assert (new_cnt = 3231929);
  let (c, _new_cnt) = next_key "abc" 5017308 in
  assert (Char.equal c '8')

let%test_unit "2016 day 5" =
  let open Aocaml.Test in
  let solutions = [ "173787"; "548" ] in
  test_day (module Day05 : DAY) 2016 5 solutions
