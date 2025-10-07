package main

import (
	enum "pc/Enum"
	kernel "pc/Kernel"
)

func main() {
	kernel.NewFileCommand("game", enum.Java, 0.1, "Nicolas")
}
