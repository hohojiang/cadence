// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.13.1. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// WorkflowService_RecordActivityTaskHeartbeatByID_Args represents the arguments for the WorkflowService.RecordActivityTaskHeartbeatByID function.
//
// The arguments for RecordActivityTaskHeartbeatByID are sent and received over the wire as this struct.
type WorkflowService_RecordActivityTaskHeartbeatByID_Args struct {
	HeartbeatRequest *shared.RecordActivityTaskHeartbeatByIDRequest `json:"heartbeatRequest,omitempty"`
}

// ToWire translates a WorkflowService_RecordActivityTaskHeartbeatByID_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.HeartbeatRequest != nil {
		w, err = v.HeartbeatRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RecordActivityTaskHeartbeatByIDRequest_Read(w wire.Value) (*shared.RecordActivityTaskHeartbeatByIDRequest, error) {
	var v shared.RecordActivityTaskHeartbeatByIDRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_RecordActivityTaskHeartbeatByID_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RecordActivityTaskHeartbeatByID_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RecordActivityTaskHeartbeatByID_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.HeartbeatRequest, err = _RecordActivityTaskHeartbeatByIDRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RecordActivityTaskHeartbeatByID_Args
// struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.HeartbeatRequest != nil {
		fields[i] = fmt.Sprintf("HeartbeatRequest: %v", v.HeartbeatRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_RecordActivityTaskHeartbeatByID_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RecordActivityTaskHeartbeatByID_Args match the
// provided WorkflowService_RecordActivityTaskHeartbeatByID_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) Equals(rhs *WorkflowService_RecordActivityTaskHeartbeatByID_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.HeartbeatRequest == nil && rhs.HeartbeatRequest == nil) || (v.HeartbeatRequest != nil && rhs.HeartbeatRequest != nil && v.HeartbeatRequest.Equals(rhs.HeartbeatRequest))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WorkflowService_RecordActivityTaskHeartbeatByID_Args.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.HeartbeatRequest != nil {
		err = multierr.Append(err, enc.AddObject("heartbeatRequest", v.HeartbeatRequest))
	}
	return err
}

// GetHeartbeatRequest returns the value of HeartbeatRequest if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) GetHeartbeatRequest() (o *shared.RecordActivityTaskHeartbeatByIDRequest) {
	if v.HeartbeatRequest != nil {
		return v.HeartbeatRequest
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RecordActivityTaskHeartbeatByID" for this struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) MethodName() string {
	return "RecordActivityTaskHeartbeatByID"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_RecordActivityTaskHeartbeatByID_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.RecordActivityTaskHeartbeatByID
// function.
var WorkflowService_RecordActivityTaskHeartbeatByID_Helper = struct {
	// Args accepts the parameters of RecordActivityTaskHeartbeatByID in-order and returns
	// the arguments struct for the function.
	Args func(
		heartbeatRequest *shared.RecordActivityTaskHeartbeatByIDRequest,
	) *WorkflowService_RecordActivityTaskHeartbeatByID_Args

	// IsException returns true if the given error can be thrown
	// by RecordActivityTaskHeartbeatByID.
	//
	// An error can be thrown by RecordActivityTaskHeartbeatByID only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RecordActivityTaskHeartbeatByID
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// RecordActivityTaskHeartbeatByID into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by RecordActivityTaskHeartbeatByID
	//
	//   value, err := RecordActivityTaskHeartbeatByID(args)
	//   result, err := WorkflowService_RecordActivityTaskHeartbeatByID_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RecordActivityTaskHeartbeatByID: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.RecordActivityTaskHeartbeatResponse, error) (*WorkflowService_RecordActivityTaskHeartbeatByID_Result, error)

	// UnwrapResponse takes the result struct for RecordActivityTaskHeartbeatByID
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if RecordActivityTaskHeartbeatByID threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_RecordActivityTaskHeartbeatByID_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_RecordActivityTaskHeartbeatByID_Result) (*shared.RecordActivityTaskHeartbeatResponse, error)
}{}

