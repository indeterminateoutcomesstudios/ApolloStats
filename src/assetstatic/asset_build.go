// +build embed

// Automagically generated by yaber v0.2 (https://github.com/lmas/yaber),
// please avoid editting this file as it might be regenerated again.

package assetstatic

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"
)

func Asset(path string) ([]byte, error) {
	body, ok := _rawAssets[path]
	if !ok {
		return nil, os.ErrNotExist
	}
	return decompress(body)
}

func AssetDir(dir string) (map[string][]byte, error) {
	var e error
	files := make(map[string][]byte)
	for path, body := range _rawAssets {
		if strings.HasPrefix(path, dir) {
			files[path], e = decompress(body)
			if e != nil {
				return nil, e
			}
		}
	}
	return files, nil
}

func decompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	gr, e := gzip.NewReader(buf)
	if e != nil {
		return nil, e
	}
	defer gr.Close()
	return ioutil.ReadAll(gr)
}

var _rawAssets = map[string][]byte{

	"static/css/style.css": []byte("\x1f\x8b\b\x00\x00\tn\x88\x00\xff\xa4W\xddr\xa36\x14\xbe\xf7S0\xb3\u04db\x9d\x85\xd8`\xbc[\xe7*\xcd\xc6\x0f\xd0i\x1f@ \x19\xd4\b\xe4\x11r\x12\xb7\xd3w\uf444@\b\x11\xd2I.v\xcd\xd1\xf9\xd39\xdf\xf9\xd1\xe6\xeek\xf4\x93\x9c\u0455\u0248\xa1\x1b\xbf\xca(\xfa\xfa\x89\xbf\xbbM-\x1b\xf6-*8\xbeE\xffl\"\xf8+9\xe3\xe2\x18}yxx\xb8\u05c4\x02\x95\u03d5\xe0\xd7\x16\xc7\xf6l\xf7\x98f\xe9\xd3\xfd\xe6\xdf\xcd&A\x85\xf2\xc2\xc8J\xf2&c\xc4h\xd5\x1e\xa3\x92\xb4\x92\b\xa3\xa2Ao\xf1+\u0172>F\xf9v{y\xb3TQQ\xe0\xdc\x01%BW\xc9\r\xf9\xcc[\x19\xbf\x12Z\xd5\xf2\b\x8e1\xec\x90;\xfa7\x01\x81dO\x1am\xfd\xcaz\u02ccvp*o\x8c\xc4\xf2v\x01\x9e\x96\xb7\xc4\b^\x10\u01b4\xad\x8e\xd1\u0330&h-\t#g\xd9QLzug\xc6\x11XWT\xcb!\x94Cs\x16M\xd6<\x8c&\x12\x15\v\n\xbe\\PEb\x1d\xe6\xa4E/\x05\x12=\xa3\xf1\x05N\xa4\u4349\x85q\xb2\xee#\x90e\x96\x12HD\x9a\xa7\xbfe\x99ID\xafV\xfd\x1f\x83#\x9d\xf9\xc5h\xfb\u071b\x9ak\x84C\x12\xcf\xc9n\u0122]j\xc9n\x06\x06\"\xa6\xdd\x05\x90x\x8ch\xab\xb5\x15\x8c\x97\xcf\xf7.\x92h[\x13A\u5493\xf0\xef\xb1\xe6/Dx\xf8;\x9dN\x8b\xd7~\xcc\xf2\x1f\xe9\xdeh<s\x81\x8a\xaa\x17~'DZ7#H(P\xc9\u06bb\xe9~\x8a\x8d!\x1f\xfb\x1e\"\x90\u071a <x9\x1aR&\x1e\xf7\xbb\xfc\xe4i\xccCa\x1b\x88\xc3-w\xa7\xc3\xe9q\x8e\x1f{~\xc8\u007f\xcd~\x1e\xf49\xfa\x16!\x13)\xf5\xe3\x85vT\x12\xbc\xc0>\x14#&%\xc4GR\xde\u0692P\x9a\x82\x01w\x83J\x84\xe0\u07b9 X\x9fA\a\xfa\x03\x15\x8ct\x9fj<N\a\x92\")\xb8\x80\xd0v6\xb6\xfa+\x96\xfc\x02\x11\x03\x04v\x9cQ\fQ~\xca\xf2\xccb\u00b0\fU3r\xed~@s\xdaiO%NT6\x86 \xb9\xa9\xe9\v\xcdp5\x90p\u075elM\xdao\u041cY>\xb8\xf7\xef*\xdf\xd0\u007f%\xa2\xec\xb3\u05ff\xdb$\x90\xbf&\x9e\xa0\xaa\a\x9f\xb9\xf9vV\x8e\xa9r<\u07c6\x80e:b\xb8\x04\n\x86lI\xaa\xb4Ckz=F5\x85\u0234&\u06ee#\xdd\x05\xb5\xbd7\xef\x94v\x1f\x9d4\xb7\u0459\xe8H\xba\x061\x86u\x9cl?\x1c}\xdd&\a\xeb\xabga\xae\t0?\xfd\x9e w\x19\xe1\x13\xa1K(\xbaIn\x9d\xf0*~k\x14\xe8\xcb\xc6j\xacL\xa4gS\u0128\xcb\u01ce\xba8\x87&\xe9J\xcd\x00s\xcd0\x1a\x80i\ue6f3n\xe6\xcbSa5\u07ceQL_\u0799\xdd\x1e\xf3:6<\x01\x95>\xf7\xf3\xa3\xd9sd\x12j\xcbwX ,\xe4\\\xaeb\xf0\xcc\x16n\x9a\x87\xf8$m\x88\xc7x\b*T\xdeM\x19\xd3m\x90\x91\xf1\xb6\x9a2\x1e\xc6\x15\xc7\r\xaa\xd9\x05\x14\x15b i\x89\x98=i 7,p\xf5y\x19\xf5\x16\xbe\x87z\xc06\xf9\xdeoES%\u05f2$]\xb7\xa4\xa0\xc7L%H\x00\x1d\xc9\x19L\x13\xbc\"k'\xc34!\xe2*\xbdHg\xbe\xe0\x8d0\xc0fH\x16\xd1vEv\xc9a*\u05ac\x06\xdd\xe5o\xb7\x8a\xac\xd9,\xd85\x90\xa5\xbfx\xf1\xf1\x86\x19\u0190PU\x80u\xbfZ\xd8\x1f\x03\xecv<\xbb;\x8d\xcfY\xa1\x86\xc4\r\xc7>\xe6\xc7\"\xea$\x92\xdd\xe2\xf6\xf4n+\x99\xaf;\xa3\u00a4\xa8\xfc\xb2\xfd\xc5[\xbe\xfb\x95\u007f\x94\xf9\xc82?o\xa2\xb3\xde<*\xfc_\xb3l\xec\x030\xe9\xff\x94\x9f\x9f\xf0\u03a4\xef Ee\x1d\u00ea\u06ac\xbe\x94\xc67\x91\xb9\x8a~`,\xbd9\xf4\xbbe\t2F\xefZ{\x87\a\x96;\xa3\xa7\x0f/\xc5\x01p B\x05m\xa9o\x0f\f\x9a\xff?\x00\x00\x00\xff\xff\x01\x00\x00\xff\xff\xed\x0e\xb8{\xa3\x0e\x00\x00"),

	"static/favicon.ico": []byte("\x1f\x8b\b\x00\x00\tn\x88\x00\xff\x00\x00\x00\xff\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00"),

	"static/robots.txt": []byte("\x1f\x8b\b\x00\x00\tn\x88\x00\xff\xe2\n-N-\xd2MLO\xcd+\xb1R\xd0\xe2r\xc9,N\xcc\xc9\xc9/\xb7R\xd0/.I,\xc9L\xd6G\x16JLN\xce/\xcd+\x89\xcf,I\xcd-F\x96HJ\xccC\xe1\x17\x01\x95\xa5`\x8a\xe8s\x01\x00\x00\x00\xff\xff\x01\x00\x00\xff\xff0\x84+<o\x00\x00\x00"),
}
