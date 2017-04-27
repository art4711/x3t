package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"x3t/xt"
)

var _ = tmpls.Add("map", `
{{template "header"}}
<div style="width: 95%; height: 95%">
	<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="100%" height="100%" viewBox="0 0 24 20" id="themap">
		<defs>
		  <g id="sunny">
		    <path d="M0 0h24v24H0z" fill="none"/>
		    <path d="M6.76 4.84l-1.8-1.79-1.41 1.41 1.79 1.79 1.42-1.41zM4 10.5H1v2h3v-2zm9-9.95h-2V3.5h2V.55zm7.45 3.91l-1.41-1.41-1.79 1.79 1.41 1.41 1.79-1.79zm-3.21 13.7l1.79 1.8 1.41-1.41-1.8-1.79-1.4 1.4zM20 10.5v2h3v-2h-3zm-8-5c-3.31 0-6 2.69-6 6s2.69 6 6 6 6-2.69 6-6-2.69-6-6-6zm-1 16.95h2V19.5h-2v2.95zm-7.45-3.91l1.41 1.41 1.79-1.8-1.41-1.41-1.79 1.8z"/>
		  </g>
		  <g id="silicon">
		    <path d="M-74 29h48v48h-48V29z" fill="none"/>
		    <path d="M22 9V7h-2V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-2h2v-2h-2v-2h2v-2h-2V9h2zm-4 10H4V5h14v14zM6 13h5v4H6zm6-6h4v3h-4zM6 7h5v5H6zm6 4h4v6h-4z"/>
		    <path d="M0 0h24v24H0zm0 0h24v24H0z" fill="none"/>
		  </g>
		  <g id="ore">
		    <path d="M0 0h24v24H0z" fill="none"/>
		    <path d="M14 6l-3.75 5 2.85 3.8-1.6 1.2C9.81 13.75 7 10 7 10l-6 8h22L14 6z"/>
		  </g>
		  <g id="dock">
		    <path d="M0 0h24v24H0z" fill="none"/>
		    <path d="M20 21c-1.39 0-2.78-.47-4-1.32-2.44 1.71-5.56 1.71-8 0C6.78 20.53 5.39 21 4 21H2v2h2c1.38 0 2.74-.35 4-.99 2.52 1.29 5.48 1.29 8 0 1.26.65 2.62.99 4 .99h2v-2h-2zM3.95 19H4c1.6 0 3.02-.88 4-2 .98 1.12 2.4 2 4 2s3.02-.88 4-2c.98 1.12 2.4 2 4 2h.05l1.89-6.68c.08-.26.06-.54-.06-.78s-.34-.42-.6-.5L20 10.62V6c0-1.1-.9-2-2-2h-3V1H9v3H6c-1.1 0-2 .9-2 2v4.62l-1.29.42c-.26.08-.48.26-.6.5s-.15.52-.06.78L3.95 19zM6 6h12v3.97L12 8 6 9.97V6z"/>
		  </g>
		</defs>
		<style>
			.s {
				width: 0.8;
				height: 0.8;
				stroke: black;
				stroke-width: 0.02;
			}
			{{/*Argon*/}}
			.r1 {
				fill: #a0a0ff;
			}
			{{/*Boron*/}}
			.r2 {
				fill: #a0ffa0;
			}
			{{/*Split*/}}
			.r3 {
				fill: #ffa0ff;
			}
			{{/*Paranid*/}}
			.r4 {
				fill: #ffa0a0;
			}
			{{/*Teladi*/}}
			.r5 {
				fill: #ffffa0;
			}
			{{/*Xenon*/}}
			.r6 {
				fill: #b06666;
			}
			{{/*Kha'ak*/}}
			.r7 {
				fill: #caa5a5;
			}
			{{/*Pirates*/}}
			.r8 {
				fill: #727272;
			}
			{{/*Goner*/}}
			.r9 {
				fill: #6060ff;
			}
			{{/*ufo?*/}}
			.r10 {
			}
			{{/*hostile?*/}}
			.r11 {
			}
			{{/*neutral*/}}
			.r12 {
				fill: #a0a0a0;
			}
			{{/*friendly?*/}}
			.r13 {
			}
			{{/*unknown*/}}
			.r14 {
				fill: #a0a0a0;
			}
			{{/*unused?*/}}
			.r15 {
			}
			{{/*unused?*/}}
			.r16 {
			}
			{{/*ATF*/}}
			.r17 {
				fill: #80ff80;
			}
			{{/*Terran*/}}
			.r18 {
				fill: #b0ffb0;
			}
			{{/*Yaki*/}}
			.r19 {
				fill: #ffff80;
			}
			.sectorname {
				font-size: 1;
			}
			.sectordesc {
				font-size: 0.6;
			}
			.zoomedsector {
				transform: scale(3);
			}
			.sectorhover {
				fill-opacity: 0;
				stroke-width: 0;
			}
		</style>
		<g>
{{- range .Sectors}}
{{template "map-sector" .}}
{{- end}}
		</g>
	</svg>
</div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
<script src='/js/svg-pan-zoom.min.js'></script>
<script>
svgPanZoom("#themap")
$(document).ready(function() {
	$(".sectorhover").hover(
	  function() {
	    var t = this.parentElement;
	    t.parentElement.appendChild(t);
	    $(t).find("g:first").addClass("zoomedsector");
	  }, function() {
	    var t = this.parentElement;
	    $(t).find("g:first").removeClass("zoomedsector");
  	});
});
</script>
{{template "footer"}}
`)

var _ = tmpls.Add("map-sector", `
<g transform="translate({{.X}} {{.Y}})" class="sector">
 <g>
  <rect class="s r{{.R}}" />
  <g transform="scale(0.12) translate(0.5 1.3)">
{{- range $i, $row := (sectorName .)}}
    <text y="{{$i}}" class="sectorname">{{$row}}</text>
{{- end}}
{{- range $i, $icon := (sectorIcons .)}}
    <use transform="translate({{$i}} 4) scale(0.04)" xlink:href="#{{$icon}}" />
{{- end}}
  </g>
 </g>
 <rect class="s sectorhover" />
</g>
`)

func (st *state) mapFuncs(fm template.FuncMap) {
	fm["sectorName"] = func(s xt.Sector) []string {
		sp := strings.Split(s.Name(st.text), " ")
		ret := make([]string, 0)
		for _, e := range sp {
			// If two substrings are shorter than 11, combine them
			if len(ret) != 0 && len(ret[len(ret)-1])+len(e) < 11 {
				ret[len(ret)-1] = ret[len(ret)-1] + " " + e
			} else {
				ret = append(ret, e)
			}
		}
		return ret
	}
	fm["sectorIcons"] = func(s xt.Sector) []string {
		ret := []string{}
		if s.SunPercent() > 150 {
			ret = append(ret, "sunny")
		}
		sil, ore := 0, 0
		for i := range s.Asteroids {
			switch s.Asteroids[i].Type {
			case 0:
				ore += s.Asteroids[i].Amount
			case 1:
				sil += s.Asteroids[i].Amount
			}
		}
		if sil > 300 {
			ret = append(ret, "silicon")
		}
		if ore > 300 {
			ret = append(ret, "ore")
		}
		if len(s.Docks) > 0 {
			ret = append(ret, "dock")
		}
		return ret
	}
}

func (st *state) showMap(w http.ResponseWriter, req *http.Request) {
	err := st.tmpl.ExecuteTemplate(w, "map", st.u)
	if err != nil {
		log.Fatal(err)
	}
}
