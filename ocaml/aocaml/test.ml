open Day_intf
open Input
open Timing

let test_day (module D : DAY) year day solutions =
  let part1_solution, part2_solution =
    match solutions with
    | [s1; s2] -> s1, s2
    | _ -> failwith "should have 2 solutions (part 1 and part 2)"
  in

  Printf.printf "%d Day %d: %s\n%!" year day D.name;
  
  let t0 = Core_unix.gettimeofday () in
  let raw = load_input_file_tests year day in
  let input = D.parse_input raw in
  let output = D.solve_part1 input in
  let out = D.string_of_output output in
  assert (String.equal out part1_solution);
  let t1 = Core_unix.gettimeofday () in
  let dt = time_to_str (t1 -. t0) in
  Printf.printf " |> Part 1 \x1b[0;32mPASSED\x1b[0m \x1b[90m%s\x1b[0m\n%!" dt;
  
  let t0 = Core_unix.gettimeofday () in
  let output = D.solve_part2 input in
  let out = D.string_of_output output in
  assert (String.equal out part2_solution);
  let t1 = Core_unix.gettimeofday () in
  let dt = time_to_str (t1 -. t0) in
  Printf.printf " |> Part 2 \x1b[0;32mPASSED\x1b[0m \x1b[90m%s\x1b[0m\n%!" dt;
