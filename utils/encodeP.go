package utils

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/robertkrimen/otto"
)

var (
	script     string
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	input, err := ReadFromFile(path.Join(basepath, "..", "converters", "rawdeflate.js"), "")
	if err != nil {
		panic(err)
	}
	script = input
	script += `
	function encode64(data) {
		r = "";
		for (i=0; i<data.length; i+=3) {
			 if (i+2==data.length) {
				r += append3bytes(data.charCodeAt(i), data.charCodeAt(i+1), 0);
			} else if (i+1==data.length) {
				r += append3bytes(data.charCodeAt(i), 0, 0);
			} else {
				r += append3bytes(data.charCodeAt(i), data.charCodeAt(i+1),
					data.charCodeAt(i+2));
			}
			console.log("r:\n" + r);
		}
		return r;
	}
	
	function append3bytes(b1, b2, b3) {
		c1 = b1 >> 2;
		c2 = ((b1 & 0x3) << 4) | (b2 >> 4);
		c3 = ((b2 & 0xF) << 2) | (b3 >> 6);
		c4 = b3 & 0x3F;
		r = "";
		r += encode6bit(c1 & 0x3F);
		r += encode6bit(c2 & 0x3F);
		r += encode6bit(c3 & 0x3F);
		r += encode6bit(c4 & 0x3F);
		return r;
	}
	
	function encode6bit(b) {
		if (b < 10) {
			 return String.fromCharCode(48 + b);
		}
		b -= 10;
		if (b < 26) {
			 return String.fromCharCode(65 + b);
		}
		b -= 26;
		if (b < 26) {
			 return String.fromCharCode(97 + b);
		}
		b -= 26;
		if (b == 0) {
			 return '-';
		}
		if (b == 1) {
			 return '_';
		}
		return '?';
	}
	(function(){
			console.log("encoding\n" + input + "\n")
			s = unescape(encodeURIComponent(input));
			return encode64(zip_deflate(s, 9))
		})();
	`
}

// EncodeP - returns the base64 similar ecnoded string of plantuml.
// see http://plantuml.com/de/code-javascript-synchronous
func EncodeP(input string) string {
	vm := otto.New()
	vm.Set("input", input)
	result, err := vm.Run(script)
	if err != nil {
		log.Printf("Error %s", err.Error())
		return ""
	}
	value, err := result.ToString()
	if err != nil {
		log.Printf("Error %s", err.Error())
		return ""
	}
	return value
}
