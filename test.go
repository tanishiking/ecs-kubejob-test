package main

import (
	"encoding/json"
	"fmt"
	"os"
)


type ECSFormatLog struct {
	Timestamp         string `json:"@timestamp"`
	Loglevel          string `json:"log.level"`
	Message           string `json:"message"`
	ProcessThreadName string `json:"process.thread.name"`
	LoggerName        string `json:"log.logger"`
	EventDataset      string `json:"event.dataset"`
	EcsVersion        string `json:"ecs.version"`
	ServiceName       string `json:"service.name"`
	StackTrace        string `json:"error.stack_trace"`
}

func main() {
	b := []byte(`{"@timestamp":"2021-11-20T18:57:42.615Z", "log.level":"ERROR", "message":"foo", "ecs.version": "1.2.0","service.name":"my-application","service.node.name":"my-application-cluster-node","event.dataset":"my-application","process.thread.name":"main","log.logger":"Hello$","error.type":"java.lang.Exception","error.message":"foo","error.stack_trace":"java.lang.Exception: foo\n\tat Hello$.delayedEndpoint$Hello$1(Hello.scala:4)\n\tat Hello$delayedInit$body.apply(Hello.scala:3)\n\tat scala.Function0.apply$mcV$sp(Function0.scala:39)\n\tat scala.Function0.apply$mcV$sp$(Function0.scala:39)\n\tat scala.runtime.AbstractFunction0.apply$mcV$sp(AbstractFunction0.scala:17)\n\tat scala.App.$anonfun$main$1(App.scala:76)\n\tat scala.App.$anonfun$main$1$adapted(App.scala:76)\n\tat scala.collection.IterableOnceOps.foreach(IterableOnce.scala:563)\n\tat scala.collection.IterableOnceOps.foreach$(IterableOnce.scala:561)\n\tat scala.collection.AbstractIterable.foreach(Iterable.scala:919)\n\tat scala.App.main(App.scala:76)\n\tat scala.App.main$(App.scala:74)\n\tat Hello$.main(Hello.scala:3)\n\tat Hello.main(Hello.scala)\n"}`)

	var ecsFormatLog ECSFormatLog
	if err := json.Unmarshal([]byte(b), &ecsFormatLog); err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", ecsFormatLog)
	fmt.Fprintf(os.Stdout, "%s", ecsFormatLog)
}
