package Languages

import (
	model "pc/Model"
)

func JavaLang() []string{
	lang:=model.Lang{
		MainCodeExample:`public class Main {
	public static void main(String[] args) {
		System.out.println("Hello!");
	}
}`,MainFileName:"Main.java",
	}

	return []string{lang.MainFileName, lang.MainCodeExample}
}