use std::time::Instant;

#[cfg(target_arch = "x86_64")]
#[inline]
pub fn rdtsc() -> u64 {
    // SAFETY: inline x86_64 rdtsc; portable only on x86_64.
    unsafe { core::arch::x86_64::_rdtsc() }
}

#[cfg(not(target_arch = "x86_64"))]
#[inline]
pub fn rdtsc() -> u64 {
    0u64 // not supported on this arch
}

/// Return resident set size in KB (best-effort).
pub fn current_rss_kb() -> Option<usize> {
    // Linux: parse /proc/self/status for VmRSS
    #[cfg(target_os = "linux")]
    {
        if let Ok(mut f) = std::fs::File::open("/proc/self/status") {
            let mut s = String::new();
            if f.read_to_string(&mut s).is_ok() {
                for line in s.lines() {
                    if line.starts_with("VmRSS:") {
                        // format: "VmRSS:\t  1234 kB"
                        let parts: Vec<&str> = line.split_whitespace().collect();
                        if parts.len() >= 2 {
                            if let Ok(kb) = parts[1].parse::<usize>() {
                                return Some(kb);
                            }
                        }
                    }
                }
            }
        }
        None
    }

    // Other Unixes: use getrusage ru_maxrss (note: ru_maxrss might be KB or bytes depending on platform)
    #[cfg(all(unix, not(target_os = "linux")))]
    {
        use libc::{RUSAGE_SELF, getrusage, rusage};
        unsafe {
            let mut usage: rusage = std::mem::zeroed();
            if getrusage(RUSAGE_SELF, &mut usage) == 0 {
                // On many unixes ru_maxrss is in kilobytes.
                return Some(usage.ru_maxrss as usize);
            }
        }
        None
    }

    #[cfg(windows)]
    {
        // Not implemented for Windows in this helper
        None
    }
}

/// Convenience: measure wall-clock and cycles and return a tuple.
pub fn time_and_cycles<F, R>(f: F) -> (R, std::time::Duration, u64)
where
    F: FnOnce() -> R,
{
    let start_cycles = rdtsc();
    let start = Instant::now();
    let r = f();
    let dur = start.elapsed();
    let end_cycles = rdtsc();
    let cycles = end_cycles.saturating_sub(start_cycles);
    (r, dur, cycles)
}
