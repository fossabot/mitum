package common

const (
	_ uint = iota
	OverflowErrorCode
	UnknownSealTypeCode
	InvalidVoteCode
	JSONUnmarshalCode
	NotImplementedCode
	InvalidHashCode
	EmptyHashCode
	InvalidHashHintCode
)

var (
	OverflowError        Error = NewError("common", OverflowErrorCode, "overflow number")
	UnknownSealTypeError Error = NewError("common", UnknownSealTypeCode, "unknown seal type found")
	InvalidVoteError     Error = NewError("common", InvalidVoteCode, "invalid vote found")
	JSONUnmarshalError   Error = NewError("common", JSONUnmarshalCode, "failed json unmarshal")
	NotImplementedError  Error = NewError("common", NotImplementedCode, "not implemented")
	InvalidHashError     Error = NewError("common", InvalidHashCode, "invalid has found")
	EmptyHashError       Error = NewError("common", EmptyHashCode, "hash is empty")
	InvalidHashHintError Error = NewError("common", InvalidHashHintCode, "invalid hash hint")
)
