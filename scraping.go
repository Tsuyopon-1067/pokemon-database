package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/gocolly/colly/v2"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

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
			tmp := createPokemonFromString(cells[0], cells[1], cells[4], cells[5], cells[6], cells[7], cells[8], cells[9], cells[10])
			pokemonList = append(pokemonList, &tmp)
		}
	})

	url := "https://game-e.com/pokemon-hgss/zukan_johto/"
	c.Visit(url)

	return pokemonList
}

func printList(pokemonList []*Pokemon) {
	for _, v := range pokemonList {
		fmt.Println(pokemonToString(*v))
	}
}