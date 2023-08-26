# pokemon-database
# 概要
ポケモンのデータをデータベースに登録したり削除したり表示したりする。 <br>
GoとMariaDB（MySQL）の接続練習。<br>

# 使い方
## ビルド
シェルスクリプトを動かせばビルドできる。
```
./build.sh
```
もしくは普通にビルド。
```
go build pokemon_database.go connectSql.go pokemonStruct.go scraping.go
```
## データベース
起動前にSQLの準備が必要。詳細はdoneList.mdを参照。

## オプション
起動時は以下のようにしてオプションを指定。\<option\>を他の文字に置き換えて指定する。
```
./pokemon_database <option>
```
- なし 通常モード
- h ヘルプの表示
- r データベースのリセット（初回起動に行う）

## 通常モード
commandは以下の通り。
- insert：データの挿入
- delete：データの削除
- print：データベースの全レコードを標準出力
- quit：終了




# 関数とかの仕様
## build pokemon_database.go
コマンドとかを受け付ける
### func main()
コマンドライン引数かそれ以外で分岐
### func normal()
通常モードの処理（無限ループ）
### func help()
ヘルプを表示
### func other()
コマンドライン引数が想定外だったときに呼び出される簡易ヘルプ
### func commandInsert()
標準入力で受け取ったデータから構造体を生成してデータベースに登録する
### func commandInsert()
標準入力で受け取った値にマッチするものをデータベースから削除する
### func isNumber()
引数が自然数ならtrueを返す（つまり負数や少数は非対応）
### func commandPrint()
データベースの全レコードを出力する

## connectSql.go
SQLとのやりとり
### func printAll()
テーブル内のデータを標準出力
### func insert(pokemon Pokemon)
テーブルにデータ追加
### func deleteId(id string)
データベースから指定Idのレコードを削除
### func deleteId(name string)
データベースから指定Nameのレコードを削除
### func resetLoad()
テーブルからレコードを全削除してからスクレイピングしたデータを登録する

## scraping.go
スクレイピング関連
### func scraping() []*Pokemon
wikiからスクレイピングしたデータをスライスに格納して返す
### func printList(pokemonList []*Pokemon)
スライスの各要素を標準出力

## pokemonStruct.go
構造体関連
### type Pokemon struct
ポケモンデータの構造体 <ID, 名前, HP, 攻撃, 防御, 特攻, 特防, 素早さ, 合計>
### func pokemonToString(pokemon Pokemon) string
pokemon構造体を標準出力できる形のstringに変換する
### func createPokemon(id int, name string, h int, a int, b int, c int, d int, s int, sum int) Pokemon
各パラメータからpokemon構造体を生成する
### func createPokemonFromString(id string, name string, h string, a string, b string, c string, d string, s string, sum string) Pokemon
createPokemonのパラメータを全部stringとした関数