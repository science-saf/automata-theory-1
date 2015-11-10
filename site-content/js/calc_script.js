$(function() {
  var isProcessing = false;
  var errorsBlock = $("#errors");
  var resultBlock = $("#result");

  $("#calculator_form").submit(function(event) {
    event.preventDefault();
    sendExpression();

    return false;
  });

  $("#expression").on("keyup focus blur", sendExpression);

  function sendExpression() {
    if (!isProcessing) {
      isProcessing = true;
      var url = $(this).attr("action");
      $.post(url, {
          expression: $("#expression").val(),
          is_ajax: true
        },
        function(){},
        "JSON")
        .done(handlePost)
        .fail(handleError)
        .always(function() { isProcessing = false; });
    }
  }

  function handleError(obj, status) {
    errorsBlock.html("Request has not succeed. (" + status + ")");
  }

  function handlePost(response) {
    if ((response.errors === null) || (response.errors.length === 0)) {
      errorsBlock.hide();
      resultBlock.html(response.result);
      resultBlock.show();
    } else {
      var errorsBlockContent = "";
      for (var i = 0; i < response.errors.length; ++i) {
        errorsBlockContent += response.errors[i];
        if (i < (response.errors.length - 1)) {
          errorsBlockContent += "<br />";
        }
      }
      errorsBlock.html(errorsBlockContent);
      errorsBlock.show();
      resultBlock.hide();
    }
  }
});
