package stats

import (
	"fmt"
    "time"
    "sync"
	"image/color"
    "net/http"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

// Query speed statistics
type Stats struct {
    data []time.Duration
    sync.RWMutex // read-write mutex
}

// Decorate Handler which stats need to be recorded
func (s *Stats) Decorate(inner http.HandlerFunc) http.HandlerFunc {
	var handler http.Handler = inner
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        handler.ServeHTTP(w, r)

        elapsed := time.Since(start)

		s.Lock()
		s.data = append(s.data, elapsed)
		if len(s.data) > 1000 {
			s.data = s.data[len(s.data)-1000:]
		}
		s.Unlock()
    })
}

// Handlers to display statistics
func (s *Stats) Statz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", `
			<h1>Latency stats</h1>
			<img src="/statz/scatter.png?rand=0" style="float:right; width:50%">
			<img src="/statz/hist.png?rand=0" style="float:left; width:50%">
			<script>
			setInterval(function() {
				var imgs = document.getElementsByTagName("IMG");
				for (var i=0; i < imgs.length; i++) {
					var eqPos = imgs[i].src.lastIndexOf("=");
					var src = imgs[i].src.substr(0, eqPos+1);
					imgs[i].src = src + Math.random();
				}
			}, 500);
			</script>
		`)
	// Math.random() is to prevent caching.
}

func (s *Stats) Scatter(w http.ResponseWriter, r *http.Request) {
	s.RLock() // Read lock stats struct
    defer s.RUnlock()

	xys := make(plotter.XYs, len(s.data))
    for i, d := range s.data {
		xys[i].X = float64(i)
		xys[i].Y = float64(d) / float64(time.Millisecond)
	}
	sc, err := plotter.NewScatter(xys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	sc.GlyphStyle.Shape = draw.CrossGlyph{}

	avgs := make(plotter.XYs, len(s.data))
	sum := 0.0
	for i, d := range s.data {
		avgs[i].X = float64(i)
		sum += float64(d)
		avgs[i].Y = sum / (float64(i+1) * float64(time.Millisecond))
	}
	l, err := plotter.NewLine(avgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	l.Color = color.RGBA{G: 255, A: 255}

	g := plotter.NewGrid()
	g.Horizontal.Color = color.RGBA{R: 255, A: 255}
	g.Vertical.Width = 0

	p, err := plot.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	p.Add(sc, l, g)
	p.Title.Text = "Latency"
	p.Y.Label.Text = "ms"
	p.X.Label.Text = "sample"

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = wt.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Stats) Hist(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	vs := make(plotter.Values, len(s.data))
	for i, d := range s.data {
		vs[i] = float64(d) / float64(time.Millisecond)
	}

	h, err := plotter.NewHist(vs, 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p, err := plot.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p.Add(h)
	p.Title.Text = "Distribution"
	p.X.Label.Text = "ms"

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = wt.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
