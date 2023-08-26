# pokemon-database

## connectSql.go
### type Pokemon struct
ポケモンデータの構造体 <ID、名前、HP、攻撃、防御、特攻、特防、素早さ、合計>

### func scraping() []*Pokemon
wikiからスクレイピングしたデータをスライスに格納して返す
### func printList(pokemonList []*Pokemon)
スライスの各要素を標準出力

## scraping.go