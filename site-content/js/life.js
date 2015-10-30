var Point = function(){};
Point.prototype = {
	x: null,
	y: null
};

var Cell = function(){};
Cell.prototype = {
	age: 0,

	grow: function()
	{
		this.age++;
	},

	getColor: function()
	{
		color = this.age * 5;
		if (color > 250)
			color = 250;

		return 'hsl(' + color + ', 100%, 50%)';
	}
}

var World = function()
{
	var x = 10,
		y = 10;

	if (arguments.length === 1)
		x = y = parseInt(arguments[0]);
	else if (arguments.length === 2) {
		x = parseInt(arguments[0]);
		y = parseInt(arguments[1]);
	}

	this.init(x, y);
};

World.prototype = {
	size: new Point,
	speed: 100,

	drawCellCallback: function(x, y, color){},
	eraseCellCallback: function(x, y){},

	age: 0,
	map: [],

	started: false,

	_borderMap: [],
	_changeQueue: [],
	_timeout: null,

	init: function(x, y)
	{
		this.age = 0;

		this.size.x = x;
		this.size.y = y;

		this.map = [];
		for (var i = 0; i <= this.size.x; i++)
			this.map[i] = [];

		x = y = 0;

		for (var i = 0; i < (this.size.x * this.size.y) / 2; i++) {

			x = Math.round(Math.random() * size.x);
			if (!this.map[x])
				this.map[x] = [];

			y = Math.round(Math.random() * size.y);
			this.map[x][y] = new Cell;

			this.drawCellCallback(x, y);
		}

		this._changeQueue = [
			[undefined, undefined],
			[undefined, undefined]
		];

	},

	start: function()
	{
		this.started = true;
		this.step();
	},

	stop: function()
	{
		clearTimeout(this._timeout);
		this.started = false;
	},

	step: function()
	{
		this.age++;

		this._borderMap = [];

		var burnNote = [],
			deathNote = [];

		var cnt = 0;
		var color;

		for (var i in this.map) {
			for (var j in this.map[i]) {

				cnt = this.checkNeighbors(i, j, true);

				if (cnt < 2 || cnt > 3)
					deathNote.push([i, j]);

				this.drawCellCallback(i, j, this.map[i][j].getColor());

				this.map[i][j].grow();
			}
		}

		for (var i in this._borderMap) {
			for (var j in this._borderMap[i]) {

				cnt = this.checkNeighbors(i, j, false);

				if (cnt === 3)
					burnNote.push([i, j]);
			}
		}

		for (var i in burnNote)
			this.burnCell(burnNote[i][0], burnNote[i][1]);


		for (var i in deathNote)
			this.killCell(
				deathNote[i][0],
				deathNote[i][1]
			);

		if (!deathNote.length)
			this.stop();

		this._changeQueue.push([burnNote.length, deathNote.length]);
		if (this._changeQueue.length > 3)
			this._changeQueue.shift();

		if (this._changeQueue[2][0] === deathNote.length &&
			this._changeQueue[2][1] === burnNote.length &&
			this._changeQueue[1][0] === deathNote.length &&
			this._changeQueue[1][1] === burnNote.length &&
			this._changeQueue[0][0] === deathNote.length &&
			this._changeQueue[0][1] === burnNote.length
		)
			this.stop();

		if (this.started) {
			var that = this;
			this._timeout = setTimeout(
				function()
				{
					if (that.started)
						that.step();
				},
				this.speed
			);
		}
	},

	checkNeighbors: function(x, y, toCheckList)
	{

		var count = 0,
			ox = 0,
			oy = 0;

		x = parseInt(x);
		y = parseInt(y);

		var width = this.size.x + 1,
			height = this.size.y + 1;

		for (var i = x - 1; i <= x + 1; i++) {
			ox = (i + width) % width;

			for (var j = y - 1; j <= y + 1; j++) {
				if (i === x && j === y)
					continue;

				oy = (j + height) % height;

				if (this.map[ox] && this.map[ox][oy] !== undefined)
					count++;
				else if (toCheckList) {
					if (!this._borderMap[ox])
						this._borderMap[ox] = [];
					this._borderMap[ox][oy] = true;
				}
			}
		}

		return count;
	},

	burnCell: function(x, y)
	{
		if (!this.map[x])
			this.map[x] = [];
		this.map[x][y] = new Cell;

		this.drawCellCallback(x, y, 'hsl(350, 100%, 50%)');
	},

	killCell: function(x, y)
	{
		delete this.map[x][y];

		this.eraseCellCallback(x, y);
	},

	clear: function()
	{
		for (var i in this.map)
			for (var j in this.map[i])
				this.killCell(i, j);
	}
}
