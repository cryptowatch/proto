package ProtobufClient

import "github.com/golang/protobuf/proto"

// DeserializeClientMessage decodes the client message and switches based on the message format.
// This is needed because older clients may be pushing ClientIdentificationMessages not wrapped
// in the ClientMessage envelope.
func DeserializeClientMessage(msgData []byte) (msg ClientMessage, err error) {
	// Unmarshal client message envelope
	err = proto.Unmarshal(msgData, &msg)
	if err != nil {
		var identificationMsg ClientIdentificationMessage
		err = proto.Unmarshal(msgData, &identificationMsg)
		msg = ClientMessage{
			Body: &ClientMessage_Identification{
				Identification: &identificationMsg,
			},
		}
	}
	return msg, err
}
