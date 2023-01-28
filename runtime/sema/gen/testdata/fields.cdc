pub struct Test {
    /// This is a test integer.
    let testInt: UInt64

    /// This is a test optional integer.
    let testOptInt: UInt64?

    /// This is a test integer reference.
    let testRefInt: &UInt64

    /// This is a test variable-sized integer array.
    let testVarInts: [UInt64]

    /// This is a test constant-sized integer array.
    let testConstInts: [UInt64; 2]

    /// This is a test parameterized-type field.
    let testParam: Foo<Bar>
}
