//http://us1.internet-radio.com:8180/
package xmp

const Tmpl = `{{define "raxx"}}
<!DOCTYPE html>
<html>
<head>
<style type=text/css>
body {padding: 0px;margin: 0px; border: 0px;}
#dx { width: 20px;height: 30px;background: #232323;padding: 0px;margin: 0px;border: 0px;position: fixed;left: 0px; top: 0px;cursor: default;}
#d0 { width: 200px;height: 30px;background: #232323;border: 0px;position: fixed;top: 0px;left: 20px;color: lightgray;font-family:Arial;font-size: 10px;font-weight: lighter;font-style: normal;text-align: center;cursor: default;}
#p0 { width: 200px;height: 30px;border: 0px;padding: 0px;background: #232323;color: lightgray;font-family:Arial;font-size:10px;font-weight: lighter;font-style: italic;text-align: center;display: table-cell;vertical-align: middle;top: 0px;left: 20px;cursor: default;}
</style>
  <script>
  var i = 0;
  var sr =
  [{{range .}}{sn:"http://{{.Ln}}/;", tl:"{{.Nm}}", nm:"{{.Sn}}", src:"http://{{.Ln}}/"},{{end}}];

    function rwnd(){
    var al = document.getElementById("al");
    var cl = document.getElementById("cl");
    var d0 = document.getElementById("d0");
    al.pause();
    cl.style.fill='#232323';
    i++;
    if (i==sr.length){i=0};
    d0.innerHTML='<p id=p0>'+sr[i].tl+'</p>';
  }

  function rest(){
    i=0;
  }

  function play(){
    var al = document.getElementById("al");
    var cl = document.getElementById("cl");
    var d0 = document.getElementById("d0");

    if (al.paused) {
      al.src=sr[i].sn;
      al.play();
      cl.style.fill="red";
      anim(0);
    } else {
      al.pause();
      cl.style.fill="#232323";
      anim(1);
    }
  }

  function anim(i){
    var ca = document.getElementById("ca");
    if (i==0){
      ca.setAttribute("from",0);
      ca.setAttribute("to",5);
      ca.setAttribute("begin","0s");
      ca.setAttribute("dur","1s");
      ca.setAttribute("fill","remove");
      ca.setAttribute("repeatCount","indefinite");
    }else{
      ca.setAttribute("from",3);
      ca.setAttribute("to",0);
      ca.setAttribute("begin","");
      ca.setAttribute("dur","");
      ca.setAttribute("fill","");
      ca.setAttribute("repeatCount","");
    }
  }

  function httpGet(theUrl){
    var xmlHttp = null;
    xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", theUrl, false );
    xmlHttp.send(null);
    return xmlHttp.responseText;}

  function out(){
    var d0 = document.getElementById("d0");
    var a = httpGet(sr[i].src);
    var a = a.substr(a.search('Current Song'),a.length-1);
    var a = a.substr(a.search('<b>')+3,a.search('</b>'));
    console.log(a);
    d0.innerHTML='<p id=p0>' + a + '</p>'}

   function over(){
      var d0 = document.getElementById("d0");
      d0.innerHTML ='<p>'+sr[i].tl+'</p>';
    }

  function vol(delta) {
    vl=document.getElementById("al").volume;
    if (delta < 0){vl-=0.05;}else{vl+=0.05;}
    document.getElementById("al").volume=vl;
  }

  function wheel(event){
    var delta = 0;
    if (!event) event = window.event;
      if (event.wheelDelta) {
        delta = event.wheelDelta/120;
        if (window.opera) delta = -delta;
      } else if (event.detail) {
        delta = -event.detail/3;
      }
    if (delta)
    vol(delta);
}

if (window.addEventListener)
  window.addEventListener('DOMMouseScroll', wheel, false);
  window.onmousewheel = document.onmousewheel = wheel;

window.setInterval(out, 10000);

 </script>

</head>
<body>
  <div id=dx>
    <audio id=al autobuffer></audio>
    <svg width="20" height="30" onclick=play() oncontextmenu=rwnd() onmousewheel=vol(delta) ondblclick=rest() onmouseover=over() onmouseout=out()>
      <polygon id=pl points="0,0 15,0 20,15 15,30, 0,30" style="fill:lightgray;fill-rule:evenodd;"/>
      <circle id=cl cx="9" cy="15" r="3" stroke="" stroke-width="0" fill="#232323">
        <animate id=ca attributeName="r" attributeType="XML"/>
      </circle>
  </div>
  <div id=d0></div>
</body>
</html>
{{end}}
`
