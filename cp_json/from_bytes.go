package cp_json

import (
	"github.com/v2pro/plz/util"
	"github.com/json-iterator/go"
	"unsafe"
	"github.com/v2pro/plz/lang"
	"reflect"
	"github.com/v2pro/plz/lang/tagging"
)

var byteArrayType = reflect.TypeOf([]byte{})

func provideFromBytesCopier(dstType, srcType reflect.Type) (util.ObjectCopier, error) {
	isFromBytes := srcType == byteArrayType && dstType.Kind() == reflect.Ptr && tagging.Get(dstType.Elem()).Tags["codec"].Text() == "json"
	if !isFromBytes {
		return nil, nil
	}
	srcAcc := lang.AccessorOf(reflect.TypeOf((*jsoniter.Iterator)(nil)), "json")
	dstAcc := lang.AccessorOf(dstType, "json")
	copier, err := util.CopierOf(dstAcc, srcAcc)
	if err != nil {
		return nil, err
	}
	return &fromBytesCopier{copier}, nil
}

type fromBytesCopier struct {
	copier util.Copier
}

func (objCopier *fromBytesCopier) Copy(dst, src interface{}) error {
	bytes := src.([]byte)
	iter := jsoniter.ParseBytes(jsoniter.ConfigDefault, bytes)
	err := objCopier.copier.Copy(lang.AddressOf(dst), unsafe.Pointer(iter))
	if err != nil {
		return err
	}
	return nil
}
