package main

func Chapter5() {

	canvas := NewCanvas(101, 101)

	rayOrigin := PointV(0, 0, -5)
	s := NewSphere(0, 0, 0, 1)
	s.Transform = CT(ScaleM(0.5, 1, 1), TranslateM(2, 0, 0))
	wallWidth := 10.
	wallHeight := 10.

	w, h := 0., 0. // wall position
	wc, hc := 0, 0 // pixel counter
	for w < wallWidth {
		h = 0.
		hc = 0
		for h < wallHeight {
			dir := SubV(PointV(w-5, h-5, 10), rayOrigin)

			// log.Println(dir)
			r := NewRay(rayOrigin, dir)
			xs := IntersectRaySphere(r, s)
			hit := xs.Hit()
			if len(hit) == 0 {
				canvas.Write(NewColour(0, 0, 0), wc, hc)
			} else {
				canvas.Write(NewColour(1, 0, 0), wc, hc)
			}
			h = h + wallHeight/100.
			hc = hc + 1
		}

		w = w + wallWidth/100.
		wc = wc + 1
	}

	canvas.WritePPM("chapter5.ppm")
}
