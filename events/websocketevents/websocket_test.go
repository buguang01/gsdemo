package websocketevents_test

import (
	websocketevents "GameService/events/websocketevents"
	"fmt"
	"go/importer"
	"testing"
)

func TestStruct(t *testing.T) {
	_ = websocketevents.WsocketListenEvent{}
	pkg, _ := importer.Default().Import(`go/importer`)
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
}
