package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tormoder/fit"
	"github.com/tormoder/fit/dyncrc16"
)

func main() {
	start := flag.Int(
		"start",
		-1,
		"start of the corruption",
	)
	size := flag.Int(
		"size",
		-1,
		"number of bytes to strip",
	)
	strip := flag.Bool(
		"strip",
		false,
		"strip zeroes from end of file",
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: recoverer [flags] [path to fit file]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("flags %d %d\n", *start, *size)
		flag.Usage()
		os.Exit(2)
	}

	fitFile := flag.Arg(0)
	fmt.Printf("flags %d %d %s\n", *start, *size, fitFile)

	var corrected []byte
	var err error

	if *strip {
		corrected, err = forceRepair(fitFile)
	} else {
		corrected, err = trimMiddle(fitFile, *start, *size)
	}
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(newFileName(fitFile), corrected, 0644)
	if err != nil {
		panic(err)
	}
}

func newFileName(fileName string) string {
	d := filepath.Dir(fileName)
	e := filepath.Ext(fileName)
	n := strings.TrimSuffix(filepath.Base(fileName), e)
	return filepath.Join(d, n+"-corrected"+e)
}

func readFile(fitFile string) (fit.Header, []byte, error) {
	raw, err := os.ReadFile(fitFile)
	if err != nil {
		return fit.Header{}, nil, err
	}

	// read only the header
	header, err := fit.DecodeHeader(bytes.NewReader(raw))
	if err != nil {
		return header, nil, err
	}

	data := raw[header.Size:]
	return header, data, nil
}

func reconstructFile(header fit.Header, data []byte) ([]byte, error) {
	// process crc
	crc := dyncrc16.New()
	header.DataSize = uint32(len(data))
	hdr, err := header.MarshalBinary()
	if err != nil {
		return nil, err
	}

	_, err = crc.Write(hdr)
	if err != nil {
		return nil, err
	}
	_, err = crc.Write(data)
	if err != nil {
		return nil, err
	}

	// write the file
	buf := &bytes.Buffer{}
	_, err = buf.Write(hdr)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(data)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, crc.Sum16())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func forceRepair(fitFile string) ([]byte, error) {
	header, data, err := readFile(fitFile)
	if err != nil {
		return nil, err
	}

	// strip trailing zeroes
	for data[len(data)-1] == 0x00 {
		data = data[:len(data)-1]
	}

	// try 10 times
	for i := 0; i < 10; i++ {
		reconstructed, err := reconstructFile(header, data)
		if err != nil {
			// should not fail
			return nil, err
		}

		// attempt to read
		_, err = fit.Decode(bytes.NewReader(reconstructed))
		if err != nil {
			// remove last byte and try again
			fmt.Println("fit file invalid, stripping more")
			data = data[:len(data)-1]
			continue
		}
		return reconstructed, nil
	}

	return nil, fmt.Errorf("failed to create a valid fit file")
}

func trimMiddle(fitFile string, start int, size int) ([]byte, error) {
	header, data, err := readFile(fitFile)
	if err != nil {
		return nil, err
	}

	// reconstruct the data portions
	beginning := data[header.Size:start]
	end := data[start+size : len(data)-2]
	combined := append(beginning, end...)

	return reconstructFile(header, combined)
}
