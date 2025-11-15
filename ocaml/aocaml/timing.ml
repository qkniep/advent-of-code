let time f =
  let t0 = Unix.gettimeofday () in
  let res = f () in
  let t1 = Unix.gettimeofday () in
  (* Printf.printf "Time: %.3f ms\n%!" (t1 -. t0); *)
  (res, t1 -. t0)

let benchmark ~reps f =
  let t0 = Unix.gettimeofday () in
  for _ = 1 to reps do
    ignore (f ())
  done;
  let t1 = Unix.gettimeofday () in
  (* Printf.printf "Average time: %.3f ms\n%!" ((t1 -. t0) /. float_of_int reps) *)
  (t1 -. t0) /. float_of_int reps

let time_to_str secs =
  match secs with
  | s when s < 1e-6 -> Printf.sprintf "%.1f ns" (s *. 1e9)
  | s when s < 1e-3 -> Printf.sprintf "%.1f μs" (s *. 1e6)
  | s when s < 1.0 -> Printf.sprintf "%.1f ms" (s *. 1e3)
  | s -> Printf.sprintf "%.1f s" s
