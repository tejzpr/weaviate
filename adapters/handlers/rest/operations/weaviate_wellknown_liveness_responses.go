//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateWellknownLivenessNoContentCode is the HTTP code returned for type WeaviateWellknownLivenessNoContent
const WeaviateWellknownLivenessNoContentCode int = 204

/*WeaviateWellknownLivenessNoContent The application is able to respond to HTTP requests

swagger:response weaviateWellknownLivenessNoContent
*/
type WeaviateWellknownLivenessNoContent struct {
}

// NewWeaviateWellknownLivenessNoContent creates WeaviateWellknownLivenessNoContent with default headers values
func NewWeaviateWellknownLivenessNoContent() *WeaviateWellknownLivenessNoContent {

	return &WeaviateWellknownLivenessNoContent{}
}

// WriteResponse to the client
func (o *WeaviateWellknownLivenessNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}
