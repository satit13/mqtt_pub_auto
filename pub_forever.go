package main

import (
"fmt"
//import the Paho Go MQTT library
MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
"os"
"time"
"math/rand"
	"encoding/json"
)


//define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
fmt.Printf("TOPIC: %s\n", msg.Topic())
fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
//create a ClientOptions struct setting the broker address, clientid, turn
//off trace output and set the default message handler
opts := MQTT.NewClientOptions().AddBroker("tcp://nava.work:1883")
opts.SetClientID("go-simple")
opts.SetDefaultPublishHandler(f)

//create and start a client using the above ClientOptions
c := MQTT.NewClient(opts)
if token := c.Connect(); token.Wait() && token.Error() != nil {
panic(token.Error())
}

//subscribe to the topic /go-mqtt/sample and request messages to be delivered
//at a maximum qos of zero, wait for the receipt to confirm the subscription


//if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
//fmt.Println(token.Error())
//os.Exit(1)
//}

//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
//from the server after sending each message



	type payload struct {

		Vendorid int
		Value int
		Jobid int
		Siteid int
		Cardid int
		Timestamp string
	}




for {


p := payload{
	Cardid : rand.Intn(100),
	Jobid  : rand.Intn(5),
	Siteid : rand.Intn(100),
	Value  : rand.Intn(100)*10,
	Vendorid : rand.Intn(2000),
	Timestamp :time.Now().Format(time.RFC3339),
}

	ret ,_ :=  json.Marshal(p)
	val := string(ret)


text := val
	fmt.Println(text)
token := c.Publish("go-mqtt/sample", 0, false, text)
token.Wait()
time.Sleep(20 * time.Second)

}



//unsubscribe from /go-mqtt/sample
if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
fmt.Println(token.Error())
os.Exit(1)
}

c.Disconnect(250)
}
