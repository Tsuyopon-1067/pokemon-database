package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		normal()
		return
	}

	switch os.Args[1] {
		case "h":
			help()
			break
		case "r":
			resetLoad()
			break
		default:
			other()
	}
}

func normal() {
	for {
		fmt.Print("command: ")
		var command string
		fmt.Scanf("%s", &command)

		isQuit := false
		switch command {
			case "quit":
				isQuit = true
				break
			case "insert":
				commandInsert()
				break
			case "delete":
				commandDelete()
				break
			case "print":
				commandPrint()
				break
			default:
				fmt.Println("command（通常モード）")
				fmt.Println("insert：データの挿入")
				fmt.Println("delete：データの削除")
				fmt.Println("print：データベースの全レコードを標準出力")
				fmt.Println("quit：終了")
		}

		if isQuit {
			break
		}
	}
}

func help() {
	fmt.Println("Help")
	fmt.Println("使用方法：pokemon_database <option>")
	fmt.Println("option")
	fmt.Println("h：ヘルプ")
	fmt.Println("r：データベースのリセット")
	fmt.Println("なし：通常モード")
	fmt.Println("")
	fmt.Println("command（通常モード）")
	fmt.Println("insert：データの挿入")
	fmt.Println("delete：データの削除")
	fmt.Println("print：データベースの全レコードを標準出力")
	fmt.Println("quit：終了")
}

func other() {
	fmt.Println("使用方法：pokemon_database <option>")
	fmt.Println("使用可能なオプションのリストについては、hを使用します")
}

func commandInsert() {
	fmt.Println("enter data")
	fmt.Println("ID, Name, HP, Attack, Defense, Special Attack, Special Defense, Speed, Sum")
	var id, name, h, a, b, c, d, s, sum string
	fmt.Scanf("%s %s %s %s %s %s %s %s %s", &id, &name, &h, &a, &b, &c, &d, &s, &sum)
	pokemon := createPokemonFromString(id, name, h, a, b, c, d, s, sum)
	insert(pokemon)
}

func commandDelete() {
	fmt.Println("enter ID or Name")
	var s string
	fmt.Scanf("%s", &s)

	if isNumber(s) {
		deleteId(s)
	} else {
		deleteName(s)
	}
}

func isNumber(s string) bool {
    res := true
    for _, v := range s {
        if v < '0' || '9' < v {
            res = false
            break
        }
    }
    return res
}

func commandPrint() {
	printAll()
}