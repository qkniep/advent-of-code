open Aocaml.Day_intf
open Aocaml.Input
open Aocaml.Timing

let run_day (module D : DAY) year day =
  Printf.printf "⁘⁙⁘⁙⁘ AOCaml %d ⁘⁙⁘⁙⁘\n%!" year;
  Printf.printf "Day %d: %s\n%!" day D.name;

  let raw_input, t = time (fun () -> load_input_file year day) in
  Printf.printf " -> Read input \x1b[90m%s\x1b[0m\n%!" (time_to_str t);

  let input, t = time (fun () -> D.parse_input raw_input) in
  Printf.printf " -> Parse input \x1b[90m%s\x1b[0m\n%!" (time_to_str t);

  let run_variants part variants =
    let num_variants = List.length variants in
    List.iteri
      (fun i (name, f) ->
        let res, t = time (fun () -> f input) in
        let t = time_to_str t in
        let () =
          if name = "" then
            Printf.printf " |> Part %d \x1b[90m%s\x1b[0m\t%d\n%!" part t res
          else if i < num_variants - 1 then
            Printf.printf "    ├ %s \x1b[90m%s\x1b[0m\t%d\n%!" name t res
          else Printf.printf "    └ %s \x1b[90m%s\x1b[0m\t%d\n%!" name t res
        in
        ())
      variants
  in

  run_variants 1 ([ ("", D.solve_part1) ] @ D.solve_part1_variants);
  run_variants 2 ([ ("", D.solve_part2) ] @ D.solve_part2_variants)

let () =
  (* default to day 1 of 2016 *)
  let year, day =
    if Array.length Sys.argv > 2 then
      (int_of_string Sys.argv.(1), int_of_string Sys.argv.(2))
    else (2016, 1)
  in

  let day_module =
    match (year, day) with
    | 2016, 1 -> (module Year_2016.Day01.Day01 : DAY)
    | 2016, 2 -> (module Year_2016.Day02.Day02 : DAY)
    | _ ->
        failwith (Printf.sprintf "Day %d of year %d not implemented\n" day year)
  in
  run_day day_module year day
