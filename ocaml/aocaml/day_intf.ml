module type DAY = sig
  val name : string

  (* Required: parsed input type *)
  type input
  (** Input type is a list of strings *)

  (* Required: parse input and solve parts 1 and 2 *)
  val parse_input : string -> input
  val solve_part1 : input -> int
  val solve_part2 : input -> int

  (* Optional: additional implementations per part *)
  val solve_part1_variants : (string * (input -> int)) list
  val solve_part2_variants : (string * (input -> int)) list
end
