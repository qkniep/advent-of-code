open Aocaml.Day_intf
open Aocaml.Util

let run_day (module D : DAY) year day =
  Printf.printf "⁘⁙⁘⁙⁘ Advent of Code %d ⁘⁙⁘⁙⁘\n%!" year;
  Printf.printf "Day %d: %s\n%!" day D.name;

  let input, t = time (fun () -> D.read_input ()) in
  Printf.printf "  - Read input in %s\n%!" (time_to_str t);

  let run_variants part variants =
    List.iter (fun (name, f) ->
      (* let _res, t = time (label ^ "/" ^ name) (fun () -> f input) in *)
      let res, t = time (fun () -> f input) in
      let t = time_to_str t in
      let () = if name = "" then
        Printf.printf "  - Part %d %s %d\n%!" part t res
      else
        Printf.printf "  - Part %d (%s) %s %d\n%!" part name t res
      in
      ()) variants
  in

  run_variants 1 ([("", D.solve_part1)] @ D.solve_part1_variants);
  run_variants 2 ([("", D.solve_part2)] @ D.solve_part2_variants)

let () =
  (* default to day 1 of 2016 *)
  let year = if Array.length Sys.argv > 1 then int_of_string Sys.argv.(1) else 2016 in
  let day = if Array.length Sys.argv > 2 then int_of_string Sys.argv.(2) else 1 in

  let day_module = match year, day with
    | 2016, 1 -> (module Year_2016.Day01.Day01 : DAY)
    | 2016, 2 -> (module Year_2016.Day02.Day02 : DAY)
    | _ -> failwith (Printf.sprintf "Day %d of year %d not implemented\n" day year)
  in
  run_day day_module year day
