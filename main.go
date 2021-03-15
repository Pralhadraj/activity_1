package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type user struct {
	name    string
	phoneno uint64
	address string
}

type users struct {
	us []user
}

type product struct {
	protype string
}

type ord struct {
	ordid       string
	productname string
	quantity    int
}

type ords struct {
	od []ord
}

/*type userver interface {
	adduser(u []user, phoneno uint64) bool
}*/

func (u *users) adduser(pno uint64) bool {
	var flag bool = false
	for _, elem := range u.us {
		fmt.Printf("%v\n", elem)
		if elem.phoneno == pno {
			flag = true
			break
		}
	}

	return flag
}

func getproducts(pro []product) {
	fmt.Println("list of products : ")
	for _, ele := range pro {
		fmt.Printf("%s\n", ele.protype)
	}
	fmt.Printf("\n\n")
}

func (o *ords) palceord(item string, quan int) {
	var ordid string = "oid" + strconv.Itoa(rand.Intn(1000))
	var ornew ord = ord{ordid: ordid, productname: item, quantity: quan}
	o.od = append(o.od, ornew)
}

func (o *ords) changeord(odid string, quan int) {

	for i, elem := range o.od {
		if elem.ordid == odid {
			o.od[i].quantity = quan
			break
		}
	}

}

func (o *ords) getords() {
	for i := 0; i < len(o.od); i++ {
		fmt.Printf("ords :%v\n", o.od[i])
	}

	fmt.Printf("\n\n")
}

func (o *ords) deleteord(doid string) {
	for i := 0; i < len(o.od); i++ {
		if o.od[i].ordid == doid {
			o.od = append(o.od[:i], o.od[i+1:]...)
			break
		}
	}
}

func main() {
	var u users
	var ods ords
	var first user = user{name: "Pralhadraj", phoneno: 9535294676, address: "Hubli"}
	u.us = append(u.us, first)
	//fmt.Printf("%v\n", u[0])
	var uname string
	var phno uint64
	var addr string
	//var uv userver
	var flag bool

	var choice int

	var prot []product = []product{
		{protype: "Shoes"},
		{protype: "Refrigerator"},
		{protype: "Bat"},
		{protype: "Ball"}}

	for {
		fmt.Printf("enter :\n1-add user\n2-To get products list\n3-place ord\n4-changing ord\n5-delete ord \n6-get ords\n")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			{
				fmt.Println("enter user name, phoneno, place")
				fmt.Scanln(&uname)
				fmt.Scanln(&phno)
				fmt.Scanln(&addr)

				first = user{name: uname, phoneno: phno, address: addr}

				flag = u.adduser(phno)

				if !flag {
					u.us = append(u.us, first)
					fmt.Printf("user registered\n")
				} else {
					fmt.Printf("phone no already registered\n")
				}
				fmt.Printf("\n\n users present\n\n")
				for _, elem := range u.us {
					fmt.Printf("%v\n", elem)
				}
			}
		case 2:
			{
				getproducts(prot)
			}

		case 3:
			{

				fmt.Println("place ord")
				var proname string
				var quan int
				fmt.Println("enter product name and quantity for ord")
				fmt.Scanln(&proname)
				fmt.Scanln(&quan)

				ods.palceord(proname, quan)

				fmt.Println("ords paced are:")
				ods.getords()

			}

		case 4:
			{
				var ordid string
				var quan int

				fmt.Println("enter  ordid for changing the ord")
				fmt.Scanln(&ordid)

				fmt.Println("enter  quantity to change for ord")
				fmt.Scanln(&quan)

				ods.changeord(ordid, quan)

				fmt.Println("ords after changing are:")
				ods.getords()

			}
		case 5:
			{
				var delid string
				fmt.Printf("enter the ordid to delete ord:\n")
				fmt.Scanln(&delid)

				ods.deleteord(delid)

				fmt.Println("ords after deleteing are:")
				ods.getords()

			}
		case 6:
			{
				ods.getords()
			}

		default:
			os.Exit(1)

		}
	}

}
