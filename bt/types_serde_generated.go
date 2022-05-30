
// Code generated by serde-go; DO NOT EDIT.
package bt

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

var _ = errors.New

type serdeStructEnum_StoredConfirmedBlock = int

const (
	serdeStructEnumSkip_StoredConfirmedBlock serdeStructEnum_StoredConfirmedBlock = 0
	
	serdeStructEnum_StoredConfirmedBlock_PreviousBlockhash  serdeStructEnum_StoredConfirmedBlock = 1
	
	
	serdeStructEnum_StoredConfirmedBlock_Blockhash  serdeStructEnum_StoredConfirmedBlock = 2
	
	
	serdeStructEnum_StoredConfirmedBlock_ParentSlot  serdeStructEnum_StoredConfirmedBlock = 3
	
)


type serdeStructFieldVisitor_StoredConfirmedBlock struct {
	e *serdeStructEnum_StoredConfirmedBlock

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_StoredConfirmedBlock(e *serdeStructEnum_StoredConfirmedBlock) serdeStructFieldVisitor_StoredConfirmedBlock {
	return serdeStructFieldVisitor_StoredConfirmedBlock{
		e: e,
		DummyVisitor: serde.NewDummyVisitor("StoredConfirmedBlock Field"),
	}
}

func (s serdeStructFieldVisitor_StoredConfirmedBlock) VisitString(v string) (err error) {
	switch v {
	
	case "PreviousBlockhash":
		*s.e = serdeStructEnum_StoredConfirmedBlock_PreviousBlockhash
	
	
	case "Blockhash":
		*s.e = serdeStructEnum_StoredConfirmedBlock_Blockhash
	
	
	case "ParentSlot":
		*s.e = serdeStructEnum_StoredConfirmedBlock_ParentSlot
	
	default:
		return errors.New("invalid field")
	}
	return nil
}

type serdeStructVisitor_StoredConfirmedBlock struct {
	v *StoredConfirmedBlock

	serde.DummyVisitor
}

func serdeNewStructVisitor_StoredConfirmedBlock(v *StoredConfirmedBlock) serdeStructVisitor_StoredConfirmedBlock {
	return serdeStructVisitor_StoredConfirmedBlock{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("StoredConfirmedBlock"),
	}
}

func (s serdeStructVisitor_StoredConfirmedBlock) VisitNil() (err error) {
	return nil
}

func (s serdeStructVisitor_StoredConfirmedBlock) VisitMap(m serde.MapAccess) (err error) {
	var fieldValue serdeStructEnum_StoredConfirmedBlock
	field := serdeNewStructFieldVisitor_StoredConfirmedBlock(&fieldValue)
	for {
		ok, err := m.NextKey(field)
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		var v serde.Visitor
		switch *field.e {
		case serdeStructEnumSkip_StoredConfirmedBlock:
			v = serde.SkipVisitor{}
	
		case serdeStructEnum_StoredConfirmedBlock_PreviousBlockhash:
			v = serde.NewStringVisitor(&s.v.PreviousBlockhash)
	
	
		case serdeStructEnum_StoredConfirmedBlock_Blockhash:
			v = serde.NewStringVisitor(&s.v.Blockhash)
	
	
		case serdeStructEnum_StoredConfirmedBlock_ParentSlot:
			v = serde.NewUint64Visitor(&s.v.ParentSlot)
	
		default:
			return errors.New("invalid field")
		}
		err = m.NextValue(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StoredConfirmedBlock) Deserialize(de serde.Deserializer) (err error) {
	return de.DeserializeStruct("StoredConfirmedBlock", nil, serdeNewStructVisitor_StoredConfirmedBlock(s))
}



func (s StoredConfirmedBlock) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("StoredConfirmedBlock", 3)
	if err != nil {
		return err
	}
	
	err = st.SerializeField(
		serde.StringSerializer("PreviousBlockhash"),
		serde.StringSerializer(s.PreviousBlockhash),
	)
	if err != nil {
		return
	}
	
	
	err = st.SerializeField(
		serde.StringSerializer("Blockhash"),
		serde.StringSerializer(s.Blockhash),
	)
	if err != nil {
		return
	}
	
	
	err = st.SerializeField(
		serde.StringSerializer("ParentSlot"),
		serde.Uint64Serializer(s.ParentSlot),
	)
	if err != nil {
		return
	}
	
	err = st.EndStruct()
	if err != nil {
		return
	}
	return nil
}

