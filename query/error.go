package query

import "errors"

var errWrongType = errors.New("ease-query: wrong type")

var errBedrockWrongMsgId = errors.New("ease-query: bedrock: wrong incoming message id")
var errBedrockWrongPongId = errors.New("ease:query: bedrock: wrong pong id")
var errBedrockWrongMagic = errors.New("ease-query: bedrock: wrong magic")
var errBedrockWrongLen = errors.New("ease-query: bedrock: wrong packet length")
var errBedrockWrongFmt = errors.New("ease-query: bedrock: wrong data format")
