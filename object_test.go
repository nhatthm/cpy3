package cpy3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy3"
)

func TestAttrString(t *testing.T) {
	cpy3.Py_Initialize()

	sys := cpy3.PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttrString("stdout"))

	stdout := sys.GetAttrString("stdout")

	assert.NotNil(t, stdout)
	assert.Zero(t, sys.DelAttrString("stdout"))
	assert.Nil(t, sys.GetAttrString("stdout"))

	cpy3.PyErr_Clear()

	assert.Zero(t, sys.SetAttrString("stdout", stdout))
}

func TestAttr(t *testing.T) {
	cpy3.Py_Initialize()

	name := cpy3.PyUnicode_FromString("stdout")
	defer name.DecRef()

	sys := cpy3.PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttr(name))

	stdout := sys.GetAttr(name)

	assert.NotNil(t, stdout)
	assert.Zero(t, sys.DelAttr(name))
	assert.Nil(t, sys.GetAttr(name))

	cpy3.PyErr_Clear()

	assert.Zero(t, sys.SetAttr(name, stdout))
}

func TestRichCompareBool(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := cpy3.PyUnicode_FromString("test1")
	s2 := cpy3.PyUnicode_FromString("test2")

	assert.Zero(t, s1.RichCompareBool(s2, cpy3.Py_EQ))
	assert.NotZero(t, s1.RichCompareBool(s1, cpy3.Py_EQ))
}

func TestRichCompare(t *testing.T) {
	cpy3.Py_Initialize()

	s1 := cpy3.PyUnicode_FromString("test1")
	s2 := cpy3.PyUnicode_FromString("test2")

	b1 := s1.RichCompare(s2, cpy3.Py_EQ)
	defer b1.DecRef()

	assert.Equal(t, cpy3.Py_False, b1)

	b2 := s1.RichCompare(s1, cpy3.Py_EQ)
	defer b2.DecRef()

	assert.Equal(t, cpy3.Py_True, b2)
}

func TestRepr(t *testing.T) {
	cpy3.Py_Initialize()

	list := cpy3.PyList_New(0)
	defer list.DecRef()

	repr := list.Repr()

	assert.Equal(t, "[]", cpy3.PyUnicode_AsUTF8(repr))
}

func TestStr(t *testing.T) {
	cpy3.Py_Initialize()

	list := cpy3.PyList_New(0)
	defer list.DecRef()

	str := list.Str()

	assert.Equal(t, "[]", cpy3.PyUnicode_AsUTF8(str))
}

func TestASCII(t *testing.T) {
	cpy3.Py_Initialize()

	list := cpy3.PyList_New(0)
	defer list.DecRef()

	ascii := list.ASCII()

	assert.Equal(t, "[]", cpy3.PyUnicode_AsUTF8(ascii))
}

func TestCallable(t *testing.T) {
	cpy3.Py_Initialize()

	builtins := cpy3.PyEval_GetBuiltins()

	assert.True(t, cpy3.PyDict_Check(builtins))

	len := cpy3.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy3.PyCallable_Check(len))

	emptyList := cpy3.PyList_New(0)

	assert.True(t, cpy3.PyList_Check(emptyList))

	args := cpy3.PyTuple_New(1)
	defer args.DecRef()

	assert.True(t, cpy3.PyTuple_Check(args))

	cpy3.PyTuple_SetItem(args, 0, emptyList)

	length := len.Call(args, nil)
	defer length.DecRef()

	assert.True(t, cpy3.PyLong_Check(length))
	assert.Equal(t, 0, cpy3.PyLong_AsLong(length))

	length = len.CallObject(args)
	defer length.DecRef()

	assert.True(t, cpy3.PyLong_Check(length))
	assert.Equal(t, 0, cpy3.PyLong_AsLong(length))

	length = len.CallFunctionObjArgs(emptyList)
	defer length.DecRef()

	assert.True(t, cpy3.PyLong_Check(length))
	assert.Equal(t, 0, cpy3.PyLong_AsLong(length))
}

