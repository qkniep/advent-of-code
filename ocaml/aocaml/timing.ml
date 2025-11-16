open Core

let time f =
  let t0 = Core_unix.gettimeofday () in
  let res = f () in
  let t1 = Core_unix.gettimeofday () in
  (res, t1 -. t0)

let benchmark ~reps f =
  let t0 = Core_unix.gettimeofday () in
  for _ = 1 to reps do
    ignore (f ())
  done;
  let t1 = Core_unix.gettimeofday () in
  (t1 -. t0) /. Int.to_float reps

let time_to_str secs =
  let open Float.O in
  if secs < 1e-6 then Printf.sprintf "%.1f ns" (secs * 1e9)
  else if secs < 1e-3 then Printf.sprintf "%.1f μs" (secs * 1e6)
  else if secs < 1.0 then Printf.sprintf "%.1f ms" (secs * 1e3)
  else Printf.sprintf "%.1f s" secs
