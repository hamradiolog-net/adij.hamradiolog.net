package adif

import (
	"encoding/json"
	"io"

	"github.com/hamradiolog-net/spec/v6/adifield"
)

var _ ADIFRecordReader = (*adijReader)(nil)

// adijReader implements ADIFRecordReader for reading ADIF records in ADIJ format.
type adijReader struct {
	document     *adifDocument
	currentIndex int
	skipHeader   bool
}

// NewADIJReader returns an ADIFRecordReader that can parse ADIF records in ADIJ JSON format.
// If skipHeader is true, Next() will not return the header record if it exists.
// This implementation reads the entire JSON document into memory when this constructor is called.
func NewADIJReader(r io.Reader, skipHeader bool) (*adijReader, error) {
	var doc adifDocument
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&doc)
	if err != nil {
		return nil, err
	}

	reader := &adijReader{
		document:     &doc,
		currentIndex: -1,
		skipHeader:   skipHeader,
	}

	return reader, nil
}

// Next reads and returns the next Record.
// It returns io.EOF when no more records are available.
func (r *adijReader) Next() (ADIFRecord, error) {
	if r.currentIndex >= len(r.document.Records) {
		return nil, io.EOF
	}
	if r.currentIndex == -1 {
		r.currentIndex = 0
		if !r.skipHeader && r.document.Header != nil {
			header := r.mapToADIRecord(r.document.Header)
			header.SetIsHeader(true)
			return header, nil
		}
		if len(r.document.Records) == 0 {
			return nil, io.EOF
		}
	}

	record := r.mapToADIRecord(r.document.Records[r.currentIndex])
	r.currentIndex++
	return record, nil
}

// mapToADIRecord converts a map of ADIField to string into an ADIFRecord
func (r *adijReader) mapToADIRecord(fieldMap map[adifield.ADIField]string) ADIFRecord {
	record := NewADIRecord()
	for field, value := range fieldMap {
		record.Set(field, value)
	}
	return record
}
