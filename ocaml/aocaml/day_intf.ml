module type DAY = sig
  (** A single day's puzzle interface.  
      Each day provides a name, a parser for its input, and one or more
      solution functions for part 1 and part 2. *)

  val name : string
  (** Human-readable name of the day, e.g. ["Not Quite Lisp"].
      Should not contain the day number, year, or any newlines. *)

  type input
  (** The structured representation of that day's puzzle input.  
      Implementations are free to choose any type that best fits the puzzle. *)

  type output
  (** The solution type for part 1 and part 2.  
      Usually this will be [int], sometimes [string]. *)

  val parse_input : string -> input
  (** [parse_input raw] transforms the raw input text into the structured
      {!type:input} value consumed by solution functions.  
      Should raise an exception on malformed input. *)

  val string_of_output : output -> string
  (** [string_of_output output] transforms the {!type:output} value
      into the human-readable submission string.  
      Should not contain any whitespace. *)

  val solve_part1 : input -> output
  (** Compute the primary solution for part 1.  
      This is the default implementation used unless a variant is selected. *)

  val solve_part2 : input -> output
  (** Compute the primary solution for part 2.  
      This is the default implementation used unless a variant is selected. *)

  val solve_part1_variants : (string * (input -> output)) list
  (** Optional additional implementations of part 1.  
      Each entry is [(label, f)] where [label] is a short description
      (e.g. ["optimized"], ["naive"], ["vectorized"]) and [f] is an alternative solver.
      The empty list indicates that only the default implementation exists. *)

  val solve_part2_variants : (string * (input -> output)) list
  (** Optional additional implementations of part 2.  
      Follows the same conventions as {!val:solve_part1_variants}. *)
end
