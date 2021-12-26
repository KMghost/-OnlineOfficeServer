package Services

import (
	"OnlineOfficeServer/Models"
)

type Position Models.Position

func (this Position) Select() (result []Models.Position) {
	var position Models.Position
	result = position.Select()
	return
}

func (this Position) Insert(name string) error {
	var position Models.Position
	err := position.Insert(name)
	return err
}
