
//This is still experimental
//Begin map data
var Rendered=new Array();
var TraverseX=new Array();
var TraverseY=new Array();

function gkRenderMap (mapId,size) {
	//MapIds: 0=desert, 1=ocean, 2=fire, 3=grassland, 4=bog
	var MapData=new Array();
	var field = document.getElementById("gkField");
	var a;
	var isoXYZ;
	var k = 0;
	var map = MapData[mapId];
//	var l = map.length
	for (var i=1; i<=size; i++) {
		for (var j=0; j<=size; j++) {
			if (mapId=0) {
				MapData=["PapayaWhip","IndianRed","LightSalmon","Wheat","Salmon","PaleGoldenRod","LightSalmon","Moccasin","NavajoWhite","SaddleBrown","Peru","Tan","Wheat","Moccasin","IndianRed","SandyBrown","PeachPuff","Bisque","Brown","BlanchedAlmond","Chocolate","Coral","DarkSalmon"];
				a = Math.floor((Math.random()*23)); 
			}
			else if (mapId=1) {
				MapData=["AliceBlue","Aquamarine","Aqua","Blue","CornflowerBlue","CadetBlue","Cyan","DarkSlateBLue","DarkSeaGreen","LightSeaGreen","MediumSeaGreen","MediumSpringGreen","SeaGreen","Teal"];
				a = Math.floor((Math.random()*14)); 
			}
			else if (mapId=2) {
				MapData=["Salmon","Red","Orange","OrangeRed","Tomato","Yellow","DimGrey"];
				a = Math.floor((Math.random()*7)); 
			}
			else if (mapId=3) {
				MapData=["Yellow","YellowGreen","SpringGreen","MediumSeaGreen","MediumSpringGreen","LimeGreen","LightGreen","LawnGreen","Green","GreenYellow","ForestGreen","DarkSeaGreen","DarkGreen","Chartreuse","OliveDrab"];
				a = Math.floor((Math.random()*15)); 
			}		
			else if (mapId=4) {
				MapData=["DarkGoldenRod","DarkGray","DarkKhaki","DarkOliveGreen","Olive","OliveDrab","Peru","SaddleBrown","Sienna"];
				a = Math.floor((Math.random()*9)); 
			}
			isoXYZ = new GkIsoXYZDef(i, j, 0);
			diamond = gkIsoCreateSingleDiamond(isoXYZ, MapData[a]);
			field.appendChild(diamond);
			Rendered[k] = MapData[a];
			k++;
		}
	}
}

function gkDetermineTerrainDiamondColor (mapId) {

}

function gkRestorePixel (xv,yv,size) {
	isoXYZ = new GkIsoXYZDef(xv, yv, 0);
	diamond = gkIsoCreateSingleDiamond(isoXYZ, map[k]);
	field.appendChild(diamond);
}

//This has been delayed due to the me needing to talk to Turnip about some dine details here.
/*
function gkLoadSpecialMap (location,size) {
	var pos;
	var k;
	var mapRequest = makeHttpObject();
	mapRequest.open("GET", "assets/gk/javascript/game/gkAudio.js", false);
	mapRequest.send(null);
	print(request.responseText);
	for (var i=1; i<=size; i++) {
		for (var j=1; j<=size; j++) {
			isoXYZ = new GkIsoXYZDef(i, j, 0);
			diamond = gkIsoCreateSingleDiamond(isoXYZ, SpecialMap[k]);
			field.appendChild(diamond);
			Rendered[k] = SpecialMap[k];
			k++;
		}
	}
}
*/

function gkTraverseAll (size) {
	var a = 0;
	var b = 0;
	var c = 0;
	var x = 0;
	var isoXYZ = new GkIsoXYZDef(a, b, c);
	for (; a<=size; a++) {
		for (var b=1; b<=size; b++) {
			winX, winY = GkIsoXYZDef(a, b, c);
			TraverseX[x] = winX;
			TraverseY[x] = winY;
		}
	}
}

/* This under development too
function gkRenderTexelsAll (texel,size) {
//	This will render a texture for all iso squares.
	var a = 0;
	var b = 0;
	var c = 0;
	var x = 0;
	var isoXYZ = new GkIsoXYZDef(a, b, c);
	for (; a<=size; a++) {
		for (var b=1; b<=size; b++) {
			winx, winy = GkIsoXYZDef(a, b, c)
			
		}
	}
}
*/

function gkPutShrub (x,y,z,location) {
	var putLocationX;
	var putLocationY;
	putLocationX, putLocationY = GkIsoXYZDef(x, y, z);
	field.innerHTML += '\x3Cdiv id\x3Dshrub' + location + ' style\x3D\x22position\x3A absolute\x3B top\x3A ' + putLocationY + 'px\x3B left\x3A ' + putLocationX + 'px\x3B\x22\x3E\x3Cimg src\x3D\x22' + location + '\x3D\x3E\x3C\x2Fdiv\x3E';
}