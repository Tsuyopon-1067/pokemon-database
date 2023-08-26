# pokemon-database

## connectSql.go
## scraping.go
### func scraping() []*Pokemon
wikiからスクレイピングしたデータをスライスに格納して返す
### func printList(pokemonList []*Pokemon)
スライスの各要素を標準出力

## pokemonStruct.go
### type Pokemon struct
ポケモンデータの構造体 <ID、名前、HP、攻撃、防御、特攻、特防、素早さ、合計>
### func pokemonToString(pokemon Pokemon) string
pokemon構造体を標準出力できる形のstringに変換する
### func createPokemon(id int, name string, h int, a int, b int, c int, d int, s int, sum int) Pokemon
各パラメータからpokemon構造体を生成する
### func createPokemonFromString(id string, name string, h string, a string, b string, c string, d string, s string, sum string) Pokemon
createPokemonのパラメータを全部stringとした関数