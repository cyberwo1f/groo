package replacer

type UnknownReplaceError struct {}

func (e *UnknownReplaceError) Error() string {
	return "Unknown Replace pattern."
}

type UnknownProtocolError struct {}

func (e *UnknownProtocolError) Error() string {
	return "Unknown Protocol"
}

type UnknownDomainError struct {}

func (e *UnknownDomainError) Error() string {
	return "Unknown Domain Service"
}
