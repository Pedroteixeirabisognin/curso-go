package matematica

import "testing"

//Use go test para testar ou go test ./... go test -v para mostrar o print
const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel() //Permite que um teste seja feito em paralelo com outro
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
