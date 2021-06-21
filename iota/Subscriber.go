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
	transport := C.tsp_client_new_from_url(C.CString(nodeConfig.Url))

	// Generate Subscriber instance
	sub := Subscriber {
		C.sub_new(C.CString(subConfig.Seed), C.CString(subConfig.Encoding), PAYLOAD_LENGTH, transport),
		nil,
		http.DefaultClient,
	}

	// Process announcement message
	annAddr := getAnnouncementId(configuration.AuthConsoleUrl)
	address := C.address_from_string(C.CString(annAddr))
	C.sub_receive_announce(sub.Subscriber, address)

	// Fetch sub link and pk for subscription
	subLink := C.sub_send_subscribe(sub.Subscriber, address)
	subPk := C.sub_get_public_key(sub.Subscriber)

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

	C.sub_send_signed_packet(
		sub.Subscriber,
		*sub.Keyload,
		nil, 0,
		(*C.uchar) (messageBytes), C.size_t(messageLen))
}

func (sub *Subscriber) Drop() {
	C.sub_drop(sub.Subscriber)
	C.drop_links(*sub.Keyload)
}

func (sub *Subscriber) AwaitKeyload() {
	exists := false
	for exists == false {
		// Gen next message ids to look for existing messages
		msgIds := C.sub_gen_next_msg_ids(sub.Subscriber)
		// Search for keyload message from these ids and try to process it
		processed := C.sub_receive_keyload_from_ids(sub.Subscriber, msgIds)
		// Free memory for c msgids object
		C.drop_next_msg_ids(msgIds)

		if processed != nil {
			// Store keyload links for attaching messages to
			sub.InsertKeyload(processed)
			exists = true
		} else {
			// Loop until keyload is found and processed
			time.Sleep(time.Second)
		}
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
