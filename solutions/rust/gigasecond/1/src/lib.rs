use time::{Duration, PrimitiveDateTime as DateTime};

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    let gigaseconds = Duration::seconds(1_000_000_000);

    // Add the gigasecond duration to the start time
    start + gigaseconds
}
