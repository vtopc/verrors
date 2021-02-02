package verrors

// Err general error
type Err struct {
	Msg string
	ID  string // optional
}

func (e Err) Error() string {
	ret := e.Msg

	if e.ID != "" {
		ret = "ID(" + e.ID + "): " + ret
	}

	return ret
}

type InvalidArgument struct {
	Err
}

type NotFound struct {
	Err
}

func (e NotFound) Error() string {
	if e.Msg == "" {
		e.Msg = "not found"
	}

	return e.Err.Error()
}

type AlreadyExists struct {
	Err
}

func (e AlreadyExists) Error() string {
	if e.Msg == "" {
		e.Msg = "already exists"
	}

	return e.Err.Error()
}
