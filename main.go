package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/vivlis/eh-sender-test/model"
)

func init() {

}

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), 20000*time.Second)
	// defer cancel()

	cs := os.Getenv("EHProducer")

	clProd, err := azeventhubs.NewProducerClientFromConnectionString(cs, "", nil)
	if err != nil {
		log.Println(`Error: ` + err.Error())
		panic(err)
	}
	defer clProd.Close(context.Background())

	data, props := getMsg()

	batchSize := 10
	batchMsg := 0
	msgs := 1
	maxNumPart := 4

	pID := fmt.Sprintf(`%d`, rand.Intn(maxNumPart))
	batch, _ := clProd.NewEventDataBatch(context.Background(), &azeventhubs.EventDataBatchOptions{
		PartitionID: &pID,
	})

	// time stamp for id
	t := time.Now().Format("2006-01-02-15:04:05")

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
			pID = fmt.Sprintf(`%d`, rand.Intn(maxNumPart))
			// pID = `3`
			batch, err = clProd.NewEventDataBatch(context.Background(), &azeventhubs.EventDataBatchOptions{
				PartitionID: &pID,
			})
			if err != nil {
				log.Printf("Failed to create batch: %s", err.Error())
				return
			}
		}
		// create event
		ct := "application/json"
		id := fmt.Sprintf(`%s-%d`, t, i)
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

			log.Printf("Batch of %d message partition, %s\n", batch.NumEvents(), pID)

			// clear event list
			batchMsg = 0
			batch = nil
			log.Printf("Message of %d messages %d, partition %s\n", i, msgs, pID)
		}
	}
}

func getMsg() (string, map[string]interface{}) {

	prop := map[string]interface{}{`type`: `security`}

	sec := model.SecurityData{ISIN: "", SecurityId: "123", Market: "XMAL", CreationTimestamp: "2002-01-01.12:12::12346578"}
	dt := model.Data{Security: sec}
	msg := model.SecurityRule{Service: `cms`, RuleSet: `security`, Version: `1.0.1`, Data: dt}

	log.Printf("%+v", msg)
	m, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(m), prop
}
