use std::time::Duration;

pub mod graph;
pub mod grid;
pub mod input;

pub fn format_duration(dt: Duration) -> String {
    let nanos = dt.as_nanos();

    if nanos < 1_000 {
        format!("{nanos} ns")
    } else if nanos < 1_000_000 {
        format!("{:.1} µs", nanos as f64 / 1_000.0)
    } else if nanos < 1_000_000_000 {
        format!("{:.1} ms", nanos as f64 / 1_000_000.0)
    } else {
        format!("{:.1} s", nanos as f64 / 1_000_000_000.0)
    }
}
