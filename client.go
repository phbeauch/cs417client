package main

import (
	"bytes"
	"fmt"
	"flag"
	"os"
  "net/http"
	"io/ioutil"
)

func main(){
	urlPtr:=flag.String("url","","the url")
	methodPtr:=flag.String("method","","the method")
	dataPtr:=flag.String("data","","the data")
	yearPtr:=flag.String("year","","the year")
	flag.Parse()
	switch{
		case *methodPtr=="list":
			Get(*urlPtr)
    case *methodPtr=="List":
      Get(*urlPtr)
		case *methodPtr=="create":
			Post(*urlPtr,*dataPtr)
    case *methodPtr=="Create":
      Post(*urlPtr,*dataPtr)
		case *methodPtr=="update":
			Update(*urlPtr)
    case *methodPtr=="Update":
      Update(*urlPtr)
    case *methodPtr=="remove":
      Delete(*urlPtr,*yearPtr)
		case *methodPtr=="Remove":
			Delete(*urlPtr,*yearPtr)
	}
}

func Get(url string){
  req, err:=http.NewRequest("GET",url,nil)
  req.Header.Set("Content-Type","application/json")
  client:=&http.Client{}
    resp, err:=client.Do(req)
    if err!=nil {
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, _:=ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func Post(url string,data string){
  var jsonStr=[]byte(data)
  req, err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonStr))
  req.Header.Set("Content-Type","application/json")
  client:=&http.Client{}
    resp, err:=client.Do(req)
    if err!=nil {
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, _:=ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func Update(url string){
  req, err:=http.NewRequest("PUT",url,nil)
  req.Header.Set("Content-Type","application/json")
  client:=&http.Client{}
    resp, err:=client.Do(req)
    if err!=nil {
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, _:=ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}

func Delete(url string,year string){
  req, err:=http.NewRequest("DELETE",url+"/"+year,nil)
  req.Header.Set("Content-Type","application/json")
  client:=&http.Client{}
    resp, err:=client.Do(req)
    if err!=nil {
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, _:=ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
