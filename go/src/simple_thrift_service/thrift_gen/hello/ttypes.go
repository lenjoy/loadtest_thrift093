// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hello

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Message
//  - InputID
//  - Dimension
type HelloRequest struct {
	Message   *string `thrift:"message,1" json:"message,omitempty"`
	InputID   *int32  `thrift:"input_id,2" json:"input_id,omitempty"`
	Dimension *int16  `thrift:"dimension,3" json:"dimension,omitempty"`
}

func NewHelloRequest() *HelloRequest {
	return &HelloRequest{}
}

var HelloRequest_Message_DEFAULT string

func (p *HelloRequest) GetMessage() string {
	if !p.IsSetMessage() {
		return HelloRequest_Message_DEFAULT
	}
	return *p.Message
}

var HelloRequest_InputID_DEFAULT int32

func (p *HelloRequest) GetInputID() int32 {
	if !p.IsSetInputID() {
		return HelloRequest_InputID_DEFAULT
	}
	return *p.InputID
}

var HelloRequest_Dimension_DEFAULT int16

func (p *HelloRequest) GetDimension() int16 {
	if !p.IsSetDimension() {
		return HelloRequest_Dimension_DEFAULT
	}
	return *p.Dimension
}
func (p *HelloRequest) IsSetMessage() bool {
	return p.Message != nil
}

func (p *HelloRequest) IsSetInputID() bool {
	return p.InputID != nil
}

func (p *HelloRequest) IsSetDimension() bool {
	return p.Dimension != nil
}

func (p *HelloRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *HelloRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *HelloRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.InputID = &v
	}
	return nil
}

func (p *HelloRequest) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Dimension = &v
	}
	return nil
}

func (p *HelloRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *HelloRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *HelloRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetInputID() {
		if err := oprot.WriteFieldBegin("input_id", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:input_id: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.InputID)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.input_id (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:input_id: ", p), err)
		}
	}
	return err
}

func (p *HelloRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetDimension() {
		if err := oprot.WriteFieldBegin("dimension", thrift.I16, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:dimension: ", p), err)
		}
		if err := oprot.WriteI16(int16(*p.Dimension)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.dimension (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:dimension: ", p), err)
		}
	}
	return err
}

func (p *HelloRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloRequest(%+v)", *p)
}

// Attributes:
//  - Message
//  - Results
type HelloResponse struct {
	Message *string     `thrift:"message,1" json:"message,omitempty"`
	Results []*HelloDoc `thrift:"results,2" json:"results,omitempty"`
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{}
}

var HelloResponse_Message_DEFAULT string

func (p *HelloResponse) GetMessage() string {
	if !p.IsSetMessage() {
		return HelloResponse_Message_DEFAULT
	}
	return *p.Message
}

var HelloResponse_Results_DEFAULT []*HelloDoc

func (p *HelloResponse) GetResults() []*HelloDoc {
	return p.Results
}
func (p *HelloResponse) IsSetMessage() bool {
	return p.Message != nil
}

func (p *HelloResponse) IsSetResults() bool {
	return p.Results != nil
}

func (p *HelloResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *HelloResponse) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *HelloResponse) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*HelloDoc, 0, size)
	p.Results = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &HelloDoc{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Results = append(p.Results, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *HelloResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *HelloResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
		}
	}
	return err
}

func (p *HelloResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetResults() {
		if err := oprot.WriteFieldBegin("results", thrift.LIST, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:results: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Results)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Results {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:results: ", p), err)
		}
	}
	return err
}

func (p *HelloResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloResponse(%+v)", *p)
}

// Attributes:
//  - DocID
//  - Vec
//  - Score
type HelloDoc struct {
	DocID int32     `thrift:"doc_id,1" json:"doc_id"`
	Vec   []float64 `thrift:"vec,2" json:"vec"`
	Score *float64  `thrift:"score,3" json:"score,omitempty"`
}

func NewHelloDoc() *HelloDoc {
	return &HelloDoc{}
}

func (p *HelloDoc) GetDocID() int32 {
	return p.DocID
}

func (p *HelloDoc) GetVec() []float64 {
	return p.Vec
}

var HelloDoc_Score_DEFAULT float64

func (p *HelloDoc) GetScore() float64 {
	if !p.IsSetScore() {
		return HelloDoc_Score_DEFAULT
	}
	return *p.Score
}
func (p *HelloDoc) IsSetScore() bool {
	return p.Score != nil
}

func (p *HelloDoc) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *HelloDoc) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DocID = v
	}
	return nil
}

func (p *HelloDoc) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]float64, 0, size)
	p.Vec = tSlice
	for i := 0; i < size; i++ {
		var _elem1 float64
		if v, err := iprot.ReadDouble(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Vec = append(p.Vec, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *HelloDoc) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Score = &v
	}
	return nil
}

func (p *HelloDoc) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HelloDoc"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *HelloDoc) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("doc_id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:doc_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.DocID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.doc_id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:doc_id: ", p), err)
	}
	return err
}

func (p *HelloDoc) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vec", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:vec: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.DOUBLE, len(p.Vec)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Vec {
		if err := oprot.WriteDouble(float64(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:vec: ", p), err)
	}
	return err
}

func (p *HelloDoc) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetScore() {
		if err := oprot.WriteFieldBegin("score", thrift.DOUBLE, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:score: ", p), err)
		}
		if err := oprot.WriteDouble(float64(*p.Score)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.score (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:score: ", p), err)
		}
	}
	return err
}

func (p *HelloDoc) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloDoc(%+v)", *p)
}
