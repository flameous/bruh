package problems

// actually this problem from Yandex final interview

// problem: there is a string
// that contains smiles - substrings started with colon, hyphen and at least one round bracket. example : ":-)", ":-((((("
// string example: "hello :-)))))) world :-)))(foobar)"
// modify a string and remove all the smiles
// result: "hello  world (foobar)"
// here I use byte slice instead of string to solve a problem with O(1) extra space
// also I assume that string contains only ASCII characters to simplify the iteration over the slice
func removeSmiles(origin []byte) []byte {
	if len(origin) < 3 {
		return origin
	}

	ret := origin[:0] // ret uses the same underlaying array as origin

	i := 0
	for i < len(origin) {
		if origin[i] == ':' && len(origin)-i > 2 {
			if origin[i+1] == '-' && (origin[i+2] == '(' || origin[i+2] == ')') {
				i += 2
				bracket := origin[i]
				for origin[i] == bracket {
					i++
					if i == len(origin) {
						return ret
					}
				}
				continue
			}
		}

		ret = append(ret, origin[i])
		i++
	}
	return ret
}

//TODO: advanced version
// "hello :-:-(()))world" -> "hello world"
func removeNestedSmiles(origin []byte) []byte {
	return nil
}