func TestCallMethod(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("hello world")
	defer s.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(s))

	sep := cpy3.PyUnicode_FromString(" ")
	defer sep.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(sep))

	split := cpy3.PyUnicode_FromString("split")
	defer split.DecRef()

	assert.True(t, cpy3.PyUnicode_Check(split))

	words := s.CallMethodObjArgs(split, sep)
	defer words.DecRef()

	assert.True(t, cpy3.PyList_Check(words))
	assert.Equal(t, 2, cpy3.PyList_Size(words))

	hello := cpy3.PyList_GetItem(words, 0)
	world := cpy3.PyList_GetItem(words, 1)

	assert.True(t, cpy3.PyUnicode_Check(hello))
	assert.True(t, cpy3.PyUnicode_Check(world))
	assert.Equal(t, "hello", cpy3.PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", cpy3.PyUnicode_AsUTF8(world))

	words = s.CallMethodArgs("split", sep)
	defer words.DecRef()

	assert.True(t, cpy3.PyList_Check(words))
	assert.Equal(t, 2, cpy3.PyList_Size(words))

	hello = cpy3.PyList_GetItem(words, 0)
	world = cpy3.PyList_GetItem(words, 1)

	assert.True(t, cpy3.PyUnicode_Check(hello))
	assert.True(t, cpy3.PyUnicode_Check(world))
	assert.Equal(t, "hello", cpy3.PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", cpy3.PyUnicode_AsUTF8(world))
}

func TestIsTrue(t *testing.T) {
	cpy3.Py_Initialize()

	b := cpy3.Py_True.IsTrue() != 0

	assert.True(t, b)

	b = cpy3.Py_False.IsTrue() != 0

	assert.False(t, b)
}

func TestNot(t *testing.T) {
	cpy3.Py_Initialize()

	b := cpy3.Py_True.Not() != 0

	assert.False(t, b)

	b = cpy3.Py_False.Not() != 0

	assert.True(t, b)
}

func TestLength(t *testing.T) {
	cpy3.Py_Initialize()
	length := 6

	list := cpy3.PyList_New(length)
	defer list.DecRef()

	listLength := list.Length()

	assert.Equal(t, length, listLength)
}

func TestLengthHint(t *testing.T) {
	cpy3.Py_Initialize()

	length := 6

	list := cpy3.PyList_New(length)
	defer list.DecRef()

	listLength := list.LengthHint(0)

	assert.Equal(t, length, listLength)
}

func TestObjectItem(t *testing.T) {
	cpy3.Py_Initialize()

	key := cpy3.PyUnicode_FromString("key")
	defer key.DecRef()

	value := cpy3.PyUnicode_FromString("value")
	defer value.DecRef()

	dict := cpy3.PyDict_New()
	err := dict.SetItem(key, value)

	assert.Zero(t, err)

	dictValue := dict.GetItem(key)

	assert.Equal(t, value, dictValue)

	err = dict.DelItem(key)

	assert.Zero(t, err)
}

func TestDir(t *testing.T) {
	cpy3.Py_Initialize()

	list := cpy3.PyList_New(0)
	defer list.DecRef()

	dir := list.Dir()
	defer dir.DecRef()

	repr := dir.Repr()
	defer repr.DecRef()

	assert.Equal(t, "['__add__', '__class__', '__class_getitem__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getstate__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']", cpy3.PyUnicode_AsUTF8(repr))
}

func TestReprEnterLeave(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("hello world")
	defer s.DecRef()

	assert.Zero(t, s.ReprEnter())
	assert.True(t, s.ReprEnter() > 0)

	s.ReprLeave()
	s.ReprLeave()
}

func TestIsSubclass(t *testing.T) {
	cpy3.Py_Initialize()

	assert.Equal(t, 1, cpy3.PyExc_Warning.IsSubclass(cpy3.PyExc_Exception))
	assert.Equal(t, 0, cpy3.Bool.IsSubclass(cpy3.Float))
}

func TestHash(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.NotEqual(t, -1, s.Hash())
}

func TestObjectType(t *testing.T) {
	cpy3.Py_Initialize()

	i := cpy3.PyLong_FromGoInt(23543)
	defer i.DecRef()

	assert.Equal(t, cpy3.Long, i.Type())
}

func TestHashNotImplemented(t *testing.T) {
	cpy3.Py_Initialize()

	s := cpy3.PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.Equal(t, -1, s.HashNotImplemented())
	assert.True(t, cpy3.PyErr_ExceptionMatches(cpy3.PyExc_TypeError))

	cpy3.PyErr_Clear()
}

func TestObjectIter(t *testing.T) {
	cpy3.Py_Initialize()

	i := cpy3.PyLong_FromGoInt(23)
	defer i.DecRef()

	assert.Nil(t, i.GetIter())
	assert.True(t, cpy3.PyErr_ExceptionMatches(cpy3.PyExc_TypeError))

	cpy3.PyErr_Clear()

	list := cpy3.PyList_New(23)
	defer list.DecRef()

	iter := list.GetIter()
	defer iter.DecRef()

	assert.NotNil(t, iter)
}
