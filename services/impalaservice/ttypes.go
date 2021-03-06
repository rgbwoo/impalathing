// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package impalaservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/koblas/impalathing/services/beeswax"
	"github.com/koblas/impalathing/services/cli_service"
	"github.com/koblas/impalathing/services/status"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__
var GoUnusedProtection__ int

type TImpalaQueryOptions int64

const (
	TImpalaQueryOptions_ABORT_ON_ERROR            TImpalaQueryOptions = 0
	TImpalaQueryOptions_MAX_ERRORS                TImpalaQueryOptions = 1
	TImpalaQueryOptions_DISABLE_CODEGEN           TImpalaQueryOptions = 2
	TImpalaQueryOptions_BATCH_SIZE                TImpalaQueryOptions = 3
	TImpalaQueryOptions_MEM_LIMIT                 TImpalaQueryOptions = 4
	TImpalaQueryOptions_NUM_NODES                 TImpalaQueryOptions = 5
	TImpalaQueryOptions_MAX_SCAN_RANGE_LENGTH     TImpalaQueryOptions = 6
	TImpalaQueryOptions_MAX_IO_BUFFERS            TImpalaQueryOptions = 7
	TImpalaQueryOptions_NUM_SCANNER_THREADS       TImpalaQueryOptions = 8
	TImpalaQueryOptions_ALLOW_UNSUPPORTED_FORMATS TImpalaQueryOptions = 9
	TImpalaQueryOptions_DEFAULT_ORDER_BY_LIMIT    TImpalaQueryOptions = 10
	TImpalaQueryOptions_DEBUG_ACTION              TImpalaQueryOptions = 11
)

func (p TImpalaQueryOptions) String() string {
	switch p {
	case TImpalaQueryOptions_ABORT_ON_ERROR:
		return "ABORT_ON_ERROR"
	case TImpalaQueryOptions_MAX_ERRORS:
		return "MAX_ERRORS"
	case TImpalaQueryOptions_DISABLE_CODEGEN:
		return "DISABLE_CODEGEN"
	case TImpalaQueryOptions_BATCH_SIZE:
		return "BATCH_SIZE"
	case TImpalaQueryOptions_MEM_LIMIT:
		return "MEM_LIMIT"
	case TImpalaQueryOptions_NUM_NODES:
		return "NUM_NODES"
	case TImpalaQueryOptions_MAX_SCAN_RANGE_LENGTH:
		return "MAX_SCAN_RANGE_LENGTH"
	case TImpalaQueryOptions_MAX_IO_BUFFERS:
		return "MAX_IO_BUFFERS"
	case TImpalaQueryOptions_NUM_SCANNER_THREADS:
		return "NUM_SCANNER_THREADS"
	case TImpalaQueryOptions_ALLOW_UNSUPPORTED_FORMATS:
		return "ALLOW_UNSUPPORTED_FORMATS"
	case TImpalaQueryOptions_DEFAULT_ORDER_BY_LIMIT:
		return "DEFAULT_ORDER_BY_LIMIT"
	case TImpalaQueryOptions_DEBUG_ACTION:
		return "DEBUG_ACTION"
	}
	return "<UNSET>"
}

func TImpalaQueryOptionsFromString(s string) (TImpalaQueryOptions, error) {
	switch s {
	case "ABORT_ON_ERROR":
		return TImpalaQueryOptions_ABORT_ON_ERROR, nil
	case "MAX_ERRORS":
		return TImpalaQueryOptions_MAX_ERRORS, nil
	case "DISABLE_CODEGEN":
		return TImpalaQueryOptions_DISABLE_CODEGEN, nil
	case "BATCH_SIZE":
		return TImpalaQueryOptions_BATCH_SIZE, nil
	case "MEM_LIMIT":
		return TImpalaQueryOptions_MEM_LIMIT, nil
	case "NUM_NODES":
		return TImpalaQueryOptions_NUM_NODES, nil
	case "MAX_SCAN_RANGE_LENGTH":
		return TImpalaQueryOptions_MAX_SCAN_RANGE_LENGTH, nil
	case "MAX_IO_BUFFERS":
		return TImpalaQueryOptions_MAX_IO_BUFFERS, nil
	case "NUM_SCANNER_THREADS":
		return TImpalaQueryOptions_NUM_SCANNER_THREADS, nil
	case "ALLOW_UNSUPPORTED_FORMATS":
		return TImpalaQueryOptions_ALLOW_UNSUPPORTED_FORMATS, nil
	case "DEFAULT_ORDER_BY_LIMIT":
		return TImpalaQueryOptions_DEFAULT_ORDER_BY_LIMIT, nil
	case "DEBUG_ACTION":
		return TImpalaQueryOptions_DEBUG_ACTION, nil
	}
	return TImpalaQueryOptions(0), fmt.Errorf("not a valid TImpalaQueryOptions string")
}

func TImpalaQueryOptionsPtr(v TImpalaQueryOptions) *TImpalaQueryOptions { return &v }

func (p TImpalaQueryOptions) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TImpalaQueryOptions) UnmarshalText(text []byte) error {
	q, err := TImpalaQueryOptionsFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - RowsAppended
type TInsertResult_ struct {
	RowsAppended map[string]int64 `thrift:"rows_appended,1,required" json:"rows_appended"`
}

func NewTInsertResult_() *TInsertResult_ {
	return &TInsertResult_{}
}

func (p *TInsertResult_) GetRowsAppended() map[string]int64 {
	return p.RowsAppended
}
func (p *TInsertResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRowsAppended bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetRowsAppended = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRowsAppended {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field RowsAppended is not set"))
	}
	return nil
}

func (p *TInsertResult_) readField1(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]int64, size)
	p.RowsAppended = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		var _val1 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val1 = v
		}
		p.RowsAppended[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *TInsertResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TInsertResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TInsertResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rows_appended", thrift.MAP, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:rows_appended: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.I64, len(p.RowsAppended)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.RowsAppended {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteI64(int64(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:rows_appended: ", p), err)
	}
	return err
}

func (p *TInsertResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TInsertResult_(%+v)", *p)
}
