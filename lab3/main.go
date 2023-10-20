package main

import (
	"fmt"
	"github.com/ichiban/prolog"
	"os"
	"regexp"
	"bufio"
)

// Для проверки вхождения элемента в массив
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// Ввод данных для проголога по шаблону
func input() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	re := regexp.MustCompile(`Мне нужен игрок из (\w+) и дешевле (\d+) рублей`)
	matches := re.FindStringSubmatch(text)
	if len(matches) >= 3 {
		country := matches[1]
		price := matches[2]
		players := getPlayersByPrice(country, price)
		for _, player := range players {
			fmt.Println(player)
		}
	} else {
		re = regexp.MustCompile(`Мне нужен самый дорогой игрок`)
		matches = re.FindStringSubmatch(text)
		if len(matches) >= 1 {
			//country := matches[1]
			fmt.Println(getMostExpensivePlayer())
		} else {
			panic("Не удалось извлечь данные из запроса")
		}
	}
}

// Получение самого дорогого игрока стране
func getMostExpensivePlayer() (string) {
	p := prolog.New(os.Stdin, os.Stdout)
	if err := p.Exec(`
		:- [lab1].  % Подключение файла lab1.pl
	`); err != nil {
		panic(err)
	}
	// Отправка запроса на пролог
	sols, err := p.Query(`find_max_player(Player, MaxPrice).`)
	if err != nil {
		panic(err)
	}
	defer sols.Close()
	player := ""
	// Циклом забираются все результаты
	for sols.Next() {
		var s struct {
			Player string
		}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		player = s.Player
	}
	return player
}

// Получение игроков из страны и с максимальной ценой
func getPlayersByPrice(country string, price string) ([]string) {
	p := prolog.New(os.Stdin, os.Stdout)
	if err := p.Exec(`
		:- [lab1].  % Подключение файла lab1.pl
	`); err != nil {
		panic(err)
	}
	// Отправка запроса на пролог
	sols, err := p.Query(`less_than_price(Player, ` + price + `), 
						  player_country(Player, ` + country + `).`)
	if err != nil {
		panic(err)
	}
	defer sols.Close()
	var players []string
	// Циклом забираются все результаты
	for sols.Next() {
		var s struct {
			Player string
		}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		players = append(players, s.Player)
	}
	return players
}

func main() {
	fmt.Println("Здравствуйте, вы можете отправить запрос для подбора игрока!")
	fmt.Println("Вы можете получить самого дорого игрока или ввести свою сумму")
	input()
}
// Мне нужен игрок из russia и дешевле 1000 рублей
// Мне нужен самый дорогой игрок
