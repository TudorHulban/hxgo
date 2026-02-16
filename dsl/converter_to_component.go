package dsl

import "unsafe"

// TODO: move to map for performance.
func dataForTag(tag string) unsafe.Pointer {
	switch tag {
	case "div":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'd', 'i', 'v'},
			closeTag: []byte{'<', '/', 'd', 'i', 'v', '>'},
		})
	case "span":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 's', 'p', 'a', 'n'},
			closeTag: []byte{'<', '/', 's', 'p', 'a', 'n', '>'},
		})
	case "a":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'a'},
			closeTag: []byte{'<', '/', 'a', '>'},
		})
	case "p":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'p'},
			closeTag: []byte{'<', '/', 'p', '>'},
		})
	case "img":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'i', 'm', 'g'},
			closeTag: nil,
		})
	case "ul":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'u', 'l'},
			closeTag: []byte{'<', '/', 'u', 'l', '>'},
		})
	case "ol":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'o', 'l'},
			closeTag: []byte{'<', '/', 'o', 'l', '>'},
		})
	case "li":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'l', 'i'},
			closeTag: []byte{'<', '/', 'l', 'i', '>'},
		})
	case "nav":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'n', 'a', 'v'},
			closeTag: []byte{'<', '/', 'n', 'a', 'v', '>'},
		})
	case "h1":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'h', '1'},
			closeTag: []byte{'<', '/', 'h', '1', '>'},
		})
	case "h2":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'h', '2'},
			closeTag: []byte{'<', '/', 'h', '2', '>'},
		})
	case "h3":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'h', '3'},
			closeTag: []byte{'<', '/', 'h', '3', '>'},
		})
	case "h4":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'h', '4'},
			closeTag: []byte{'<', '/', 'h', '4', '>'},
		})
	case "h5":
		return unsafe.Pointer(&struct {
			openTag  []byte
			closeTag []byte
		}{
			openTag:  []byte{'<', 'h', '5'},
			closeTag: []byte{'<', '/', 'h', '5', '>'},
		})
	case "h6":
		return unsafe.Pointer(
			&struct {
				openTag  []byte
				closeTag []byte
			}{
				openTag:  []byte{'<', 'h', '6'},
				closeTag: []byte{'<', '/', 'h', '6', '>'},
			},
		)
	default:
		return buildDataForTag(tag)
	}
}
