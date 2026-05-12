package main

import (
	"adventofcode/ponteiro"
	"testing"
)

func TestRotacaoHorario(t *testing.T) {
	ponteiro.Reset()
	ponteiro.OnGiro(func(posicaoAtual int) {
		if posicaoAtual != 65 {
			t.Errorf("Esperado 65, obteve %d", posicaoAtual)
		}
	})
	ponteiro.GirarHorario(15)

	ponteiro.OnGiro(func(posicaoAtual int) {
		if posicaoAtual != 75 {
			t.Errorf("Esperado 75, obteve %d", posicaoAtual)
		}
	})
	ponteiro.GirarHorario(10)

}

func TestRotacaoAntiHorario(t *testing.T) {
	ponteiro.Reset()
	ponteiro.OnGiro(func(posicaoAtual int) {
		if posicaoAtual != 20 {
			t.Errorf("Esperado 20, obteve %d", posicaoAtual)
		}
	})
	ponteiro.GirarAntiHorario(30)
}

func TestRotacaoAntiHorario2Posicoes(t *testing.T) {
	ponteiro.Reset()
	ponteiro.GirarAntiHorario(30)

	ponteiro.OnGiro(func(posicaoAtual int) {
		if posicaoAtual != 95 {
			t.Errorf("Esperado %d, obteve %d", 95, posicaoAtual)
			t.Context().Done()
		}
	})
	ponteiro.GirarAntiHorario(25)
}

func TestRotacaoHorario2Posicoes(t *testing.T) {
	ponteiro.Reset()
	ponteiro.GirarAntiHorario(55)
	ponteiro.OnGiro(func(posicaoAtual int) {
		if posicaoAtual != 5 {
			t.Errorf("Esperado %d, obteve %d", 5, posicaoAtual)
		}
	})
	ponteiro.GirarHorario(10)
}

func TestRotacaoAntihorario(t *testing.T) {
	ponteiro.SetPonteiro(0)
	ponteiro.GirarAntiHorario(1)

	if ponteiro.GetPonteiro() != 99 {
		t.Errorf("Esperado %d, obteve %d", 99, ponteiro.GetPonteiro())
	}
}

func TestGiroHorario2(t *testing.T) {
	ponteiro.SetPonteiro(99)
	ponteiro.GirarHorario(1)
	if ponteiro.GetPonteiro() != 0 {
		t.Errorf("Esperado %d, obteve %d", 0, ponteiro.GetPonteiro())
	}
}

func TestGiroHorario3(t *testing.T) {
	ponteiro.SetPonteiro(99)
	ponteiro.GirarHorario(100)
	if ponteiro.GetPonteiro() != 99 {
		t.Errorf("Esperado %d, obteve %d", 99, ponteiro.GetPonteiro())
	}
}

func TestGirosHorarios(t *testing.T) {
	ponteiro.SetPonteiro(99)
	for i := 0; i < 100; i++ {
		ponteiro.GirarHorario(1)
		if ponteiro.GetPonteiro() != i {
			t.Errorf("Esperado %d, obteve %d", i, ponteiro.GetPonteiro())
		}
	}
}

func TestGirosAntiHorarios(t *testing.T) {
	ponteiro.SetPonteiro(0)
	for i := 0; i < 100; i++ {
		ponteiro.GirarAntiHorario(1)
		if ponteiro.GetPonteiro() != 99-i {
			t.Errorf("Esperado %d, obteve %d", 99-i, ponteiro.GetPonteiro())
		}
	}
}

func TestGirosAltos(t *testing.T) {
	ponteiro.SetPonteiro(0)
	for i := 0; i < 783; i++ {
		ponteiro.GirarAntiHorario(1)
		t.Log(ponteiro.GetPonteiro())
	}

	if ponteiro.GetPonteiro() != 17 {
		t.Errorf("Esperado %d, obteve %d", 17, ponteiro.GetPonteiro())
	}
}

func TestGirosAltosHorarios(t *testing.T) {
	ponteiro.SetPonteiro(0)
	for i := 0; i < 783; i++ {
		ponteiro.GirarHorario(1)
		t.Log(ponteiro.GetPonteiro())
	}

	if ponteiro.GetPonteiro() != 83 {
		t.Errorf("Esperado %d, obteve %d", 83, ponteiro.GetPonteiro())
	}
}
