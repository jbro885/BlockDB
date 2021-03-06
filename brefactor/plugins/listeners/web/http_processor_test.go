package web

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpListener_Handle(t *testing.T) {
	httpListener := HttpListener{
		Config: HttpListenerConfig{MaxContentLength: 1e7},
	}
	data := `{
		"op": "insert",
		"collection": "sample_collection",
		"data": {
			"name": "fudan",
			"address": {
				"city": "Shanghai",
				"road": "xxx"
			},
			"logo": {
				"url": "http://a.png"
			},
			"teachers": [
				"T1", "T2", "T3"
			]
		},
		"public_key": "02c3a28b7e83c83f90c56861210b418dfc7a7302a9449c4c4597eb6e0ce415b944"
    }`

	priBytes, _ := hex.DecodeString("42f909a1a4cc546f270306b1b69c45434a1e37cddf2d834ea377cd5e92c5d3d5")
	pri, err := crypto.UnmarshalSecp256k1PrivateKey(priBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	dataBytes := []byte(Normalize(data))
	hash := sha256.Sum256(dataBytes)
	signature, _ := pri.Sign(hash[:])
	message := &Message{
		OpStr:     dataBytes,
		OpHash:    hex.EncodeToString(hash[:]),
		PublicKey: "02c3a28b7e83c83f90c56861210b418dfc7a7302a9449c4c4597eb6e0ce415b944",
		Signature: hex.EncodeToString(signature),
	}
	msg, _ := json.Marshal(message)
	req := httptest.NewRequest(http.MethodPost, "http://url.com", bytes.NewReader(msg))
	w := httptest.NewRecorder()
	httpListener.Handle(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("Handle function work incorrectly.")
	}
}
