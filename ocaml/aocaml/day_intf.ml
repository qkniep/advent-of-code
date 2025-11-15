module type DAY = sig
  val year : int
  val day : int
  val name : string

  val read_input : unit -> string list

  val solve_part1 : string list -> int
  val solve_part2 : string list -> int

  (* Optional: multiple variants per part *)
  val solve_part1_variants : (string * (string list -> int)) list
  val solve_part2_variants : (string * (string list -> int)) list
end
