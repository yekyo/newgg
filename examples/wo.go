package main

import "github.com/yekyo/newgg"

func main() {
	const w = 7736
	const h = 9236
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	if err := dc.LoadNewFontFace("/Users/pengye/Library/Fonts/941-CAI978_7.ttf", 288, 150); err != nil {
		panic(err)
	}
	//dc.SetRGB255(0, 0, 0)
	dc.SetRGB255(168,165,162)
	s := "Grandma"
	n := 48 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		/*go func(dy int) {
			for dx := -n; dx <= n; dx++ {
				if dx*dx+dy*dy >= n*n {
					// give it rounded corners
					continue
				}
				x := w/2 + float64(dx)
				y := h/2 + float64(dy)
				dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
			}
		}(dy)*/
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := w/2 + float64(dx)
			y := h/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB255(255, 255, 255)
	dc.DrawStringAnchored(s, w/2, h/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}

