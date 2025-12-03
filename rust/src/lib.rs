mod alloc_tracker;
pub mod core;
pub mod day_run_macro;
pub mod day_tests_macro;
pub mod measure_helpers;
pub mod measure_macro;
pub mod years;

pub use self::years::*;
pub use alloc_tracker::*;

pub const FIRST_YEAR: u16 = 2015;
pub const LAST_YEAR: u16 = 2025;

#[global_allocator]
static GLOBAL_ALLOC: alloc_tracker::CountingAllocator = alloc_tracker::CountingAllocator;
