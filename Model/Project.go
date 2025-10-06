package model

import (
	"pc/Enum"
	"time"
)
type Project struct{
	NameProject string 
	LangProject enum.Lang
	DateCreation string
	Version float64
	Author string
}
func (p *Project) SetDateCreationForNow() {
	p.DateCreation = time.Now().Format("02/01/2006 15:04:05")
}
func (p *Project) SetVersion(newVersion float64){
	p.Version = newVersion
}
func (p *Project) GetInformations() *Project{
	return p
}