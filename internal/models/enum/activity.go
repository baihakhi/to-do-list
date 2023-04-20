package enum

const (
	// constanta for priority of activities
	HIGHEST string = "highest"
	HIGH    string = "high"
	NORMAL  string = "normal"
	LOW     string = "low"

	// constanta for state of activities
	TODO      string = "todo"
	INPROGRES string = "in-progress"
	DONE      string = "done"
)

var IsValidPriorty = map[string]bool{
	HIGHEST: true,
	HIGH:    true,
	NORMAL:  true,
	LOW:     true,
}

var IsValidState = map[string]bool{
	TODO:      true,
	INPROGRES: true,
	DONE:      true,
}
