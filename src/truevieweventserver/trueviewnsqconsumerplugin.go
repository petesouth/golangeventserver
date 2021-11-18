package truevieweventserver

import "github.com/nsqio/go-nsq"
import "log"
import "fmt"
import "math/rand"
import "sync"
import "io/ioutil"

const TRUEVIEW_CONSUMER_PLUGIN string = "trueview_consumer_plugin"

/**
 * Main exported function.  Runs a configuration against values in an influx Database.
 */
func TrueviewConsumerPlugin(configuration Configuration) error {

	cfg := nsq.NewConfig()

	channel := fmt.Sprintf("TrueviewConsumerPlugin%06d#ephemeral", rand.Int()%999999)
	topic := configuration.ConsumerPluginConfig.Topic
	nsqdTcpAddress := configuration.ConsumerPluginConfig.NsqdTcpAddress
	messageBatchRemoveSize := configuration.ConsumerPluginConfig.MessageBatchRemoveSize
	outputDir := configuration.ConsumerPluginConfig.OutputDir

	wg := &sync.WaitGroup{}
	wg.Add(1)

	consumer, errConsumer := nsq.NewConsumer(topic, channel, cfg)
	if errConsumer != nil {
		log.Println("Error: creating consumer in trueviewConsumerPlugin:", errConsumer)
		return errConsumer
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		influxNResponseNSQContainer := InfluxNResponseNSQContainer{}
		err := JsonStringStruct(string(message.Body), &influxNResponseNSQContainer)
		if err != nil {
			log.Printf("Error converting incomming NSQ message to InfluxResponseNSQContainer", err)
		}

		dataAsString, err := StructToJson(influxNResponseNSQContainer.Data)
		if err != nil {
			log.Printf("Error converting incomming NSQ message to InfluxResponseNSQContainer", err)
		}

		outputDirLen := len(outputDir)

		for iOutputDir := 0; iOutputDir < outputDirLen; iOutputDir += 1 {
			theTargetDir := outputDir[iOutputDir]
			incommingFileWritePath := theTargetDir + "/" + influxNResponseNSQContainer.Name

			log.Printf("Got a message writing JSON Report to %s:", incommingFileWritePath)

			ioutil.WriteFile(incommingFileWritePath, []byte(dataAsString), 0666)
		}

		// Decrement the message batch size.. Once done.. This consumer exits.. For rexecution by plugin runner.
		messageBatchRemoveSize = messageBatchRemoveSize - 1

		if messageBatchRemoveSize < 1 {
			//wg.Done()
		}
		return nil
	}))

	err := consumer.ConnectToNSQLookupd(nsqdTcpAddress)
	if err != nil {
		log.Println("Error: connectiong to NSQ in trueviewConsumerPlugin:", err)
		return err
	}
	wg.Wait()

	consumer.Stop()

	return nil
}
