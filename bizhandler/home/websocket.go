package home

import (
    "time"

    "github.com/lessgo/lessgo"
)

var WebSocket = lessgo.ApiHandler{
    Desc:   "websocket",
    Method: "WS",
    Params: []lessgo.Param{},
    Handler: func(c *lessgo.Context) error {
        for {
            var req interface{}
            if err := c.WsRecvJSON(&req); err != nil {
                return err
            }
            c.Log().Info("req: %v", req)
            if _, err := c.WsSendJSON(map[string]string{"back": time.Now().String()}); err != nil {
                return err
            }
        }
        return nil
    },
}.Reg()
