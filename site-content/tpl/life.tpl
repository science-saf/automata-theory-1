<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>{{ .title}}</title>

	<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css" />
	<link rel="stylesheet" type="text/css" href="/css/bootstrap-theme.min.css" />

	<style>
	* {
		margin: 0;
		padding: 0;
	}
	html, body {
		height: 100%;
	}
	canvas {
		display: block;
		margin: 5% auto;
	}
	</style>

	<script type="text/javascript" src="/js/jquery.js"></script>
	<script type="text/javascript" src="/js/life.js"></script>
	<script type="text/javascript" src="/js/script.js"></script>
</head>
<body>
	<div>
		<button id="add_cell" class="control btn btn-default">Add cell</button>
		<button id="add_square" class="control btn btn-default">Add square</button>
		<button id="add_glider" class="control btn btn-default">Add glider</button>
		<button id="add_pentapole" class="control btn btn-default">Add pentapole</button>
		<button id="add_spaceship" class="control btn btn-default">Add spaceship</button>
		<button id="add_gun" class="control btn btn-default">Add glider gun</button>
		<button id="add_weekender" class="control btn btn-default">Add weekender</button>
		<button id="add_pulsar" class="control btn btn-default">Add pulsar</button>
	</div>
	<div>
		<button id="start_button" class="btn btn-default">Start</button>
		<button id="chaotic_start_button" class="btn btn-default">Chaotic start</button>
	</div>
<canvas id="viewport"></canvas>
<script>

var bgColor = '#EEE',
	lineColor = '#BBB',
	cellSize = 10;

var canvas = document.getElementById('viewport');
canvas.height = 0.8 * document.body.clientHeight;
canvas.width = 0.8 * document.body.clientWidth;

var context = canvas.getContext('2d');

var topLeft = new Point;
	bottomRight = new Point,
	size = new Point;

function calcOffset()
{
	var hor = canvas.width % cellSize,
		vert = canvas.height % cellSize;

	topLeft.x = Math.round(hor / 2);
	topLeft.y = Math.round(vert / 2);

	bottomRight.x = Math.round(canvas.width - hor / 2);
	bottomRight.y = Math.round(canvas.height - vert / 2);

	size.x = Math.floor(canvas.width / cellSize) - 1,
	size.y = Math.floor(canvas.height / cellSize) - 1;
}
calcOffset();

function drawBackground()
{
	context.fillStyle = lineColor;
	context.fillRect(0, 0, canvas.width, canvas.height);
	context.fillStyle = bgColor;
	context.fillRect(
        topLeft.x,
        topLeft.y,
        bottomRight.x - topLeft.x,
        bottomRight.y - topLeft.y
    );
}
drawBackground();

function drawGreed()
{
	context.strokeStyle = lineColor;
	context.lineWidth = 1;

	context.beginPath();

	for (var i = topLeft.x; i <= bottomRight.x; i += cellSize) {
		context.moveTo(i, topLeft.y);
		context.lineTo(i, bottomRight.y);
	}

	for (var i = topLeft.y; i <= bottomRight.y; i += cellSize) {
		context.moveTo(topLeft.x, i);
		context.lineTo(bottomRight.x, i);
	}

	context.stroke();
}
drawGreed();


function fillCell(x, y, color)
{
	context.fillStyle = color;
	context.fillRect(
		topLeft.x + x * cellSize + 1,
		topLeft.y + y * cellSize + 1,
		cellSize - 2,
		cellSize - 2
	);
}


var world = new World(size.x, size.y);

world.drawCellCallback = function(x, y, color)
{
	fillCell(x, y, color);
};
world.eraseCellCallback = function(x, y)
{
	fillCell(x, y, bgColor);
};

world.speed = 50;

function canvasClick(evt)
{
	var x = Math.floor((evt.clientX - topLeft.x - 13) / cellSize);
	var y = Math.floor((evt.clientY - topLeft.y) / cellSize);
	x = x - 12; // added because of error in click coordinates
	y = y - 13; // added because of error in click coordinates
	if (!world.started && (activatedControlId !== "")) {
		processControlClick(x, y);
	} else {
		world.burnCell(x, y);
	}
}

canvas.onclick = canvasClick;

function keyPress(evt)
{
	switch (evt.keyCode) {
		case 13: // Enter
			if (world.started)
				world.stop();
			else
				world.start();
			break;

		case 32: // Space
			if (!world.started)
				world.step();
			break;

		case 91: // ]
			world.speed += 50;
			break;

		case 93: // [
			var speed = world.speed - 50;
			if (speed < 50)
				speed = 50;
			world.speed = speed;
			break;

		case 114: // r
			world.stop();
			drawBackground();
			drawGreed();
			world.init(world.size.x, world.size.y);
			world.start();
			break;

		case 99: // c
			world.clear();
			break;
	}
}
document.body.onkeypress = keyPress;
</script>
</body>
</html>
