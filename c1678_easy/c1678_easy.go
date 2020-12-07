package c1678_easy

func interpret(command string) string {
	var opt string
	commandRune := []rune(command)
	i := 0

	for {
		if i >= len(commandRune) {
			break
		}
		if commandRune[i] == 'G' {
			opt += "G"
			i++
		} else if commandRune[i] == '(' && commandRune[i+1] == ')' {
			opt += "o"
			i += 2
		} else {
			opt += "al"
			i += 4
		}
	}

	return opt
}

