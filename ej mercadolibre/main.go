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
err=nil

bytes:=ioutil.ReadAll(response.bytes)

var cats Categories
json.Unmarshal(bytes, &cats)

return cats, err
}
func main(){
cats,err :=GetCategories("MLA")

if err==nil{
fmt.Println("ERROR")
}
fmt.Println("dfjsbdf")

}




//pkg.go.dev/net/http