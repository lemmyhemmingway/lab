package helloworld

import "testing"

func TestHello(t *testing.T)  {
	t.Run("say hello to given name in french", func(t *testing.T) {
		got := Hello("mahmut", "french")
		want := "Bonjour, mahmut"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say hello to given name in spanish", func(t *testing.T) {
		got := Hello("mahmut", "spanish")
		want := "Hola, mahmut"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say hello to given name", func(t *testing.T) {
		got := Hello("gundi", "")
		want := "Hello, gundi"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello world if empty string given", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})


	
}
