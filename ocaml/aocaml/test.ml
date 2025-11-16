open Core
open Day_intf
open Input
open Timing

let test_day (module D : DAY) year day solutions =
  Printf.printf "%d Day %d: %s\n%!" year day D.name;

  let raw = load_input_file_tests year day in
  let input = D.parse_input raw in

  let test_variants part variants solution =
    let num_variants = List.length variants in
    List.iteri
      ~f:(fun i (name, f) ->
        let res, t = time (fun () -> f input) in
        let out = D.string_of_output res in
        assert (String.equal out solution);
        let t = time_to_str t in
        let () =
          if String.is_empty name then
            Printf.printf " |> Part %d \x1b[0;32mPASSED\x1b[0m \x1b[90m%s\x1b[0m\t%s\n%!" part t out
          else if i < num_variants - 1 then
            Printf.printf "    ├ %s \x1b[0;32mPASSED\x1b[0m \x1b[90m%s\x1b[0m\t%s\n%!" name t out
          else Printf.printf "    └ %s \x1b[0;32mPASSED\x1b[0m \x1b[90m%s\x1b[0m\t%s\n%!" name t out
        in
        ())
      variants
  in

  let part1_solution, part2_solution =
    match solutions with
    | [s1; s2] -> s1, s2
    | _ -> failwith "should have 2 solutions (part 1 and part 2)"
  in

  let variants = [ ("", D.solve_part1) ] @ D.solve_part1_variants in
  test_variants 1 variants part1_solution;
  let variants = [ ("", D.solve_part2) ] @ D.solve_part2_variants in
  test_variants 2 variants part2_solution
