package main

//usa gorilla para crear JSON-rpc

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type RpcApiArg struct {
	Message string
}

type RpcApiResp struct {
	Message string
}

type CadServ struct{}

func (h *CadServ) Length(r *http.Request, args *RpcApiArg, resp *RpcApiResp) error {
	resp.Message = "Tu cadena tiene : " + fmt.Sprintf(" %d caracteres", utf8.RuneCountInString(args.Message))
	return nil
}

func main() {
	fmt.Println("Iniciando servicio")
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "appilcation/json")
	s.RegisterService(new(CadServ), "")
	http.Handle("/rpc", s)
	http.ListenAndServe(":80", nil)
}
