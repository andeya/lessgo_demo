package Home

import (
    "time"

    "github.com/lessgo/lessgo"
)

var WebSocketHandle = lessgo.ApiHandler{
    Desc:   "websocket",
    Types:  []string{"SOCKET"},
    Params: []lessgo.Param{},
    Handler: func(c lessgo.Context) error {
        for {
            var req interface{}
            if err := c.WsRecvJSON(&req); err != nil {
                return err
            }
            c.Logger().Info("req: %v", req)
            if _, err := c.WsSendJSON(map[string]string{"back": time.Now().String()}); err != nil {
                return err
            }
        }
        return nil
    },
}.Reg()
