package servicebus

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func NewReceiver(
	connString string,
	queue string,
) (*azservicebus.Receiver, *azservicebus.Client, error) {

	fmt.Println("Connecting To Queue:", queue)

	client, err := azservicebus.NewClientFromConnectionString(
		connString,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	receiver, err := client.NewReceiverForQueue(
		queue,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Receiver Created Successfully")

	return receiver, client, nil
}

func CloseClient(client *azservicebus.Client) {
	client.Close(context.Background())
}
