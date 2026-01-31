package domain

type ChecklistItem struct {
	Text string
	Done bool
	Due  string
}

type ProjectItem struct {
	Name    string
	Subject string
	Due     string
	Status  string
}

type ExamItem struct {
	Name     string
	Date     string
	Retakes  []string
	Priority string
}

type SubjectItem struct {
	Code  string
	Name  string
	Exams []ExamItem
}

type LofiTrack struct {
	Title string
	Note  string
}

const LofiVisibleCap = 8

const (
	PriorityHigh = "HIGH"
	PriorityMed  = "MED"
	PriorityLow  = "LOW"
)

const (
	ProjectStatusNotStarted = "NOT STARTED"
	ProjectStatusInProgress = "IN PROGRESS"
	ProjectStatusDone       = "DONE"
)
