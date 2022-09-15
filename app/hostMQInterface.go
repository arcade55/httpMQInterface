package main

/*

func getSalesCustomerDetails() {
	mqValues, _ := getMQValues()
	//qMgrObject := qmgrConnect(mqValues.qmgr)

	//defer disc(qMgrObject)

}


func qmgrConnect() (ibmmq.MQQueueManager, error) {

	qmgrName := os.Getenv("QMGR")
	if qmgrName != "" {
		return ibmmq.Conn(qmgrName)
	}
	return ibmmq.MQQueueManager{}, QMGRNameNotFoundErr{}
}

func getMQValues() (mqValues mqValues, err error) {
	//TODO set error returna

	mqValues.qmgr = os.Getenv("QMGR")
	if mqValues.qmgr == "" {
		return

	}
	mqValues.requestQ = os.Getenv("REQUEST_QUEUE")
	if mqValues.requestQ == "" {
		return
	}
	mqValues.replyToQ = os.Getenv("REPLY_TO_QUEUE")
	if mqValues.replyToQ == "" {
		return
	}
	return mqValues, nil

}

/*


qMgrObject, err := ibmmq.Conn(qMgrName)
if err != nil {
	log.Fatalf("Error when connecting to queuemanager: %s with error: %s", qMgrName, err)
}
fmt.Printf("Connected to queue manager %s\n", qMgrName)

defer disc(qMgrObject)



time.Sleep(10 * time.Second)
fmt.Println("waking up from sleep ready to put message")

// The default queue manager and queue to be used. These can be overridden on command line.
// Environment variables can also be used as that works well in a number of common
// container deployment models.
qMgrName := "QM1"
qName := "getSalesCustomerDetails

fmt.Println("Sample AMQSPUT.GO start")

// This is where we connect to the queue manager. It is assumed
// that the queue manager is either local, or you have set the
// client connection information externally eg via a CCDT or the
// MQSERVER environment variable
qMgrObject, err := ibmmq.Conn(qMgrName)
if err != nil {
	log.Fatalf("Error when connectin to queuemanager: %s with error: %s", qMgrName, err)
}
fmt.Printf("Connected to queue manager %s\n", qMgrName)
defer disc(qMgrObject)

// Create the Object Descriptor that allows us to give the queue name
mqod := ibmmq.NewMQOD()

// We have to say how we are going to use this queue. In this case, to PUT
// messages. That is done in the openOptions parameter.
openOptions := ibmmq.MQOO_OUTPUT

// Opening a QUEUE (rather than a Topic or other object type) and give the name
mqod.ObjectType = ibmmq.MQOT_Q
mqod.ObjectName = qName

qObject, err = qMgrObject.Open(mqod, openOptions)
if err != nil {
	log.Fatalf("Error when opening q: %s with error: %s", qName, err)
}
fmt.Println("Opened queue", qObject.Name)
defer close(qObject)

// PUT a message to the queue
// The PUT requires control structures, the Message Descriptor (MQMD)
// and Put Options (MQPMO). Create those with default values.
putmqmd := ibmmq.NewMQMD()
pmo := ibmmq.NewMQPMO()

// The default options are OK, but it's always
// a good idea to be explicit about transactional boundaries as
// not all platforms behave the same way.
pmo.Options = ibmmq.MQPMO_NO_SYNCPOINT

// Tell MQ what the message body format is. In this case, a text string
putmqmd.Format = ibmmq.MQFMT_STRING
putmqmd.ReplyToQ = "dfd"

// And create the contents to include a timestamp just to prove when it was created
msgData := "Hello from Go at " + time.Now().Format(time.RFC3339)

// The message is always sent as bytes, so has to be converted before the PUT.
buffer := []byte(msgData)

// Now put the message to the queue
err = qObject.Put(putmqmd, pmo, buffer)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Put message to", strings.TrimSpace(qObject.Name))
	// Print the MsgId so it can be used as a parameter to amqsget
	fmt.Println("MsgId:" + hex.EncodeToString(putmqmd.MsgId))
}

// Exit with any return code extracted from the failing MQI call.
// Deferred disconnect will happen after the return
mqret := 0
if err != nil {
	mqret = int((err.(*ibmmq.MQReturn)).MQCC)
}
}


// Disconnect from the queue manager
func disc(qMgrObject ibmmq.MQQueueManager) error {
err := qMgrObject.Disc()
if err == nil {
	fmt.Printf("Disconnected from queue manager %s\n", qMgrObject.Name)
} else {
	fmt.Println(err)
}
return err
}

// Close the queue if it was opened
func close(object ibmmq.MQObject) error {
err := object.Close(0)
if err == nil {
	fmt.Println("Closed queue")
} else {
	fmt.Println(err)
}
return err
}

*/
