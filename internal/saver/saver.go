package saver

type Saver interface {
	Save(fileName string) error
}
