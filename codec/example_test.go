package codec_test

import (
	"fmt"

	"github.com/malivvan/tempest"
	"github.com/malivvan/tempest/codec/gob"
	"github.com/malivvan/tempest/codec/json"
)

func Example() {
	// The examples below show how to set up all the codecs shipped with Tempest.
	// Proper error handling left out to make it simple.
	var gobDb, _ = tempest.Open("gob.db", tempest.Codec(gob.Codec))
	var jsonDb, _ = tempest.Open("json.db", tempest.Codec(json.Codec))

	fmt.Printf("%T\n", gobDb.Codec())
	fmt.Printf("%T\n", jsonDb.Codec())

	// Output:
	// *gob.gobCodec
	// *json.jsonCodec
}
