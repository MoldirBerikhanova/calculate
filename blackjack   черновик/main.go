package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type BlackJack interface {
// }

// type Player struct {
// 	masti      string
// 	CardsValue string
// }

// type Diler struct {
// 	DilersValue string
// 	DilersMasti string
// }

func main() {

	Mast := []string{"Червы", "Пики", "Буби", "Кресть"}                           //////////////////////////////cоздание колоды
	Cards := []string{"6", "7", "8", "9", "10", "Валет", "Дама", "Король", "Туз"} /////создание колоды
	//Player := make(map[string]string)
	unique := make(map[string]bool)

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	/////////////Cоздание Дилера///////////////////////////////////////////////////////////////////////
	var DilersValue []string //////////////значения колоды
	for i := 0; i <= 4; i++ {
		randomIndex := rand.Intn(len(Cards))
		DilersValue = append(DilersValue, Cards[randomIndex])
		rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

	}

	for _, v := range DilersValue {
		if _, exists := unique[v]; !exists {
			unique[v] = true
			//		fmt.Println(v)
		}

	}

	//fmt.Println(DilersValue)

	var DilersMasti []string ///////масти колоды
	for i := 0; i <= 4; i++ {
		randomIndex := rand.Intn(len(Mast))
		DilersMasti = append(DilersMasti, Mast[randomIndex])

	}

	//fmt.Println(DilersMasti)

	fmt.Println("Diler")
	for i := 0; i <= 4; i++ {

		fmt.Println(DilersMasti[i]+" "+DilersValue[i]+";", " ")

	}

	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------")
	////////СОЗДАНИЕ ИГРОКА/////////////
	//перемешивание карт колоды по значениям
	var CardsValue []string //знаечния колоды
	for i := 0; i <= 4; i++ {
		randomIndex := rand.Intn(len(Cards))
		CardsValue = append(CardsValue, Cards[randomIndex])
		rand.Shuffle(len(Cards), func(i, j int) { Cards[i], Cards[j] = Cards[j], Cards[i] })

	}

	for _, v := range CardsValue {
		if _, exists := unique[v]; !exists {
			unique[v] = true
			//fmt.Println(v)
		}
	}

	// for _ , v := range CardsValue {
	// //	if _, exists := DilersValue[v];!exists {

	// 	}
	// }
	// 	masti = append(masti, )
	///////////////////////////////////////////////////////////////////////////////////////////////////////

	//перемешивание карт колоды по мастям
	var masti []string //масти колоды
	for i := 0; i <= 4; i++ {
		randomIndex := rand.Intn(len(Mast))
		masti = append(masti, Mast[randomIndex])

	}

	//fmt.Println(masti)

	//вывести игрока на консоль  PLAYER-1
	fmt.Println("Player")
	for i := 0; i <= 4; i++ {

		fmt.Println(masti[i]+" "+CardsValue[i]+";", " ")

	}

	fmt.Println() //Player[CardsValue[i]] = masti[i]

}

// 	Player[valuue[k]] = masti[i];
//fmt.Println("Игрок1:", Player)
// 	fmt.Println(Player)

// numb := []int{1, 5, 55} //  эту надо перевести на Латиницу

// 	var result []string
// 	for _, v := range numb {
// 		if x, found := R[v]; found {
// 			result = append(result, x)
// 		} else {
// 			result = append(result, "")
// 		}
// 	}
// 	fmt.Println(result)

// var Player [] string

// }

// fmt.Println(Player, len(Player))

// fmt.Println(Player)

// // source := rand.NewSource(time.Now().UnixNano())
// rand := rand.New(source)
// //var result [] int
// for i := 0; i >= 3; i++ {

// 	utr := Cards[index]
// 	M[utr]++
// }
// fmt.Println(M)

/*
	Реализовать поэтапно:
	- Создание колоды;
	- Перемешивание карт колоды;
	- Создать как минимум одного игрока и дилера;
	- Раздать каждому игроку и дилеру по 5 карт;
	- Вывести в консоль карты игроков и дилера в формате `Карта {масть} - {номинал}`
*/

// ВАШ КОД ТУТ...

// func printCards(cards []string) {
// 	for i, val := range cards {
// 		fmt.Printf("Карта #%d - %s\n", i, val)
// 	}
// }
