// +build !wechat_debug

package callback

import (
	"io"
	"net/http"

	"github.com/lib4dev/wechat/util"
)

func DebugPrintRequest(r *http.Request) {}

func DebugPrintRequestMessage(msg []byte) {}

func EncodeXMLResponseMessage(w io.Writer, msg map[string]string) (err error) {
	return util.EncodeXMLFromMap(w, msg, "xml")
}
