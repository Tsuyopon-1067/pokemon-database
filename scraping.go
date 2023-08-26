package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

func scraping() []*Pokemon {
	var pokemonList []*Pokemon

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		// Shift_JISからUTF-8に変換するTransformerを作成
		decoder := japanese.ShiftJIS.NewDecoder()

		// Transformerを使用して変換を行う
		utf8Body, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(r.Body), decoder))
		if err != nil {
			log.Fatalf("Failed to decode: %v", err)
		}
		r.Body = utf8Body
	})

	c.OnHTML("[id='tablesorter-zukan'] tr", func(e *colly.HTMLElement) {
		// この関数はテーブ
		cells := e.ChildTexts("td")
		if len(cells) >= 11 {
			var id, h, a, b, c, d, s, sum int
			id, _ = strconv.Atoi(cells[0])
			name := cells[1]
			h, _ = strconv.Atoi(cells[4])
			a, _ = strconv.Atoi(cells[5])
			b, _ = strconv.Atoi(cells[6])
			c, _ = strconv.Atoi(cells[7])
			d, _ = strconv.Atoi(cells[8])
			s, _ = strconv.Atoi(cells[9])
			sum, _ = strconv.Atoi(cells[10])

			tmp := Pokemon{Id: id, Name: name, H: h, A: a, B: b, C: c, D: d, S: s, Sum: sum}
			pokemonList = append(pokemonList, &tmp)
		}
	})

	url := "https://game-e.com/pokemon-hgss/zukan_johto/"
	c.Visit(url)

	return pokemonList
}

func printList(pokemonList []*Pokemon) {
	for _, v := range pokemonList {
		fmt.Printf("%d, %s: [%d, %d, %d, %d, %d, %d]->%d\n", v.Id, v.Name, v.H, v.A, v.B, v.C, v.D, v.S, v.Sum)
	}
}