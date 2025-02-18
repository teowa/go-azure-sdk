package apisbytag

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeGraphql   ApiType = "graphql"
	ApiTypeHTTP      ApiType = "http"
	ApiTypeSoap      ApiType = "soap"
	ApiTypeWebsocket ApiType = "websocket"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeGraphql),
		string(ApiTypeHTTP),
		string(ApiTypeSoap),
		string(ApiTypeWebsocket),
	}
}

func parseApiType(input string) (*ApiType, error) {
	vals := map[string]ApiType{
		"graphql":   ApiTypeGraphql,
		"http":      ApiTypeHTTP,
		"soap":      ApiTypeSoap,
		"websocket": ApiTypeWebsocket,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiType(input)
	return &out, nil
}

type BearerTokenSendingMethods string

const (
	BearerTokenSendingMethodsAuthorizationHeader BearerTokenSendingMethods = "authorizationHeader"
	BearerTokenSendingMethodsQuery               BearerTokenSendingMethods = "query"
)

func PossibleValuesForBearerTokenSendingMethods() []string {
	return []string{
		string(BearerTokenSendingMethodsAuthorizationHeader),
		string(BearerTokenSendingMethodsQuery),
	}
}

func parseBearerTokenSendingMethods(input string) (*BearerTokenSendingMethods, error) {
	vals := map[string]BearerTokenSendingMethods{
		"authorizationheader": BearerTokenSendingMethodsAuthorizationHeader,
		"query":               BearerTokenSendingMethodsQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BearerTokenSendingMethods(input)
	return &out, nil
}

type ProductState string

const (
	ProductStateNotPublished ProductState = "notPublished"
	ProductStatePublished    ProductState = "published"
)

func PossibleValuesForProductState() []string {
	return []string{
		string(ProductStateNotPublished),
		string(ProductStatePublished),
	}
}

func parseProductState(input string) (*ProductState, error) {
	vals := map[string]ProductState{
		"notpublished": ProductStateNotPublished,
		"published":    ProductStatePublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductState(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
	ProtocolWs    Protocol = "ws"
	ProtocolWss   Protocol = "wss"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolHTTP),
		string(ProtocolHTTPS),
		string(ProtocolWs),
		string(ProtocolWss),
	}
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"http":  ProtocolHTTP,
		"https": ProtocolHTTPS,
		"ws":    ProtocolWs,
		"wss":   ProtocolWss,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}
