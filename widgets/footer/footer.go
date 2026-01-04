package wfooter

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/TudorHulban/hx-widgets/base"
)

func WidgetFooter() hxprimitives.Node {
	form := hxcomponents.NewFormThreeContainers(
		&hxcomponents.ParamsNewFormThreeContainers{
			IDDivEnclosing:  "footer-info",
			IDForm:          "form-footer",
			IDDivContainers: "form-footer-info",

			IDContainerLeft:   "contact-hours",
			IDContainerMiddle: "contact-address",
			IDContainerRight:  "footer-links",

			ElementsInputLeft: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Contact and Open Hours",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"+1234567890",
					),
				),
			},

			ElementsInputMiddle: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Address",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"Lorem ipsum 3",
					),
				),
			},

			ElementsInputRight: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Book an appointment",
					),
				),
			},
		},
	)

	return hxhtml.Div(
		hxprimitives.AttrClass(
			"page-footer",
		),

		form,

		hxhtml.Div(
			hxprimitives.AttrID(
				"copyright",
			),

			hxhtml.H4(
				hxprimitives.Text(
					"Barbershop ©2025",
				),
			),
		),
	)
}

func CSSWidgetFooter() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		DesktopFirst: true,

		CSSAllMedias: `
		.page-footer {
            background-color: #333;
            color: #fff;
            padding: 1.5rem 1.5rem 0.5rem 1.5rem;
            width: 100%;
            font-family: Arial, sans-serif;
        }

        #form-footer-info {
            display: flex;
            justify-content: space-between;
            max-width: 1200px;
            margin: 0 auto;
            flex-wrap: wrap;
            gap: 2rem;
        }

        #form-footer-info > div {
            flex: 1;
            min-width: 200px;
            display: flex;
            flex-direction: column;
            gap: 0.5rem;
        }

        #form-footer-info span {
            display: block;
        }

        #form-footer-info span:first-child {
            font-weight: bold;
            margin-bottom: 0.5rem;
        }

        #footer-links span {
            cursor: pointer;
            color: #4CAF50;
            transition: color 0.3s;
        }

        #footer-links span:hover {
            color: #45a049;
        }

        #copyright {
            text-align: center;
            margin-top: 1rem;
			padding-top: 1rem;
			padding-bottom: 1rem;
            border-top: 1px solid rgba(255, 255, 255, 0.1);
        }

		#copyright h4 {
    		margin-block-start: 0;
    		margin-block-end: 0;
    		margin-inline-start: 0;
    		margin-inline-end: 0;
		}
`,

		CSSResponsive: []pagecss.CSSMedia{
			{
				InflexionPointPX: base.Tablet,
				CSS: `
			.page-footer {
                padding: 1rem;
            }

			#form-footer-info {
                flex-direction: column;
                align-items: center;
                text-align: center;
				gap: 1.5rem;
            }

            #form-footer-info > div {
                min-width: 100%;
            }
				`,
			},
		},
	}
}
