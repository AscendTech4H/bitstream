package bitstream

import "io"

//Reader reads from an io.Reader bit by bit
type Reader struct {
	src io.Reader
	dat byte
	bit uint8
}

//Bit returns one bit
func (s *Reader) Bit() (uint8, error) {
	if s.bit > 7 {
		s.bit = 0
		buf := make([]byte, 1)
		_, err := s.src.Read(buf)
		if err != nil {
			return 0, err
		}
		s.dat = buf[0]
	}
	o := s.dat % 2
	s.dat = s.dat >> 1
	s.bit--
	return o, nil
}

//Bits returns a certain number of bits
func (s *Reader) Bits(n uint) (uint, error) {
	o := uint(0)
	for ; n > 0; n-- {
		b, err := s.Bit()
		if err != nil {
			return 0, err
		}
		o += uint(b)
		o = o << 1
	}
	return o, nil
}

//NewReader creates a new bitstream.Reader
func NewReader(r io.Reader) *Reader {
	reader := new(Reader)
	reader.bit = 0
	reader.dat = 0
	reader.src = r
	return reader
}
