module Input = struct
  let read_lines () =
    let rec loop acc =
      match input_line stdin with
      | line -> loop (line :: acc)
      | exception End_of_file -> List.rev acc
    in
    loop []

  let read_all () =
    let buf = Buffer.create 1024 in
    (try
       while true do
         Buffer.add_string buf (input_line stdin);
         Buffer.add_char buf '\n'
       done
     with End_of_file -> ());
    Buffer.contents buf

  let split s ~on = String.split_on_char on s

  let words s =
    s |> String.split_on_char ' '
      |> List.filter (fun x -> x <> "")

  let ints s =
    s |> words
      |> List.map int_of_string

  let comma_ints s =
    s |> split ~on:','
      |> List.map int_of_string

  let digits_of_string s =
    s |> String.to_seq
      |> Seq.map (fun c -> Char.code c - Char.code '0')
      |> List.of_seq

  let int_grid lines =
    List.map digits_of_string lines
end
