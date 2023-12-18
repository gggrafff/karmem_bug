# How to use
1) Generate serialization code
```
cd ./serialization_api
make
cd ../
```
2) Run example
```
go run .
```

# What is problem
Struct before serialization does not match struct after serialization. My output of this example:  
```
Test of karmem
{{33 42 true}} was encoded to 40 bytes
40 bytes was decoded to {{0 0 false}}
panic: 42 == msg1.A.IntField != msg2.A.IntField == 0
```

# What is solution
~~I think that the following part of generated code is invalid:~~  
```
func (x *BtypeViewer) A() (v *AtypeViewer) {
	return (*AtypeViewer)(unsafe.Add(unsafe.Pointer(&x._data), 0))
}
```
~~We need use the correct offset instead of 0~~

~~The following code fixes the problem:~~  
```
func (x *BtypeViewer) A() (v *AtypeViewer) {
	return (*AtypeViewer)(unsafe.Add(unsafe.Pointer(&x._data), x.size()))
}
```
**Error in writer (see next paragraph)**

We allocate memory for the nested structure twice, because we fall into the condition `if offset == 0`  
We need to avoid double allocation and just write data with valid offset.

# Temporary patch
It can help:
```
struct Btype inline {
    FakeField int32;
    A Atype;
}
```
And this gives a smaller encoded message than the original example. Thus, we can conclude that the **error in the writer**