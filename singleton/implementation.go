package singleton

import "fmt"

type db struct{}

var instance *db

func InitDataBase() *db {
	if instance == nil {
		fmt.Println("iniciando nova instancia...")
		instance = &db{}
	} else {
		fmt.Println("pegando instancia ja criada....")
	}

	return instance
}
