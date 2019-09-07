package utils

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"strings"
)

// see http://plantuml.com/de/code-php

/**
<?php

function encodep($text) {
	$data = utf8_encode($text);
	$compressed = gzdeflate($data, 9);
	return encode64($compressed);
}

function encode6bit($b) {
	if ($b < 10) {
		return chr(48 + $b);
	}
	$b -= 10;
	if ($b < 26) {
		return chr(65 + $b);
	}
	$b -= 26;
	if ($b < 26) {
		return chr(97 + $b);
	}
	$b -= 26;
	if ($b == 0) {
		return '-';
	}
	if ($b == 1) {
		return '_';
	}
	return '?';
}

function append3bytes($b1, $b2, $b3) {
	$c1 = $b1 >> 2;
	$c2 = (($b1 & 0x3) << 4) | ($b2 >> 4);
	$c3 = (($b2 & 0xF) << 2) | ($b3 >> 6);
	$c4 = $b3 & 0x3F;
	$r = "";
	$r .= encode6bit($c1 & 0x3F);
	$r .= encode6bit($c2 & 0x3F);
	$r .= encode6bit($c3 & 0x3F);
	$r .= encode6bit($c4 & 0x3F);
	return $r;
}

function encode64($c) {
	$str = "";
	$len = strlen($c);
	for ($i = 0; $i < $len; $i+=3) {
		if ($i+2==$len) {
			$str .= append3bytes(ord(substr($c, $i, 1)), ord(substr($c, $i+1, 1)), 0);
		} else if ($i+1==$len) {
			$str .= append3bytes(ord(substr($c, $i, 1)), 0, 0);
		} else {
			$str .= append3bytes(
					ord(substr($c, $i, 1)),
					ord(substr($c, $i+1, 1)),
					ord(substr($c, $i+2, 1)));
		}
	}
	return $str;
}
?>
**/

func encode6bit(b int8) rune {
	if b < 10 {
		return rune(48 + b)
	}
	b -= 10
	if b < 26 {
		return rune(65 + b)
	}
	b -= 26
	if b < 26 {
		return rune(97 + b)
	}
	b -= 26
	if b == 0 {
		return '-'
	}
	if b == 1 {
		return '_'
	}
	return '?'
}

func append3bytes(b1 int8, b2 int8, b3 int8) string {
	c1 := b1 >> 2
	c2 := ((b1 & 0x3) << 4) | (b2 >> 4)
	c3 := ((b1 & 0xF) << 2) | (b2 >> 6)
	c4 := b3 & 0x3F
	return string([]rune{
		encode6bit(c1 & 0x3F),
		encode6bit(c2 & 0x3F),
		encode6bit(c3 & 0x3F),
		encode6bit(c4 & 0x3F)})
}

func encode64(c string) string {
	str := ""
	len := len(c)
	fmt.Printf("str-len: %d\n", len)
	for i := 0; i < len; i += 3 {
		if i+2 == len {
			str += append3bytes(
				ord(c, i, i+1),
				ord(c, i+1, i+2),
				0)
		} else if i+1 == len {
			str += append3bytes(
				ord(c, i, i+1),
				0,
				0)
		} else {
			str += append3bytes(
				ord(c, i, i+1),
				ord(c, i+1, i+2),
				ord(c, i+2, i+3))
		}
	}
	return str
}

func ord(c string, s int, e int) int8 {
	rs := []rune(c[s:e])
	fmt.Printf("len %d, encoding %c\n", len(rs), rs[0])
	return int8(rs[0])
}

//EncodeP - returns the base64 similar ecnoded string of plantuml.
func EncodeP(input string) string {
	var b bytes.Buffer

	zw, err := flate.NewWriter(&b, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(zw, strings.NewReader(input)); err != nil {
		log.Fatal(err)
	}
	if err := zw.Flush(); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buf-len: %d\n", b.Len())

	return encode64(b.String())
}
