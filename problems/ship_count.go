package problems

// from Tinkoff interview

/*
Есть матрица NxN, состоящая из 0 и 1, и отражающая расположения кораблей на поле для морского боя.

1. Поле может быть любого размера, но обязательно квадратное.
2. Кораблей может быть любое количество
3. Размер кораблей — от 1х1 до 1хN
4. Корабли никак не соприкасаются друг с другом.
5. Могут располагаться горизонтально и вертикально

Необходимо подсчитать количество кораблей.
*/
func shipCount(matrix [][]int) int {
	count := 0

	for col := 0; col < len(matrix[0]); col++ {
		for row := 0; row < len(matrix); row++ {
			if matrix[row][col] == 0 {
				continue
			}

			var a, b bool
			if row == 0 || matrix[row-1][col] == 0 {
				a = true
			}
			if col == 0 || matrix[row][col-1] == 0 {
				b = true
			}

			if a && b {
				count++
			}
		}
	}
	return count
}
