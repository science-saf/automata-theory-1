var activatedControlId = "";

$(function() {

  var addCellButton = $("#add_cell");

	$("#start_button").click(function(event) {
    world.isChaosOnStart = false;
		world.start();
    deactivateControls();
	});


  $("#chaotic_start_button").click(function(event) {
    world.isChaosOnStart = true;
    world.init(world.size.x, world.size.y);
    world.start();
    deactivateControls();
  });

	$(".control").click(function(event){
    if (world.started)
      return;
      
		activateControl($(this));
		disableControls();
	});
});

var figuresContainer = {
  glider: [
    [0, 1, 0],
    [0, 0, 1],
    [1, 1, 1]
  ],
  square: [
    [1, 1],
    [1, 1]
  ],
  gun: [
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 1,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 1, 0, 1,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 1, 1, 0,  0, 0, 0, 0, 0,  1, 1, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 1,  1],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 1, 0, 0, 0,  1, 0, 0, 0, 0,  1, 1, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 1,  1],
    [1, 1, 0, 0, 0,  0, 0, 0, 0, 0,  1, 0, 0, 0, 0,  0, 1, 0, 0, 0,  1, 1, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],

    [1, 1, 0, 0, 0,  0, 0, 0, 0, 0,  1, 0, 0, 0, 1,  0, 1, 1, 0, 0,  0, 0, 1, 0, 1,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  1, 0, 0, 0, 0,  0, 1, 0, 0, 0,  0, 0, 0, 0, 1,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 1, 0, 0, 0,  1, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 1, 1, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0]
  ],
  spaceship: [
    [1, 1, 1, 1, 0],
    [1, 0, 0, 0, 1],
    [1, 0, 0, 0, 0],
    [0, 1, 0, 0, 1]
  ],
  weekender: [
    [0, 1, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 1,  0],
    [1, 1, 1, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 1, 1,  1],
    [1, 1, 1, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 1, 1,  1],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],

    [0, 0, 0, 1, 0,  0, 0, 1, 1, 0,  0, 0, 1, 0, 0,  0],
    [0, 1, 0, 1, 1,  1, 1, 0, 0, 1,  1, 1, 1, 0, 1,  0],
    [0, 1, 0, 1, 1,  1, 1, 0, 0, 1,  1, 1, 1, 0, 1,  0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0],
    [0, 0, 0, 0, 0,  1, 1, 1, 1, 1,  1, 0, 0, 0, 0,  0],

    [0, 0, 0, 0, 1,  1, 0, 0, 0, 0,  1, 1, 0, 0, 0,  0],
    [0, 0, 0, 0, 1,  1, 0, 0, 0, 0,  1, 1, 0, 0, 0,  0]
  ],
  pulsar: [
    [0, 0, 1, 1, 1,  0, 0, 0, 1, 1,  1, 0, 0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0],
    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],
    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],
    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],

    [0, 0, 1, 1, 1,  0, 0, 0, 1, 1,  1, 0, 0],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0],
    [0, 0, 1, 1, 1,  0, 0, 0, 1, 1,  1, 0, 0],
    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],
    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],

    [1, 0, 0, 0, 0,  1, 0, 1, 0, 0,  0, 0, 1],
    [0, 0, 0, 0, 0,  0, 0, 0, 0, 0,  0, 0, 0],
    [0, 0, 1, 1, 1,  0, 0, 0, 1, 1,  1, 0, 0]
  ],
  pentapole: [
    [1, 1, 0, 0, 0,  0, 0, 0],
    [1, 0, 1, 0, 0,  0, 0, 0],
    [0, 0, 0, 0, 0,  0, 0, 0],
    [0, 0, 1, 0, 1,  0, 0, 0],
    [0, 0, 0, 0, 0,  0, 0, 0],

    [0, 0, 0, 0, 1,  0, 1, 0],
    [0, 0, 0, 0, 0,  0, 0, 1],
    [0, 0, 0, 0, 0,  0, 1, 1],
  ]
};

function activateControl(control) {
  activatedControlId = control.attr("id");
	control.attr("activated", "true");
	control.removeClass("btn-default").addClass("btn-warning");
}

function deactivateControl(control) {
  activatedControlId = "";
	control.removeAttr("activated");
	control.removeClass("btn-warning").addClass("btn-default");
  isControlActivated = false;
}

function deactivateControls() {
  deactivateControl($(".control"));
}

function disableControls() {
	$(".control").attr("disabled", "disabled");
}

function enableControls() {
	$(".control").removeAttr("disabled");
}

function processControlClick(x, y) {
  console.log(x, y, activatedControlId);
  switch (activatedControlId) {
    case "add_cell":
      world.burnCell(x, y);
      break;
    case "add_glider":
      burnFigure(figuresContainer.glider, x, y);
      break;
    case "add_square":
      burnFigure(figuresContainer.square, x, y);
      break;
    case "add_gun":
      burnFigure(figuresContainer.gun, x, y);
      break;
    case "add_spaceship":
      burnFigure(figuresContainer.spaceship, x, y);
      break;
    case "add_weekender":
      burnFigure(figuresContainer.weekender, x, y);
      break;
    case "add_pulsar":
      burnFigure(figuresContainer.pulsar, x, y);
      break;
    case "add_pentapole":
      burnFigure(figuresContainer.pentapole, x, y);
      break;
    default:
      throw "Can't process control";
  }
  enableControls();
  deactivateControl($("#" + activatedControlId));
  activatedControlId = "";
}

function burnFigure(figure, xLeft, yTop) {
  var figureDimensionX = figure[0].length
  var figureDimensionY = figure.length;
  if (((xLeft + figureDimensionX) > world.size.x) || ((yTop + figureDimensionY) > world.size.y))
    return;

  for (var yCoord = 0; yCoord < figureDimensionY; ++yCoord) {
    for (var xCoord = 0; xCoord < figureDimensionX; ++xCoord) {
      if (figure[yCoord][xCoord] != 0) {
        world.burnCell(xLeft + xCoord, yTop + yCoord);
      }
    }
  }
}
