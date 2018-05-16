package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"fmt"
)

var (
	sqs_result_queue_url   = os.Getenv("SQS_RESULT_QUEUE_URL")
)

/*

Notes
-----

dim1 0
dim2 1
dim3 2
etc

Always pass out tasks using later-ONLY dimension results

*/

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	// note - taken directly from AWS example
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &sqs_result_queue_url,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(36000),  // 10 hours
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	if len(result.Messages) != 0 {

		// TODO - stuff we do to a specific result message

		// 1.) Tracking - how we we know when we're done listning?
		// 2.) What are the variations? (see notes)
		// 3.) (or possible 1.) Send result to result writer

		//

	}


}
