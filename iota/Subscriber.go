package iota

/*
#cgo CFLAGS: -I./include -DIOTA_STREAMS_CHANNELS_CLIENT
#cgo LDFLAGS: -L./include -liota_streams_c
#include <channels.h>
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/project-alvarium/go-simulator/configuration"
	"github.com/project-alvarium/go-simulator/simulator/configfile"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const PAYLOAD_LENGTH=1024

type Subscriber struct {
	Subscriber *C.subscriber_t
	Keyload *C.message_links_t
	client *http.Client
}

type AnnResponse struct {
	AnnId string `json:"announcement_id"`
}

func NewSubscriber(nodeConfig *configfile.NodeConfig, subConfig *configfile.SubConfig) Subscriber {
	// Generate Transport client
	transport := C.transport_client_new_from_url(C.CString(nodeConfig.Url))

	var newSub *C.subscriber_t
	cerr := C.sub_new(&newSub, C.CString(subConfig.Seed), C.CString(subConfig.Encoding), PAYLOAD_LENGTH, transport)
	if cerr != C.ERR_OK {
		fmt.Println(errCode(cerr))
	}

	// Generate Subscriber instance
	sub := Subscriber {
		newSub,
		nil,
		http.DefaultClient,
	}

	// Process announcement message
	annAddr := getAnnouncementId(configuration.AuthConsoleUrl)
	address := C.address_from_string(C.CString(annAddr))
	C.sub_receive_announce(sub.Subscriber, address)

	// Fetch sub link and pk for subscription
	var subLink *C.address_t
	var subPk *C.public_key_t

	cerr = C.sub_send_subscribe(&subLink, sub.Subscriber, address)
	if cerr != C.ERR_OK {
		fmt.Println(errCode(cerr))
	}

	cerr = C.sub_get_public_key(&subPk, sub.Subscriber)
	if cerr != C.ERR_OK {
		fmt.Println(errCode(cerr))
	}

	subIdStr := C.get_address_id_str(subLink)
	subPkStr := C.public_key_to_string(subPk)

	log.Println("Sending subscription request... ", C.GoString(subIdStr))
	sub.SendSubscriptionIdToAuthor(
		configuration.AuthConsoleUrl,
		SubscriptionRequestBody(C.GoString(subIdStr), C.GoString(subPkStr)))
	log.Println("Subscription sent")

	// Free generated c strings from mem
	C.drop_str(subIdStr)
	C.drop_str(subPkStr)

	return sub
}

func (sub *Subscriber) InsertKeyload(keyload *C.message_links_t) {
	s := sub
	s.Keyload = keyload
	*sub = *s
}

func (sub *Subscriber) SendMessage(message TangleMessage) {
	messageBytes := C.CBytes([]byte(message.message))
	messageLen := len(message.message)

	var messageLinks C.message_links_t
	log.Println("Sending streams message... ")
	cerr := C.sub_send_signed_packet(
		&messageLinks,
		sub.Subscriber,
		*sub.Keyload,
		nil, 0,
		(*C.uchar) (messageBytes), C.size_t(messageLen))

	if cerr != C.ERR_OK {
		fmt.Println(errCode(cerr))
	}
	msgStr := C.get_address_id_str(messageLinks.msg_link)
	log.Println("Streams message sent ", C.GoString(msgStr))
}

func (sub *Subscriber) Drop() {
	C.sub_drop(sub.Subscriber)
	C.drop_links(*sub.Keyload)
}

func (sub *Subscriber) AwaitKeyload() {
	exists := false
	for exists == false {
		// Gen next message ids to look for existing messages
		var msgIds *C.next_msg_ids_t
		cerr := C.sub_gen_next_msg_ids(&msgIds, sub.Subscriber)
		if cerr != C.ERR_OK {
			fmt.Println(errCode(cerr))
		}

		// Search for keyload message from these ids and try to process it
		var processed C.message_links_t
		cerr = C.sub_receive_keyload_from_ids(&processed, sub.Subscriber, msgIds)
		if cerr != C.ERR_OK {
			fmt.Println("Keyload not found yet... Checking again in 5 seconds...")
			// Loop until keyload is found and processed
			time.Sleep(time.Second * 5)
		} else {
			// Store keyload links for attaching messages to
			sub.InsertKeyload(&processed)
			exists = true
		}
		// Free memory for c msgids object
		C.drop_next_msg_ids(msgIds)
	}
}

func SubscriptionRequestBody(msgid string, pk string) []byte {
	body := "{ \"msgid\": \"" + msgid + "\", \"pk\": \"" + pk + "\" }"
	return []byte(body)
}

func (sub *Subscriber) SendSubscriptionIdToAuthor(url string, body []byte) {
	client := http.Client{}
	data := bytes.NewReader(body)
	req, err := http.NewRequest("POST", url + "/subscribe", data)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
}

func getAnnouncementId(url string) string {
	resp, err := http.Get(url + "/get_announcement_id")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var annResp AnnResponse
	if err := json.Unmarshal(bodyBytes, &annResp); err != nil {
		fmt.Println(err)
	}
	return annResp.AnnId
}

func errCode(err C.err_t) string {
	switch err {
		case C.ERR_OK: return "\nFunction completed Ok"
		case C.ERR_NULL_ARGUMENT: return "\nSTREAMS ERROR: Null argument passed to function"
		case C.ERR_BAD_ARGUMENT: return "\nSTREAMS ERROR: Bad argument passed to function"
		case C.ERR_OPERATION_FAILED: return "\nSTREAMS ERROR: Operation failed to execute properly"
	}
	return "\nError code does not match any provided error options"
}
