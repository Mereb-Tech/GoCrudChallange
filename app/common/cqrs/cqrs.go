/*
Package cqrs provides a generic interface for handling commands and queries.
*/
package cqrs

import ierr "github.com/beka-birhanu/GoCrudChallange/domain/common"

// Handler defines a generic interface for handling requests (commands or queries).
//
// Type Parameters:
// - Req: The type of the request to be handled.
// - Res: The type of the result returned after handling the request.
type Handler[Req any, Res any] interface {
	// Handle processes the provided request and returns the result and an error.
	Handle(request Req) (Res, ierr.IErr)
}
