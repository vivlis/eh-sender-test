package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/vivlis/eh-sender-test/model"
)

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), 20000*time.Second)
	// defer cancel()

	cs := `Endpoint=sb://eh-rdaas-shared-dev-ams.servicebus.windows.net/;SharedAccessKeyName=SendAndListenPolicy;SharedAccessKey=ApkGHoEYTanfJTxFPkXRNlubgufO+4689+AEhBm21BA=;EntityPath=rdh-valfrm-in2`

	clProd, err := azeventhubs.NewProducerClientFromConnectionString(cs, "", nil)
	if err != nil {
		log.Println(`Error: ` + err.Error())
		panic(err)
	}

	data, props := getMsg()

	batchSize := 100
	batchMsg := 0
	msgs := 1000

	batch, err := clProd.NewEventDataBatch(context.Background(), nil)

	for i := 1; i <= msgs; i++ {
		batchMsg++
		// err = hub.Send(ctx, &eventhub.Event{
		// 	Data:       []byte(data),
		// 	ID:         fmt.Sprintf(`%d`, i),
		// 	Properties: props,
		// })

		// events = append(events, &eventhub.Event{
		// 	Data:       []byte(data),
		// 	ID:         fmt.Sprintf(`%d`, i),
		// 	Properties: props,
		// })
		// log.Printf("Added %d message\n", i)

		if batch == nil {
			batch, err = clProd.NewEventDataBatch(context.Background(), nil)
			if err != nil {
				log.Printf("Failed to create batch: %s", err.Error())
				return
			}
		}
		// create event
		ct := "application/json"
		id := fmt.Sprintf(`%d`, i)
		ev := &azeventhubs.EventData{
			ContentType: &ct,
			Body:        []byte(data),
			Properties:  props,
			MessageID:   &id,
		}

		if err := batch.AddEventData(ev, nil); err != nil {
			log.Printf("Failed to add event: %s", err.Error())
			return
		}

		if batchMsg == batchSize || i == msgs {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			err = clProd.SendEventDataBatch(ctx, batch, nil)
			if err != nil {
				log.Printf("Failed to send message: %s", err.Error())
			}

			log.Printf("Batch of %d message\n", batch.NumEvents())

			// clear event list
			batchMsg = 0
			batch = nil
			log.Printf("Message of %d messages %d\n", i, msgs)
		}
	}
}

func getMsg() (string, map[string]interface{}) {

	prop := map[string]interface{}{`type`: `security`}

	dataSec := &model.SecurityData{
		ISINCode:    "IE123",
		WKN:         "XETR",
		MarketPlace: "XPEX",
		Market:      "XETR",
	}
	dataProd := &model.ProductData{
		MarketPlace: "XPEX",
		Market:      "XETR",
	}

	log.Printf("%#v", dataProd)
	log.Printf("%#v", dataSec)
	m, err := json.Marshal(dataSec)
	if err != nil {
		panic(err)
	}
	return string(m), prop
}
