// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type MemPoolTx struct {
	From       string `json:"from"`
	Gas        string `json:"gas"`
	GasPrice   string `json:"gasPrice"`
	Hash       string `json:"hash"`
	Input      string `json:"input"`
	Nonce      string `json:"nonce"`
	To         string `json:"to"`
	Value      string `json:"value"`
	V          string `json:"v"`
	R          string `json:"r"`
	S          string `json:"s"`
	PendingFor string `json:"pendingFor"`
	QueuedFor  string `json:"queuedFor"`
	Pool       string `json:"pool"`
}
