package main

//import "fmt"
//import "strconv"
import (
	"fmt"
	"strconv"
)
func main() {
	edad := 23;
	edad_str := strconv.Itoa(edad);//se debe convertir la variable para poder hacerlo
	fmt.Println("mi edad e: "+edad_str+"\n");
	fmt.Println("mi nueva edad es: "+strconv.Itoa(edad)+"\n")


	nombre := "2"
    nomb_str,_ := strconv.Atoi(nombre);//el retorna otro valor y el compilador espera que lo recibamos _ significa la obtengo pero no me importa
	fmt.Println(nomb_str+10);
}