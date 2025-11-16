open Core
open Aocaml.Day_intf

let has_n_zeroes hex n =
  String.sub hex ~pos:0 ~len:n |> String.for_all ~f:(fun c -> Char.equal c '0')

let hash_counter_hex input counter =
  let input = String.concat [ input; string_of_int counter ] in
  Md5.to_hex (Md5.digest_string input)

let next_relevant_hex input start =
  Sequence.unfold ~init:start ~f:(fun i -> Some (i, i + 1))
  |> Sequence.find_map ~f:(fun i ->
         let hex = hash_counter_hex input i in
         if has_n_zeroes hex 5 then Some (hex, i + 1) else None)
  |> Option.value_exn

let password_sequence input =
  Sequence.unfold ~init:0 ~f:(fun cnt ->
      let hex, next_cnt = next_relevant_hex input cnt in
      Some ((hex.[5], hex.[6]), next_cnt))

module Day05 : DAY = struct
  let name = "How About a Nice Game of Chess?"

  type input = string
  type output = string

  let parse_input raw = String.strip raw
  let string_of_output output = output

  let solve_part1 input =
    Sequence.take (password_sequence input) 8
    |> Sequence.map ~f:fst |> Sequence.to_list |> String.of_char_list

  let solve_part2 input =
    let buf = Buffer.create 8 in
    for _ = 0 to 7 do
      Buffer.add_char buf '#'
    done;
    let buf = Bytes.of_string (Buffer.contents buf) in
    password_sequence input
    |> Sequence.map ~f:(fun (p, c) ->
           let pos = Char.to_int p - Char.to_int '0' in
           (pos, c))
    |> Sequence.filter ~f:(fun (pos, _) -> pos >= 0 && pos <= 7)
    |> Sequence.iter_until
         ~f:(fun (pos, c) ->
           let char = Bytes.get buf pos in
           if Char.equal char '#' then Bytes.set buf pos c else ();
           if Bytes.contains buf '#' then Continue ()
           else Stop (Bytes.to_string buf))
         ~finish:(fun () -> Bytes.to_string buf)

  (* no variants, just empty lists *)
  let solve_part1_variants = []
  let solve_part2_variants = []
end

let%test_unit "2016 day 5" =
  let open Aocaml.Test in
  let solutions = [ "4543c154"; "1050cbbd" ] in
  test_day (module Day05 : DAY) 2016 5 solutions
