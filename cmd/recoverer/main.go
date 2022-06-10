package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
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

	corrected, err := process(fitFile, *start, *size)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(newFileName(fitFile), corrected.Bytes(), 0644)
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

func process(fitFile string, start int, size int) (*bytes.Buffer, error) {
	// read the file
	raw, err := ioutil.ReadFile(fitFile)
	if err != nil {
		return nil, err
	}

	// read only the header
	header, err := fit.DecodeHeader(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}

	// reconstruct the data portions
	beginning := raw[header.Size:start]
	end := raw[start+size : len(raw)-2]
	combined := append(beginning, end...)

	// calculate the CRC
	crc := dyncrc16.New()
	header.DataSize = uint32(len(combined))
	hdr, err := header.MarshalBinary()
	if err != nil {
		return nil, err
	}

	_, err = crc.Write(hdr)
	if err != nil {
		return nil, err
	}
	_, err = crc.Write(combined)
	if err != nil {
		return nil, err
	}

	// write the file
	buf := &bytes.Buffer{}
	_, err = buf.Write(hdr)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(combined)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, crc.Sum16())
	return buf, err
}
