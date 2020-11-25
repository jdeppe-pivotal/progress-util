package core

type Progress interface {
	ProcessLine(string)
	Results()
}
