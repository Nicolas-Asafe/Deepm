package kernel

import model "pc/Model"
import "path/filepath"

type FuncKernel struct {
	BasePath       string
	Project        model.Project
	GetProjectPath func() string
}

func NewFuncKernel(project *model.Project) *FuncKernel {
	basePath := `C:\Dev\Projects`
	lang := project.LangProject.String() + "Projects"
	name := project.NameProject
	return &FuncKernel{
		BasePath: basePath,
		Project: *project,
		GetProjectPath: func() string {
			projectPath := filepath.Join(basePath, lang, name)
			return projectPath
		},
	}
}
