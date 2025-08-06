package main

type Event struct {
	name        string
	description string
	link        string
	location    string
	startTime   string
	endTime     string
}

func (e Event) ToArgSlice() []string {
	return []string{e.name, e.description, e.link, e.location, e.startTime, e.endTime}
}