func init() {
	WorkflowService_RecordActivityTaskHeartbeatByID_Helper.Args = func(
		heartbeatRequest *shared.RecordActivityTaskHeartbeatByIDRequest,
	) *WorkflowService_RecordActivityTaskHeartbeatByID_Args {
		return &WorkflowService_RecordActivityTaskHeartbeatByID_Args{
			HeartbeatRequest: heartbeatRequest,
		}
	}

	WorkflowService_RecordActivityTaskHeartbeatByID_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.DomainNotActiveError:
			return true
		case *shared.LimitExceededError:
			return true
		case *shared.ServiceBusyError:
			return true
		default:
			return false
		}
	}

	WorkflowService_RecordActivityTaskHeartbeatByID_Helper.WrapResponse = func(success *shared.RecordActivityTaskHeartbeatResponse, err error) (*WorkflowService_RecordActivityTaskHeartbeatByID_Result, error) {
		if err == nil {
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.BadRequestError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.InternalServiceError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.EntityNotExistError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{EntityNotExistError: e}, nil
		case *shared.DomainNotActiveError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.DomainNotActiveError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{DomainNotActiveError: e}, nil
		case *shared.LimitExceededError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.LimitExceededError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{LimitExceededError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RecordActivityTaskHeartbeatByID_Result.ServiceBusyError")
			}
			return &WorkflowService_RecordActivityTaskHeartbeatByID_Result{ServiceBusyError: e}, nil
		}

		return nil, err
	}
	WorkflowService_RecordActivityTaskHeartbeatByID_Helper.UnwrapResponse = func(result *WorkflowService_RecordActivityTaskHeartbeatByID_Result) (success *shared.RecordActivityTaskHeartbeatResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.DomainNotActiveError != nil {
			err = result.DomainNotActiveError
			return
		}
		if result.LimitExceededError != nil {
			err = result.LimitExceededError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_RecordActivityTaskHeartbeatByID_Result represents the result of a WorkflowService.RecordActivityTaskHeartbeatByID function call.
//
// The result of a RecordActivityTaskHeartbeatByID execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_RecordActivityTaskHeartbeatByID_Result struct {
	// Value returned by RecordActivityTaskHeartbeatByID after a successful execution.
	Success              *shared.RecordActivityTaskHeartbeatResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError                     `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError                `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError                `json:"entityNotExistError,omitempty"`
	DomainNotActiveError *shared.DomainNotActiveError                `json:"domainNotActiveError,omitempty"`
	LimitExceededError   *shared.LimitExceededError                  `json:"limitExceededError,omitempty"`
	ServiceBusyError     *shared.ServiceBusyError                    `json:"serviceBusyError,omitempty"`
}

// ToWire translates a WorkflowService_RecordActivityTaskHeartbeatByID_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) ToWire() (wire.Value, error) {
	var (
		fields [7]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.DomainNotActiveError != nil {
		w, err = v.DomainNotActiveError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.LimitExceededError != nil {
		w, err = v.LimitExceededError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_RecordActivityTaskHeartbeatByID_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WorkflowService_RecordActivityTaskHeartbeatByID_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RecordActivityTaskHeartbeatByID_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RecordActivityTaskHeartbeatByID_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _RecordActivityTaskHeartbeatResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.DomainNotActiveError, err = _DomainNotActiveError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.LimitExceededError, err = _LimitExceededError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.DomainNotActiveError != nil {
		count++
	}
	if v.LimitExceededError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_RecordActivityTaskHeartbeatByID_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RecordActivityTaskHeartbeatByID_Result
// struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [7]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.DomainNotActiveError != nil {
		fields[i] = fmt.Sprintf("DomainNotActiveError: %v", v.DomainNotActiveError)
		i++
	}
	if v.LimitExceededError != nil {
		fields[i] = fmt.Sprintf("LimitExceededError: %v", v.LimitExceededError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}

	return fmt.Sprintf("WorkflowService_RecordActivityTaskHeartbeatByID_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RecordActivityTaskHeartbeatByID_Result match the
// provided WorkflowService_RecordActivityTaskHeartbeatByID_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) Equals(rhs *WorkflowService_RecordActivityTaskHeartbeatByID_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.DomainNotActiveError == nil && rhs.DomainNotActiveError == nil) || (v.DomainNotActiveError != nil && rhs.DomainNotActiveError != nil && v.DomainNotActiveError.Equals(rhs.DomainNotActiveError))) {
		return false
	}
	if !((v.LimitExceededError == nil && rhs.LimitExceededError == nil) || (v.LimitExceededError != nil && rhs.LimitExceededError != nil && v.LimitExceededError.Equals(rhs.LimitExceededError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of WorkflowService_RecordActivityTaskHeartbeatByID_Result.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BadRequestError != nil {
		err = multierr.Append(err, enc.AddObject("badRequestError", v.BadRequestError))
	}
	if v.InternalServiceError != nil {
		err = multierr.Append(err, enc.AddObject("internalServiceError", v.InternalServiceError))
	}
	if v.EntityNotExistError != nil {
		err = multierr.Append(err, enc.AddObject("entityNotExistError", v.EntityNotExistError))
	}
	if v.DomainNotActiveError != nil {
		err = multierr.Append(err, enc.AddObject("domainNotActiveError", v.DomainNotActiveError))
	}
	if v.LimitExceededError != nil {
		err = multierr.Append(err, enc.AddObject("limitExceededError", v.LimitExceededError))
	}
	if v.ServiceBusyError != nil {
		err = multierr.Append(err, enc.AddObject("serviceBusyError", v.ServiceBusyError))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetSuccess() (o *shared.RecordActivityTaskHeartbeatResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetBadRequestError returns the value of BadRequestError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetBadRequestError() (o *shared.BadRequestError) {
	if v.BadRequestError != nil {
		return v.BadRequestError
	}

	return
}

// GetInternalServiceError returns the value of InternalServiceError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetInternalServiceError() (o *shared.InternalServiceError) {
	if v.InternalServiceError != nil {
		return v.InternalServiceError
	}

	return
}

// GetEntityNotExistError returns the value of EntityNotExistError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetEntityNotExistError() (o *shared.EntityNotExistsError) {
	if v.EntityNotExistError != nil {
		return v.EntityNotExistError
	}

	return
}

// GetDomainNotActiveError returns the value of DomainNotActiveError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetDomainNotActiveError() (o *shared.DomainNotActiveError) {
	if v.DomainNotActiveError != nil {
		return v.DomainNotActiveError
	}

	return
}

// GetLimitExceededError returns the value of LimitExceededError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetLimitExceededError() (o *shared.LimitExceededError) {
	if v.LimitExceededError != nil {
		return v.LimitExceededError
	}

	return
}

// GetServiceBusyError returns the value of ServiceBusyError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) GetServiceBusyError() (o *shared.ServiceBusyError) {
	if v.ServiceBusyError != nil {
		return v.ServiceBusyError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "RecordActivityTaskHeartbeatByID" for this struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) MethodName() string {
	return "RecordActivityTaskHeartbeatByID"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_RecordActivityTaskHeartbeatByID_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
