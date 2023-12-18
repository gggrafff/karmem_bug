package serialization_api

import (
	"unsafe"

	karmem "karmem.org/golang"
)

var _ unsafe.Pointer

var _Null = make([]byte, 24)
var _NullReader = karmem.NewReader(_Null)

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierAtype = 5766179297190848870
	PacketIdentifierBtype = 2764376575326507992
)

type Atype struct {
	UintField uint32
	IntField  int32
	BoolField bool
}

func NewAtype() Atype {
	return Atype{}
}

func (x *Atype) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierAtype
}

func (x *Atype) Reset() {
	x.Read((*AtypeViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Atype) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Atype) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(16)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__UintFieldOffset := offset + 0
	writer.Write4At(__UintFieldOffset, *(*uint32)(unsafe.Pointer(&x.UintField)))
	__IntFieldOffset := offset + 4
	writer.Write4At(__IntFieldOffset, *(*uint32)(unsafe.Pointer(&x.IntField)))
	__BoolFieldOffset := offset + 8
	writer.Write1At(__BoolFieldOffset, *(*uint8)(unsafe.Pointer(&x.BoolField)))

	return offset, nil
}

func (x *Atype) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewAtypeViewer(reader, 0), reader)
}

func (x *Atype) Read(viewer *AtypeViewer, reader *karmem.Reader) {
	x.UintField = viewer.UintField()
	x.IntField = viewer.IntField()
	x.BoolField = viewer.BoolField()
}

type Btype struct {
	A Atype
}

func NewBtype() Btype {
	return Btype{}
}

func (x *Btype) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierBtype
}

func (x *Btype) Reset() {
	x.Read((*BtypeViewer)(unsafe.Pointer(&_Null)), _NullReader)
}

func (x *Btype) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *Btype) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(24)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__AOffset := offset + 0
	if _, err := x.A.Write(writer, __AOffset); err != nil {
		return offset, err
	}

	return offset, nil
}

func (x *Btype) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewBtypeViewer(reader, 0), reader)
}

func (x *Btype) Read(viewer *BtypeViewer, reader *karmem.Reader) {
	x.A.Read(viewer.A(), reader)
}

type AtypeViewer struct {
	_data [16]byte
}

func NewAtypeViewer(reader *karmem.Reader, offset uint32) (v *AtypeViewer) {
	if !reader.IsValidOffset(offset, 16) {
		return (*AtypeViewer)(unsafe.Pointer(&_Null))
	}
	v = (*AtypeViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *AtypeViewer) size() uint32 {
	return 16
}
func (x *AtypeViewer) UintField() (v uint32) {
	return *(*uint32)(unsafe.Add(unsafe.Pointer(&x._data), 0))
}
func (x *AtypeViewer) IntField() (v int32) {
	return *(*int32)(unsafe.Add(unsafe.Pointer(&x._data), 4))
}
func (x *AtypeViewer) BoolField() (v bool) {
	return *(*bool)(unsafe.Add(unsafe.Pointer(&x._data), 8))
}

type BtypeViewer struct {
	_data [24]byte
}

func NewBtypeViewer(reader *karmem.Reader, offset uint32) (v *BtypeViewer) {
	if !reader.IsValidOffset(offset, 24) {
		return (*BtypeViewer)(unsafe.Pointer(&_Null))
	}
	v = (*BtypeViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *BtypeViewer) size() uint32 {
	return 24
}
func (x *BtypeViewer) A() (v *AtypeViewer) {
	return (*AtypeViewer)(unsafe.Add(unsafe.Pointer(&x._data), 0)) // 0 -> x.size()
}
