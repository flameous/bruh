package problems

import (
	"testing"
)

func TestRemoveSmiles(t *testing.T) {

	t.Run("smile in the middle of the string", func(t *testing.T) {
		b := []byte("hello :-)))))world")
		expected := "hello world"

		got := removeSmiles(b)

		if &got[0] != &got[0] {
			t.Fatalf("there are different slices")
		}

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("only smiles", func(t *testing.T) {
		b := []byte(":-)))")
		expected := ""

		got := removeSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("only smiles #2", func(t *testing.T) {
		b := []byte(":-):-((((:-))))")
		expected := ""

		got := removeSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("smile at the end", func(t *testing.T) {
		b := []byte("hello:-)))")
		expected := "hello"

		got := removeSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("smiles again", func(t *testing.T) {
		b := []byte(":-)))))(foo):-((((bar():-)))(")
		expected := "(foo)bar()("

		got := removeSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("not smiles", func(t *testing.T) {
		b := []byte(":-| :-aaa ::::a::::----) :--((()")
		expected := ":-| :-aaa ::::a::::----) :--((()"

		got := removeSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})
}

func TestRemoveNestedSmiles(t *testing.T) {

	t.Run("nested smiles", func(t *testing.T) {
		b := []byte("hello :-:-)))(((world")
		expected := "hello world"

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("nested smiles #2", func(t *testing.T) {
		b := []byte("hello :-:-)))world")
		expected := "hello :-world"

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("nested smiles #3", func(t *testing.T) {
		b := []byte("hello :-:-)))((()world")
		expected := "hello )world"

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("nested smiles #4", func(t *testing.T) {
		b := []byte("hello :-:-)))(((world:--")
		expected := "hello world:--"

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("nested smiles #5", func(t *testing.T) {
		b := []byte("::-)-(")
		expected := ""

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

	t.Run("nested smiles #6", func(t *testing.T) {
		b := []byte("::-)-|")
		expected := ":-|"

		got := removeNestedSmiles(b)

		if string(got) != expected {
			t.Fatalf("expected '%s' but got '%s'", expected, string(got))
		}
	})

}
