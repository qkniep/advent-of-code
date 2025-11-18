#[macro_export]
macro_rules! measure {
    // simple style: measure!( "label", || { expr } );
    ($label:expr, $body:expr) => {{
        use std::io::Write;
        // reset allocator counters
        $crate::reset_alloc_counters();

        // sample RSS before
        let before_rss = $crate::measure_helpers::current_rss_kb();

        // run and time (and cycles if supported)
        let (res, duration, cycles) = $crate::measure_helpers::time_and_cycles(|| {
            $body
        });

        // gather allocation stats
        let alloc_count = $crate::get_alloc_count();
        let alloc_bytes = $crate::get_alloc_bytes();

        // sample RSS after
        let after_rss = $crate::measure_helpers::current_rss_kb();

        // Print a structured single-line summary (also return the result)
        // We prefer println! here so this works inside bench harnesses too.
        println!(
            "[measure] {}: time = {:?}, cycles = {}, allocs = {}, bytes = {} B, rss_before = {}, rss_after = {}",
            $label,
            duration,
            cycles,
            alloc_count,
            alloc_bytes,
            before_rss.map(|k| format!("{} KB", k)).unwrap_or("N/A".to_string()),
            after_rss.map(|k| format!("{} KB", k)).unwrap_or("N/A".to_string()),
        );

        // Additionally, flush stdout so results appear promptly in bench output
        let _ = std::io::stdout().flush();

        res
    }};
}
