package ProtobufClient

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// TestDeserializingClientMessage attempts to deserialize a client message with a Body
// type ClientIdentificationMessage. This should be the default way to pass the
// ClientIdentificationMessage.
func TestDeserializingClientMessage(t *testing.T) {
	msg := &ClientMessage{
		Body: &ClientMessage_Identification{
			Identification: &ClientIdentificationMessage{
				Useragent: "Example",
				Revision:  "1",
			},
		},
	}
	msgData, _ := proto.Marshal(msg)
	_, err := DeserializeClientMessage(msgData)
	assert.Nil(t, err, "There was an error deserializing the ClientMessage: %s", err)
}

// TestDeserializingClientIdentificationMessage confirms that DeserializeClientMessage falls
// back to deserializing a ClientIdentificationMessage if the byte buffer passed is not a
// ClientMessage.
func TestDeserializingClientIdentificationMessage(t *testing.T) {
	msg := &ClientIdentificationMessage{
		Useragent: "Example",
		Revision:  "1",
	}
	msgData, _ := proto.Marshal(msg)
	_, err := DeserializeClientMessage(msgData)
	assert.Nil(t, err, "There was an error deserializing the ClientIdentificationMessage as a ClientMessage: %s", err)
}
