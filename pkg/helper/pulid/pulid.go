package pulid

import (
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/oklog/ulid/v2"
)

// ID - Prefixed ULID. The first two characters indicate the type of entity being described.
type ID struct {
	Prefix string
	ulid.ULID
}

// EncodedSize is the length of a text encoded PULID.
const EncodedSize = ulid.EncodedSize + 2

// String returns the string encoding of the PULID.
func (id ID) String() string {
	return id.Prefix + id.ULID.String()
}

// Parse returns the string encoding of the PULID.
func Parse(pulid string) (id ID, err error) {
	id.Prefix = string(pulid[0:2])
	id.ULID, err = ulid.Parse(pulid[2:])
	if err != nil {
		return id, fmt.Errorf("issue parsing ulid: %w", err)
	}
	return id, nil
}

// MarshalText implements the encoding.TextMarshaler interface by
// returning the string encoded ULID.
func (id ID) MarshalText() ([]byte, error) {
	ulid := make([]byte, EncodedSize)
	copy(ulid, []byte(id.Prefix))
	return ulid, id.MarshalTextTo(ulid[2:])
}

// UnmarshalText implements the encoding.TextUnmarshaler interface by
func (id *ID) UnmarshalText(v []byte) error {
	id.Prefix = string(v[0:2])
	return id.ULID.UnmarshalText(v[2:])
}

// MarshalBinary implements the encoding.BinaryMarshaler interface by
// returning the ULID as a byte slice.
func (id ID) MarshalBinary() ([]byte, error) {
	ulid := make([]byte, EncodedSize)
	return ulid, id.MarshalBinaryTo(ulid)
}

// MarshalBinaryTo writes the binary encoding of the ULID to the given buffer.
// ErrBufferSize is returned when the len(dst) != 16.
func (id ID) MarshalBinaryTo(dst []byte) error {
	if len(dst) != EncodedSize {
		return ulid.ErrBufferSize
	}

	copy(dst, []byte(id.Prefix))
	copy(dst[2:], id.ULID[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface by
// copying the passed data and converting it to an ULID. ErrDataSize is
// returned if the data length is different from ULID length.
func (id *ID) UnmarshalBinary(data []byte) error {
	if len(data) != EncodedSize {
		return ulid.ErrDataSize
	}

	id.Prefix = string(data[0:2])
	copy((*id).ULID[:], data[2:])
	return nil
}

// Scan implements the sql.Scanner interface. It supports scanning
// a string or byte slice.
func (id *ID) Scan(src interface{}) error {
	switch x := src.(type) {
	case nil:
		return nil
	case string:
		return id.UnmarshalText([]byte(x))
	case []byte:
		return id.UnmarshalBinary(x)
	}

	return ulid.ErrScanValue
}

// Value implements the sql/driver.Valuer interface. This returns the value
// represented as a byte slice. If instead a string is desirable, a wrapper
// type can be created that calls String().
//
//	// stringValuer wraps a ULID as a string-based driver.Valuer.
// 	type stringValuer ULID
//
//	func (id stringValuer) Value() (driver.Value, error) {
//		return ULID(id).String(), nil
//	}
//
//	// Example usage.
//	db.Exec("...", stringValuer(id))
func (id ID) Value() (driver.Value, error) {
	return id.MarshalBinary()
}

// MarshalID providers a graphql marshaler for ULIDs.
func MarshalID(id ID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, `"`)
		io.WriteString(w, id.String())
		io.WriteString(w, `"`)
	})
}

// UnmarshalID unmarshals ULID.
func UnmarshalID(v interface{}) (id ID, err error) {
	str, ok := v.(string)
	if !ok {
		return id, fmt.Errorf("ids must be strings")
	}
	return Parse(str)
}
