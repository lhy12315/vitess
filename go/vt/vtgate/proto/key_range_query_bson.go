// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/youtube/vitess/go/bson"
	"github.com/youtube/vitess/go/bytes2"
	"github.com/youtube/vitess/go/vt/key"
)

// MarshalBson bson-encodes KeyRangeQuery.
func (keyRangeQuery *KeyRangeQuery) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	bson.EncodeString(buf, "Sql", keyRangeQuery.Sql)
	// map[string]interface{}
	{
		bson.EncodePrefix(buf, bson.Object, "BindVariables")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v1 := range keyRangeQuery.BindVariables {
			bson.EncodeInterface(buf, _k, _v1)
		}
		lenWriter.Close()
	}
	bson.EncodeString(buf, "Keyspace", keyRangeQuery.Keyspace)
	// []key.KeyRange
	{
		bson.EncodePrefix(buf, bson.Array, "KeyRanges")
		lenWriter := bson.NewLenWriter(buf)
		for _i, _v2 := range keyRangeQuery.KeyRanges {
			_v2.MarshalBson(buf, bson.Itoa(_i))
		}
		lenWriter.Close()
	}
	keyRangeQuery.TabletType.MarshalBson(buf, "TabletType")
	// *Session
	if keyRangeQuery.Session == nil {
		bson.EncodePrefix(buf, bson.Null, "Session")
	} else {
		(*keyRangeQuery.Session).MarshalBson(buf, "Session")
	}
	bson.EncodeBool(buf, "NotInTransaction", keyRangeQuery.NotInTransaction)

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into KeyRangeQuery.
func (keyRangeQuery *KeyRangeQuery) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for KeyRangeQuery", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Sql":
			keyRangeQuery.Sql = bson.DecodeString(buf, kind)
		case "BindVariables":
			// map[string]interface{}
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for keyRangeQuery.BindVariables", kind))
				}
				bson.Next(buf, 4)
				keyRangeQuery.BindVariables = make(map[string]interface{})
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v1 interface{}
					_v1 = bson.DecodeInterface(buf, kind)
					keyRangeQuery.BindVariables[_k] = _v1
				}
			}
		case "Keyspace":
			keyRangeQuery.Keyspace = bson.DecodeString(buf, kind)
		case "KeyRanges":
			// []key.KeyRange
			if kind != bson.Null {
				if kind != bson.Array {
					panic(bson.NewBsonError("unexpected kind %v for keyRangeQuery.KeyRanges", kind))
				}
				bson.Next(buf, 4)
				keyRangeQuery.KeyRanges = make([]key.KeyRange, 0, 8)
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					bson.SkipIndex(buf)
					var _v2 key.KeyRange
					_v2.UnmarshalBson(buf, kind)
					keyRangeQuery.KeyRanges = append(keyRangeQuery.KeyRanges, _v2)
				}
			}
		case "TabletType":
			keyRangeQuery.TabletType.UnmarshalBson(buf, kind)
		case "Session":
			// *Session
			if kind != bson.Null {
				keyRangeQuery.Session = new(Session)
				(*keyRangeQuery.Session).UnmarshalBson(buf, kind)
			}
		case "NotInTransaction":
			keyRangeQuery.NotInTransaction = bson.DecodeBool(buf, kind)
		default:
			bson.Skip(buf, kind)
		}
	}
}
