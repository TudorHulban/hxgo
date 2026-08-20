package dsl

// Using as Constructor Function:
//
//	func NewPageLayout(title, description, lang string, head ...dsl.Node) Layout {
//		return func(nodes ...dsl.Node) dsl.Node {
//			page := Page{
//				Title:       title,
//				Description: description,
//				Language:    lang,
//				Head:        head,
//				Body:        nodes, // Passes the layout children directly into Body
//			}
//			return page.Build()
//		}
//	}
type Layout func(nodes ...Node) Node
