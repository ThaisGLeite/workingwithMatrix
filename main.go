package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
 orders[i]=[customerNamei,tableNumberi,foodItemi]
	Input: orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
	Output: [["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
	Explanation:
	The displaying table looks like:
	Table,Beef Burrito,Ceviche,Fried Chicken,Water
	3    ,0           ,2      ,1            ,0
	5    ,0           ,1      ,0            ,1
	10   ,1           ,0      ,0            ,0
	For the table 3: David orders "Ceviche" and "Fried Chicken", and Rous orders "Ceviche".
	For the table 5: Carla orders "Water" and "Ceviche".
	For the table 10: Corina orders "Beef Burrito".
*/

func main() {

	retorno := [][]string{{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
		{"Carla", "7", "Water"}}

	tables := displayTable(retorno)
	for i := range tables {
		for _, j := range tables[i] {
			fmt.Printf("%s, ", j)
		}
		fmt.Println("")
	}
}

func displayTable(orders [][]string) [][]string {
	costumerOrders := make([][]string, 0)
	lista := make([]string, 0)

	for _, v := range orders {
		lista = highlanderFood(lista, v[2])
		costumerOrders = highlanderMesa(costumerOrders, v[1])
	}
	sort.Strings(lista)
	table := []string{"Table"}
	lista = append(table, lista...)

	costumerOrders = append([][]string{lista}, costumerOrders...)
	sort.Slice(costumerOrders, func(i, j int) bool {
		um, _ := strconv.Atoi(costumerOrders[i][0])
		dois, _ := strconv.Atoi(costumerOrders[j][0])
		return um < dois
	})

	for i := 1; i < len(costumerOrders); i++ {
		for j := 1; j < len(costumerOrders[0]); j++ {
			costumerOrders[i] = append(costumerOrders[i], "0")
		}
	}

	costumerOrders = foodOrder(costumerOrders, lista, orders)
	return costumerOrders
}

func highlanderFood(lista []string, a string) []string {
	for _, v := range lista {
		if v == a {
			return lista
		}
	}
	lista = append(lista, a)
	return lista
}

func highlanderMesa(orders [][]string, mesaNr string) [][]string {
	for _, v := range orders {
		if v[0] == mesaNr {
			return orders
		}

	}
	mesa := make([]string, 0)
	mesa = append(mesa, mesaNr)
	orders = append(orders, [][]string{mesa}...)
	return orders
}

func foodOrder(costumerOrders [][]string, lista []string, orders [][]string) [][]string {

	for _, x := range orders {
		mesa := x[1]
		for linhadaMesa, v := range costumerOrders {
			if v[0] == mesa {
				for i, y := range lista {
					if y == x[2] {
						nFood, _ := strconv.Atoi(costumerOrders[linhadaMesa][i])
						nFood++
						costumerOrders[linhadaMesa][i] = strconv.Itoa(nFood)
					}
				}
			}

		}
	}

	return costumerOrders

}
