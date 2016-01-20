package main

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
)

func main() {

	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)


	type payload struct {

		Vendorid int
		Value int
		Jobid int
		Siteid int
		Cardid int
		Timestamp string
	}

	p := payload{
		Cardid 		:	rand.Intn(100),
		Jobid		:	rand.Intn(100),
		Vendorid 	:	rand.Intn(200),
		Value		:	rand.Intn(100),
		Siteid		:	rand.Intn(100),
		Timestamp	:	time.Now().Format(time.RFC3339),
	}


	fmt.Println(p.Value)
	fmt.Println(p.Timestamp)

//	p.cardid = rand.Intn(100)
//	p.jobid  = rand.Intn(5)
//	p.siteid = rand.Intn(100)
//	p.value  = rand.Intn(100)*10
//	p.vendorid = rand.Intn(2000)
//	p.timestamp =time.Now().Format(time.RFC3339)


	//q := []byte(p)

	fmt.Println(p)
	ret ,err :=  json.Marshal(p)
	val := string(ret)

	fmt.Println(val)




	//fmt.Println(time.Now().Format(time.RFC850))
}

