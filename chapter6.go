package main

func Chapter6() {
	canvas := NewCanvas(101, 101)
	eyePos := PointV(0, 0, -5)

	wallWidth := 10. // image target width & height
	wallHeight := 10.

	s := NewSphere()
	m := NewMaterial()
	m.Colour = NewColour(1, 0.2, 1)
	s.Material = m

	light := NewPointLight(PointV(-10, 10, -10), NewColour(1, 1, 1))

	w, h := 0., 0. // wall position
	wc, hc := 0, 0 // pixel counter
	for w < wallWidth {
		h = 0.
		hc = 0
		for h < wallHeight {
			dir := SubV(PointV(w-5, h-5, 10), eyePos)

			// log.Println(dir)
			r := NewRay(eyePos, dir)
			xs := s.Intersect(r)
			hit := xs.Hit()
			if len(hit) == 0 {
				canvas.Write(NewColour(0, 0, 0), wc, hc)
			} else {
				hitPos := r.Position(hit[0].T)
				normal := s.NormalAt(hitPos)
				toEye := ScaleV(-1, r.Dir)

				col := Lighting(m, nil, light, hitPos, toEye, normal, false)
				canvas.Write(col, wc, hc)
			}
			h = h + wallHeight/100.
			hc = hc + 1
		}

		w = w + wallWidth/100.
		wc = wc + 1
	}

	canvas.WritePPM("chapter6.ppm")
}
