mod alloc_tracker;

pub use alloc_tracker::*;
pub mod core;
pub mod day_tests_macro;
pub mod measure_helpers;
pub mod measure_macro;

#[global_allocator]
static GLOBAL_ALLOC: alloc_tracker::CountingAllocator = alloc_tracker::CountingAllocator;
