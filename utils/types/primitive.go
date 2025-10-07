package types

type (
	NullableBool   = Nullable[bool]
	NullableString = Nullable[string]
	NullableByte   = Nullable[byte]
	NullableRune   = Nullable[rune]
	NullableAny    = Nullable[any]

	NullableInt   = Nullable[int]
	NullableInt8  = Nullable[int8]
	NullableInt16 = Nullable[int16]
	NullableInt32 = Nullable[int32]
	NullableInt64 = Nullable[int64]

	NullableUint   = Nullable[uint]
	NullableUint8  = Nullable[uint8]
	NullableUint16 = Nullable[uint16]
	NullableUint32 = Nullable[uint32]
	NullableUint64 = Nullable[uint64]

	NullableFloat32 = Nullable[float32]
	NullableFloat64 = Nullable[float64]

	NullableComplex64  = Nullable[complex64]
	NullableComplex128 = Nullable[complex128]
)
