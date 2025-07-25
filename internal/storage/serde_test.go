// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"io"
	"reflect"
	"testing"
	"unsafe"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
)

type MockRecordWriter struct {
	writefn func(Record) error
	closefn func() error
}

var _ RecordWriter = (*MockRecordWriter)(nil)

func (w *MockRecordWriter) Write(record Record) error {
	return w.writefn(record)
}

func (w *MockRecordWriter) Close() error {
	return w.closefn()
}

func (w *MockRecordWriter) GetWrittenUncompressed() uint64 {
	return 0
}

func TestSerDe(t *testing.T) {
	type args struct {
		dt schemapb.DataType
		v  any
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 bool
	}{
		{"test bool", args{dt: schemapb.DataType_Bool, v: true}, true, true},
		{"test bool null", args{dt: schemapb.DataType_Bool, v: nil}, nil, true},
		{"test bool negative", args{dt: schemapb.DataType_Bool, v: -1}, nil, false},
		{"test int8", args{dt: schemapb.DataType_Int8, v: int8(1)}, int8(1), true},
		{"test int8 null", args{dt: schemapb.DataType_Int8, v: nil}, nil, true},
		{"test int8 negative", args{dt: schemapb.DataType_Int8, v: true}, nil, false},
		{"test int16", args{dt: schemapb.DataType_Int16, v: int16(1)}, int16(1), true},
		{"test int16 null", args{dt: schemapb.DataType_Int16, v: nil}, nil, true},
		{"test int16 negative", args{dt: schemapb.DataType_Int16, v: true}, nil, false},
		{"test int32", args{dt: schemapb.DataType_Int32, v: int32(1)}, int32(1), true},
		{"test int32 null", args{dt: schemapb.DataType_Int32, v: nil}, nil, true},
		{"test int32 negative", args{dt: schemapb.DataType_Int32, v: true}, nil, false},
		{"test int64", args{dt: schemapb.DataType_Int64, v: int64(1)}, int64(1), true},
		{"test int64 null", args{dt: schemapb.DataType_Int64, v: nil}, nil, true},
		{"test int64 negative", args{dt: schemapb.DataType_Int64, v: true}, nil, false},
		{"test float32", args{dt: schemapb.DataType_Float, v: float32(1)}, float32(1), true},
		{"test float32 null", args{dt: schemapb.DataType_Float, v: nil}, nil, true},
		{"test float32 negative", args{dt: schemapb.DataType_Float, v: -1}, nil, false},
		{"test float64", args{dt: schemapb.DataType_Double, v: float64(1)}, float64(1), true},
		{"test float64 null", args{dt: schemapb.DataType_Double, v: nil}, nil, true},
		{"test float64 negative", args{dt: schemapb.DataType_Double, v: -1}, nil, false},
		{"test string", args{dt: schemapb.DataType_String, v: "test"}, "test", true},
		{"test string null", args{dt: schemapb.DataType_String, v: nil}, nil, true},
		{"test string negative", args{dt: schemapb.DataType_String, v: -1}, nil, false},
		{"test varchar", args{dt: schemapb.DataType_VarChar, v: "test"}, "test", true},
		{"test varchar null", args{dt: schemapb.DataType_VarChar, v: nil}, nil, true},
		{"test varchar negative", args{dt: schemapb.DataType_VarChar, v: -1}, nil, false},
		{"test array negative", args{dt: schemapb.DataType_Array, v: "{}"}, nil, false},
		{"test array null", args{dt: schemapb.DataType_Array, v: nil}, nil, true},
		{"test json", args{dt: schemapb.DataType_JSON, v: []byte("{}")}, []byte("{}"), true},
		{"test json null", args{dt: schemapb.DataType_JSON, v: nil}, nil, true},
		{"test json negative", args{dt: schemapb.DataType_JSON, v: -1}, nil, false},
		{"test float vector", args{dt: schemapb.DataType_FloatVector, v: []float32{1.0}}, []float32{1.0}, true},
		{"test float vector null", args{dt: schemapb.DataType_FloatVector, v: nil}, nil, true},
		{"test float vector negative", args{dt: schemapb.DataType_FloatVector, v: []int{1}}, nil, false},
		{"test bool vector", args{dt: schemapb.DataType_BinaryVector, v: []byte{0xff}}, []byte{0xff}, true},
		{"test float16 vector", args{dt: schemapb.DataType_Float16Vector, v: []byte{0xff, 0xff}}, []byte{0xff, 0xff}, true},
		{"test bfloat16 vector", args{dt: schemapb.DataType_BFloat16Vector, v: []byte{0xff, 0xff}}, []byte{0xff, 0xff}, true},
		{"test bfloat16 vector null", args{dt: schemapb.DataType_BFloat16Vector, v: nil}, nil, true},
		{"test bfloat16 vector negative", args{dt: schemapb.DataType_BFloat16Vector, v: -1}, nil, false},
		{"test int8 vector", args{dt: schemapb.DataType_Int8Vector, v: []int8{10}}, []int8{10}, true},
		{"test array of vector", args{dt: schemapb.DataType_ArrayOfVector, v: "{}"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := tt.args.dt
			v := tt.args.v
			builder := array.NewBuilder(memory.DefaultAllocator, serdeMap[dt].arrowType(1))
			serdeMap[dt].serialize(builder, v)
			// assert.True(t, ok)
			a := builder.NewArray()
			got, got1 := serdeMap[dt].deserialize(a, 0, false)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("deserialize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSerDeCopy(t *testing.T) {
	tests := []struct {
		name string
		dt   schemapb.DataType
		v    any
	}{
		{"test string copy", schemapb.DataType_String, "test"},
		{"test string no copy", schemapb.DataType_String, "test"},
		{"test binary copy", schemapb.DataType_JSON, []byte{1, 2, 3}},
		{"test binary no copy", schemapb.DataType_JSON, []byte{1, 2, 3}},
		{"test bool copy", schemapb.DataType_Bool, true},
		{"test bool no copy", schemapb.DataType_Bool, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := tt.dt
			v := tt.v
			builder := array.NewBuilder(memory.DefaultAllocator, serdeMap[dt].arrowType(1))
			defer builder.Release()
			serdeMap[dt].serialize(builder, v)
			a := builder.NewArray()

			// Test deserialize with shouldCopy parameter
			copy, got1 := serdeMap[dt].deserialize(a, 0, true)
			if !got1 {
				t.Errorf("deserialize() failed for %s", tt.name)
			}
			if !reflect.DeepEqual(copy, tt.v) {
				t.Errorf("deserialize() got = %v, want %v", copy, tt.v)
			}
			ref, _ := serdeMap[dt].deserialize(a, 0, false)
			// check the unsafe pointers of copy and ref are different
			switch v := copy.(type) {
			case []byte:
				if unsafe.Pointer(&v[0]) == unsafe.Pointer(&ref.([]byte)[0]) {
					t.Errorf("deserialize() got same pointer for %v", tt.v)
				}
			case string:
				if unsafe.StringData(v) == unsafe.StringData(ref.(string)) {
					t.Errorf("deserialize() got same pointer for %v", tt.v)
				}
			}

			a.Release()
		})
	}
}

func BenchmarkDeserializeReader(b *testing.B) {
	len := 1000000
	blobs, err := generateTestData(len)
	assert.NoError(b, err)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reader, err := NewBinlogDeserializeReader(generateTestSchema(), MakeBlobsReader(blobs), false)
		assert.NoError(b, err)
		defer reader.Close()
		for i := 0; i < len; i++ {
			_, err = reader.NextValue()
			assert.NoError(b, err)
		}
		_, err = reader.NextValue()
		assert.Equal(b, io.EOF, err)
	}
}

func TestCalculateArraySize(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.DefaultAllocator)
	defer mem.AssertSize(t, 0)

	tests := []struct {
		name         string
		arrayBuilder func() arrow.Array
		expectedSize uint64
	}{
		{
			name: "Empty array",
			arrayBuilder: func() arrow.Array {
				b := array.NewInt32Builder(mem)
				defer b.Release()
				return b.NewArray()
			},
			expectedSize: 0,
		},
		{
			name: "Fixed-length array",
			arrayBuilder: func() arrow.Array {
				b := array.NewInt32Builder(mem)
				defer b.Release()
				b.AppendValues([]int32{1, 2, 3, 4}, nil)
				return b.NewArray()
			},
			expectedSize: 20, // 4 elements * 4 bytes + bitmap(4bytes)
		},
		{
			name: "Variable-length string array",
			arrayBuilder: func() arrow.Array {
				b := array.NewStringBuilder(mem)
				defer b.Release()
				b.AppendValues([]string{"hello", "world"}, nil)
				return b.NewArray()
			},
			expectedSize: 23, // bytes: "hello" (5 bytes) + "world" (5 bytes)
			// offsets: 2+1 elements * 4 bytes
			// bitmap(1 byte)
		},
		{
			name: "Nested list array",
			arrayBuilder: func() arrow.Array {
				b := array.NewListBuilder(mem, arrow.PrimitiveTypes.Int32)
				defer b.Release()
				valueBuilder := b.ValueBuilder().(*array.Int32Builder)

				b.Append(true)
				valueBuilder.AppendValues([]int32{1, 2, 3}, nil)

				b.Append(true)
				valueBuilder.AppendValues([]int32{4, 5}, nil)

				b.Append(true)
				valueBuilder.AppendValues([]int32{}, nil)

				return b.NewArray()
			},
			expectedSize: 44, // child buffer: 5 elements * 4 bytes, plus bitmap (4bytes)
			// offsets: 3+1 elements * 4 bytes
			// bitmap(4 bytes)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := tt.arrayBuilder()
			defer arr.Release()

			size := arr.Data().SizeInBytes()
			if size != tt.expectedSize {
				t.Errorf("Expected size %d, got %d", tt.expectedSize, size)
			}
		})
	}
}
