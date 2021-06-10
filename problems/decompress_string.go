package problems

// from Tinkoff interview

// Дана строка из латинских заглавных букв, цифр и скобок. Скобочные выражения всегда верные.
// За цифрой всегда стоит или буква или скобочное выражение. Скобочные выражения могут быть вложенные.
// Надо преобразовать в строку умножая выражение за цифрой

// 3AB2(Z3K)CD4B → AAABZKKKZKKKCDBBBB
func decompress(s string) string {
	res, _ := foo([]byte(s), 0, len(s))
	return string(res)
}

func isChar(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func digit(b byte) (int, bool) {
	if '0' <= b && b <= '9' {
		return int(b - '0'), true
	}
	return 0, false
}

func foo(data []byte, start, end int) ([]byte, int) {
	b := []byte{}
	var shift int
	coeff := 1

	for i := start; i < end; i++ {
		if d, ok := digit(data[i]); ok {
			coeff = d
			continue
		}
		if data[i] == ')' {
			shift = i - start + 1
			break
		}

		var seq []byte
		if data[i] == '(' {
			var offset int
			seq, offset = foo(data, i+1, end)
			i += offset
		}
		if isChar(data[i]) {
			seq = []byte{data[i]}
		}

		for j := 0; j < coeff; j++ {
			b = append(b, seq...)
		}
		coeff = 1
	}

	return b, shift
}
