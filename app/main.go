package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ibm-messaging/mq-golang/ibmmq"
)

func main() {

	qmgr, err := ibmmq.Connx("QM1", newMqsno())
	if err != nil {
		fmt.Printf("Failed connecting to MQ host/manager. %s \n", err.Error())
		return
	}

	devqueue1, _ := createPutQueueObject(qmgr, "DEV.QUEUE.1")
	msgID, err := putMessageToQueue(devqueue1, []byte("here is a test message"))
	if err != nil {
		fmt.Printf("failed to put message to queue %s with error %s", devqueue1.Name, err)
	}
	fmt.Printf("MsgID of put message: %s\n", msgID)
	close(devqueue1)

	devqueue1Get, _ := createGetQueueObject(qmgr, "DEV.QUEUE.1")

	getmqmd := ibmmq.NewMQMD()
	gmo := ibmmq.NewMQGMO()

	gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT
	gmo.Options |= ibmmq.MQGMO_WAIT
	gmo.WaitInterval = 3 * 1000 //milliseconds

	gmo.MatchOptions = ibmmq.MQMO_MATCH_MSG_ID
	msgID = "414d5120514d312020202020202020207d822263013f0040"
	getmqmd.MsgId, _ = hex.DecodeString(msgID)

	buffer := make([]byte, 1024)

	datalen, err := devqueue1Get.Get(getmqmd, gmo, buffer)
	if err != nil {
		mqret := err.(*ibmmq.MQReturn)
		if mqret.MQRC != ibmmq.MQRC_NO_MSG_AVAILABLE {
			fmt.Printf("Error from get: %s", err)
			return
		}
	}

	fmt.Printf("Got message %s", strings.TrimSpace(string(buffer[:datalen])))

	disc(qmgr)

}

// Disconnect from the queue manager
func disc(qMgrObject ibmmq.MQQueueManager) error {
	err := qMgrObject.Disc()
	if err != nil {
		return err
	}
	fmt.Printf("Disconnected from queue manager %s\n", qMgrObject.Name)
	return nil
}

// Close the queue if it was opened
func close(object ibmmq.MQObject) error {
	err := object.Close(0)
	if err != nil {
		return err
	}
	return nil
}

func createPutQueueObject(qmgr ibmmq.MQQueueManager, queueName string) (ibmmq.MQObject, error) {
	mqod := ibmmq.NewMQOD()

	// We have to say how we are going to use this queue. In this case, to PUT
	// messages. That is done in the openOptions parameter.
	openOptions := ibmmq.MQOO_OUTPUT

	// Opening a QUEUE (rather than a Topic or other object type) and give the name
	mqod.ObjectType = ibmmq.MQOT_Q
	mqod.ObjectName = queueName

	qObject, err := qmgr.Open(mqod, openOptions)
	if err != nil {
		return qObject, err
	}
	return qObject, nil

}

func createGetQueueObject(qmgr ibmmq.MQQueueManager, queueName string) (ibmmq.MQObject, error) {
	mqod := ibmmq.NewMQOD()

	// We have to say how we are going to use this queue. In this case, to PUT
	// messages. That is done in the openOptions parameter.
	openOptions := ibmmq.MQOO_INPUT_EXCLUSIVE

	// Opening a QUEUE (rather than a Topic or other object type) and give the name
	mqod.ObjectType = ibmmq.MQOT_Q
	mqod.ObjectName = queueName

	qObject, err := qmgr.Open(mqod, openOptions)
	if err != nil {
		return qObject, err
	}
	return qObject, nil

}

func newMqsno() *ibmmq.MQCNO {
	var mqcno = ibmmq.NewMQCNO()
	mqcno.Options = ibmmq.MQCNO_CLIENT_BINDING

	var mqcd = ibmmq.NewMQCD()
	mqcd.ChannelName = "DEV.APP.SVRCONN"
	mqcd.ConnectionName = "mqseries(1414)"
	mqcd.SSLCipherSpec = "TLS_RSA_WITH_AES_128_CBC_SHA256"
	mqcd.CertificateLabel = "ibmwebspheremqapp"
	mqcd.SSLClientAuth = ibmmq.MQSCA_REQUIRED
	mqcno.ClientConn = mqcd

	var mqsco = ibmmq.NewMQSCO()
	mqsco.KeyRepository = "/mnt/mqm/certs/client"
	mqcno.SSLConfig = mqsco

	return mqcno
}

func putMessageToQueue(queue ibmmq.MQObject, msg []byte) (msgID string, err error) {

	// The PUT requires control structures, the Message Descriptor (MQMD)
	// and Put Options (MQPMO). Create those with default values.
	putmqmd := ibmmq.NewMQMD()
	pmo := ibmmq.NewMQPMO()

	// The default options are OK, but it's always
	// a good idea to be explicit about transactional boundaries as
	// not all platforms behave the same way.
	pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT

	// Tell MQ what the message body format is. In this case, a text string
	putmqmd.Format = "MQSTR"

	// Now put the message to the queue
	err = queue.Put(putmqmd, pmo, msg)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Put message to", strings.TrimSpace(queue.Name))

	return hex.EncodeToString(putmqmd.MsgId), nil

}
