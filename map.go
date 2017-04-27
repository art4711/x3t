package main

import (
	"log"
	"net/http"
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
    <text transform="translate(0 {{$i}})" class="sectorname">{{$row}}</text>
{{- end}}
{{if gt .SunPercent 150}}
    <use transform="translate(0 4) scale(0.04)" xlink:href="#sunny" />
{{end}}
  </g>
 </g>
 <rect class="s sectorhover" />
</g>
`)

func (st *state) showMap(w http.ResponseWriter, req *http.Request) {
	err := st.tmpl.ExecuteTemplate(w, "map", st.u)
	if err != nil {
		log.Fatal(err)
	}
}
