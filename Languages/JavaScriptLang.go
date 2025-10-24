package Languages

import (model "pc/Model")


func JavaScriptLang() []string{
	lang:=model.Lang{
		MainCodeExample: `console.log("Hello!");`,
		MainFileName: "script.js",
	}
	return []string{
		lang.MainFileName,
		lang.MainCodeExample,
	}
}
