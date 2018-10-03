package main

// https://tools.ietf.org/html/rfc4346#section-6.2.1
type protocolVersion struct {
	major, minor uint8
}

/*
 The TLS Record Layer which handles all data transport.
 The record layer is assumed to sit directly on top of some
 reliable transport such as TCP. The record layer can carry four kinds of payloads:

 1. Handshake messages—used for algorithm negotiation and key establishment.
 2. ChangeCipherSpec messages—really part of the handshake but technically a separate kind of message.
 3. Alert messages—used to signal that errors have occurred
 4. Application layer data

 The DTLS record layer is extremely similar to that of TLS 1.1.  The
 only change is the inclusion of an explicit sequence number in the
 record.  This sequence number allows the recipient to correctly
 verify the TLS MAC.
 https://tools.ietf.org/html/rfc4347#section-4.1
*/
type recordLayer struct {
	protocolVersion protocolVersion
	epoch           uint16
	sequenceNumber  uint64 // uint48 in spec
	content         content
}

func (r *recordLayer) Marshal() ([]byte, error) {
	return nil, nil
}

func (r *recordLayer) Unmarshal(data []byte) error {
	return nil
}
