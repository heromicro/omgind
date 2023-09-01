package enumtipe

import "database/sql/driver"

type DictValueTipe string

const (
	DictValueTipeInt DictValueTipe = "int"
	DictValueTipeStr DictValueTipe = "str"
)

var dictValueTipes = []string{
	DictValueTipeInt.String(),
	DictValueTipeStr.String(),
}

func (DictValueTipe) Values() (dtipes []string) {
	dtipes = append(dtipes, dictValueTipes...)
	return
}

func (dvt DictValueTipe) Value() (driver.Value, error) {
	return dvt.String(), nil
}

func (dvt DictValueTipe) String() string {
	return string(dvt)
}
