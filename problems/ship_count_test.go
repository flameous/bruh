package problems

import "testing"

func TestShipCount(t *testing.T) {

	t.Run("ships", func(t *testing.T) {
		expected := 3

		field := [][]int{
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 1},
		}

		count := shipCount(field)
		if expected != count {
			t.Fatalf("expected %d but got %d", expected, count)
		}
	})

	t.Run("ships at the corners", func(t *testing.T) {
		expected := 4

		field := [][]int{
			{1, 0, 0, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 0, 0, 1},
		}

		count := shipCount(field)
		if expected != count {
			t.Fatalf("expected %d but got %d", expected, count)
		}
	})

	t.Run("long ship", func(t *testing.T) {
		expected := 1

		field := [][]int{
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}

		count := shipCount(field)
		if expected != count {
			t.Fatalf("expected %d but got %d", expected, count)
		}
	})

	t.Run("long vertical ship", func(t *testing.T) {
		expected := 1

		field := [][]int{
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		}

		count := shipCount(field)
		if expected != count {
			t.Fatalf("expected %d but got %d", expected, count)
		}
	})

	t.Run("a lot of ships", func(t *testing.T) {
		expected := 7

		field := [][]int{
			{0, 0, 0, 0, 0, 1, 1, 0},
			{0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 1, 0, 1},
			{0, 1, 1, 1, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}

		count := shipCount(field)
		if expected != count {
			t.Fatalf("expected %d but got %d", expected, count)
		}
	})

}
