// Copyright (C) MongoDB, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package bsonx

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

// DefaultRegistry is the default bsoncodec.Registry. It contains the default codecs and the
// primitive codecs.
var DefaultRegistry = NewRegistryBuilder().Build()

// NewRegistryBuilder creates a new RegistryBuilder configured with the default encoders and
// decoders from the bsoncodec.DefaultValueEncoders and bsoncodec.DefaultValueDecoders types and the
// PrimitiveCodecs type in this package.
func NewRegistryBuilder() *bsoncodec.RegistryBuilder {
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	bson.PrimitiveCodecs{}.RegisterPrimitiveCodecs(rb)
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}