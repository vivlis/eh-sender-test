package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

	// default batch size
	batchSize := 10
	// get batch size from env
	size, _ := strconv.Atoi(os.Getenv("BatchSize"))
	if size != 0 {
		batchSize = size
	}
	batchMsg := 0

	// default number of msgs
	msgs := 1
	// get number of msgs from env
	numMsg, _ := strconv.Atoi(os.Getenv("NoEvents"))
	if numMsg != 0 {
		msgs = numMsg
	}

	// default number of partitions
	maxNumPart := 0
	// get number of partitions from env
	numPart, _ := strconv.Atoi(os.Getenv("MaxNoPart"))
	if numPart != 0 {
		maxNumPart = numPart
	}

	var pID string
	batch := &azeventhubs.EventDataBatch{}
	batch = nil
	// time stamp for id
	t := time.Now().Format("2006-01-02-15:04:05")

	for i := 1; i <= msgs; i++ {
		batchMsg++

		if batch == nil {

			bop := &azeventhubs.EventDataBatchOptions{}
			if maxNumPart > 0 {
				pID = fmt.Sprintf(`%d`, rand.Intn(maxNumPart))
				bop.PartitionID = &pID
			}

			batch, err = clProd.NewEventDataBatch(context.Background(), bop)
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
