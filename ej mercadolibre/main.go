package main

import (
	"encoding/json"
	"errors"
	"fmt"
	
	"io/ioutil"
	"net/http"
)
type Categories []Category

type Category struct{
ID string `json:"ID"`
Name string `json:"name"`
}


func GetCategories(siteID string) (Categories, error){
response,err  := http.Get("https://api.mercadolibre.com/sites/MLA/Categories")


bytes,err:=ioutil.ReadAll(response.Body)

var cats Categories
err= json.Unmarshal(bytes, &cats)

return cats, err
}
func main(){
	var categorias Categories
	var i int 
	for true{
cats,err :=GetCategories("MLA")
if err=nil{
	fmt.Println("ERROR")
  break 
}
categorias[i]=cats
i++

}
fmt.Println(categorias[1])




}




//pkg.go.dev/net/http