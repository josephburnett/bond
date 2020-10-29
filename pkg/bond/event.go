package bond

type Event string

const (
	CORRECT   Event = "correct"
	INCORRECT Event = "incorrect"
	HINT      Event = "hint"
)
