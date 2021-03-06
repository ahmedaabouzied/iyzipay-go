package iyzipay

type CreateApprovalRequest struct {
	Locale               string `json:"locale,omitempty"`
	ConversationId       string `json:"conversationId,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
}

func (request CreateApprovalRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentTransactionId", request.PaymentTransactionId)

	return pkiBuilder.getRequestString()
}
