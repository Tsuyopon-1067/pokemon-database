package main

import (
	"strconv"
	"fmt"
)

type Pokemon struct {
	Id   int
	Name string
	H    int
	A    int
	B    int
	C    int
	D    int
	S    int
	Sum  int
}

func pokemonToString(pokemon Pokemon) string {
	res := fmt.Sprintf("%d, %s: [%d, %d, %d, %d, %d, %d]->%d", pokemon.Id, pokemon.Name, pokemon.H, pokemon.A, pokemon.B, pokemon.C, pokemon.D, pokemon.S, pokemon.Sum)
	return res
}

func createPokemon(id int, name string, h int, a int, b int, c int, d int, s int, sum int) Pokemon {
	res := Pokemon{Id: id, Name: name, H: h, A: a, B: b, C: c, D: d, S: s, Sum: sum}
	return res
}

func createPokemonFromString(id string, name string, h string, a string, b string, c string, d string, s string, sum string) Pokemon {
	iId, _ := strconv.Atoi(id)
	iH, _ := strconv.Atoi(h)
	iA, _ := strconv.Atoi(a)
	iB, _ := strconv.Atoi(b)
	iC, _ := strconv.Atoi(c)
	iD, _ := strconv.Atoi(d)
	iS, _ := strconv.Atoi(s)
	iSum, _ := strconv.Atoi(sum)
	res := createPokemon(iId, name, iH, iA, iB, iC, iD, iS, iSum)
	return res
}