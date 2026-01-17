package dsl

import "unsafe"

func renderGenericElement(a *accumulator, data unsafe.Pointer) {
	tag := *(*string)(data)

	// <tag
	a.html = append(a.html, '<')
	a.html = append(a.html, tag...)

	// attributes are rendered by children with isAttribute=true
	// so do nothing here

	// >
	a.html = append(a.html, '>')

	// children are rendered by walkHTML

	// </tag>
	a.html = append(a.html, '<', '/')
	a.html = append(a.html, tag...)
	a.html = append(a.html, '>')
}

// core renderers

func renderDiv(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'd', 'i', 'v') // avoids string allocations
	// attributes rendered by children
	a.html = append(a.html, '>')
	// children rendered by walkHTML
	a.html = append(a.html, '<', '/', 'd', 'i', 'v', '>')
}

func renderSpan(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 's', 'p', 'a', 'n')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 's', 'p', 'a', 'n', '>')
}

func renderA(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'a')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'a', '>')
}

func renderP(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'p')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'p', '>')
}

func renderImg(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'i', 'm', 'g')
	// attributes rendered by children
	a.html = append(a.html, '/', '>')
}

func renderUl(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'u', 'l')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'u', 'l', '>')
}

func renderOl(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'o', 'l')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'o', 'l', '>')
}

func renderLi(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'l', 'i')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'l', 'i', '>')
}

func renderNav(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'n', 'a', 'v')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'n', 'a', 'v', '>')
}

// heading renderers

func renderH1(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '1')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '1', '>')
}

func renderH2(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '2')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '2', '>')
}

func renderH3(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '3')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '3', '>')
}

func renderH4(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '4')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '4', '>')
}

func renderH5(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '5')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '5', '>')
}

func renderH6(a *accumulator, data unsafe.Pointer) {
	a.html = append(a.html, '<', 'h', '6')
	a.html = append(a.html, '>')
	a.html = append(a.html, '<', '/', 'h', '6', '>')
}
