package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func ExampleRetrieveCards() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	card := builder.CreateCardBuilder()

	request := iyzipay.RetrieveCardListRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		CardUserKey:    card["cardUserKey"].(string),
	}

	rawResponse := iyzipay.CardList{}.Retrieve(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	fmt.Println(response["status"])
	fmt.Println(response["locale"])
	fmt.Println(response["conversationId"])
	fmt.Println(response["errorCode"])
	fmt.Println(response["errorMessage"])
	fmt.Println(response["errorGroup"])
	//Output:
	//success
	//tr
	//123456789
	//<nil>
	//<nil>
	//<nil>
}
