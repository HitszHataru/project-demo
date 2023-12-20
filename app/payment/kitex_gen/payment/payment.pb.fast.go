// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package payment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreditCardInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreditCardInfo[number], err)
}

func (x *CreditCardInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.CreditCardNumber, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreditCardInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CreditCardCvv, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CreditCardInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CreditCardExpirationYear, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *CreditCardInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CreditCardExpirationMonth, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ChargeRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ChargeRequest[number], err)
}

func (x *ChargeRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *ChargeRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v CreditCardInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CreditCard = &v
	return offset, nil
}

func (x *ChargeRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ChargeResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ChargeResponse[number], err)
}

func (x *ChargeResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TransactionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreditCardInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *CreditCardInfo) fastWriteField1(buf []byte) (offset int) {
	if x.CreditCardNumber == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCreditCardNumber())
	return offset
}

func (x *CreditCardInfo) fastWriteField2(buf []byte) (offset int) {
	if x.CreditCardCvv == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetCreditCardCvv())
	return offset
}

func (x *CreditCardInfo) fastWriteField3(buf []byte) (offset int) {
	if x.CreditCardExpirationYear == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetCreditCardExpirationYear())
	return offset
}

func (x *CreditCardInfo) fastWriteField4(buf []byte) (offset int) {
	if x.CreditCardExpirationMonth == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetCreditCardExpirationMonth())
	return offset
}

func (x *ChargeRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ChargeRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 1, x.GetAmount())
	return offset
}

func (x *ChargeRequest) fastWriteField2(buf []byte) (offset int) {
	if x.CreditCard == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetCreditCard())
	return offset
}

func (x *ChargeRequest) fastWriteField3(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOrderId())
	return offset
}

func (x *ChargeResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ChargeResponse) fastWriteField1(buf []byte) (offset int) {
	if x.TransactionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTransactionId())
	return offset
}

func (x *CreditCardInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *CreditCardInfo) sizeField1() (n int) {
	if x.CreditCardNumber == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCreditCardNumber())
	return n
}

func (x *CreditCardInfo) sizeField2() (n int) {
	if x.CreditCardCvv == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetCreditCardCvv())
	return n
}

func (x *CreditCardInfo) sizeField3() (n int) {
	if x.CreditCardExpirationYear == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetCreditCardExpirationYear())
	return n
}

func (x *CreditCardInfo) sizeField4() (n int) {
	if x.CreditCardExpirationMonth == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetCreditCardExpirationMonth())
	return n
}

func (x *ChargeRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ChargeRequest) sizeField1() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeFloat(1, x.GetAmount())
	return n
}

func (x *ChargeRequest) sizeField2() (n int) {
	if x.CreditCard == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetCreditCard())
	return n
}

func (x *ChargeRequest) sizeField3() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetOrderId())
	return n
}

func (x *ChargeResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ChargeResponse) sizeField1() (n int) {
	if x.TransactionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTransactionId())
	return n
}

var fieldIDToName_CreditCardInfo = map[int32]string{
	1: "CreditCardNumber",
	2: "CreditCardCvv",
	3: "CreditCardExpirationYear",
	4: "CreditCardExpirationMonth",
}

var fieldIDToName_ChargeRequest = map[int32]string{
	1: "Amount",
	2: "CreditCard",
	3: "OrderId",
}

var fieldIDToName_ChargeResponse = map[int32]string{
	1: "TransactionId",
}
