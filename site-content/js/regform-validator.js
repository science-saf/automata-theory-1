$(document).ready(function()
{
   var regexpsContainer = {
    nickname: new RegExp(/^[A-z0-9_]+$/),
    email: new RegExp(/^[A-z0-9_]+@(?:gmail\.com|yandex\.ru|mail\.ru)$/i),
    password1: new RegExp(/[A-z]+/),
    password2: new RegExp(/[\d]+/)
  };
  var fieldsWithErrors = [];

  var form = $("#regform");

  form.submit(function(event) {
    if (!isFormValid())
    {
      event.preventDefault();
      refreshErrors();
    }
  });

  function isFormValid()
  {
    var result = true;
    if (!isNicknameValid($("#nickname_input").val()))
    {
      fieldsWithErrors.push("#nickname_input")
      result = false;
    }

    if (!isEmailValid($("#email_input").val()))
    {
      fieldsWithErrors.push("#email_input")
      result = false;
    }

    if (!isPasswordValid($("#password_input1").val()))
    {
      fieldsWithErrors.push("#password_input1")
      result = false;
    }

    if ($("#password_input1").val() !== $("#password_input2").val())
    {
      fieldsWithErrors.push("#password_input2")
      result = false;
    }

    return result;
  }

  function isNicknameValid(nickname)
  {
    return regexpsContainer.nickname.test(nickname);
  }

  function isEmailValid(email)
  {
    return regexpsContainer.email.test(email);
  }

  function isPasswordValid(password1)
  {
    var MIN_REQUIRED_LENGTH = 6;
    var result = true;
    result &= regexpsContainer.password1.test(password1);
    result &= (password1.length >= MIN_REQUIRED_LENGTH);

    return result;
  }

  function refreshErrors()
  {
    $(".error-block").remove();
    for (var i = 0; i < fieldsWithErrors.length; ++i)
    {
      getErrorMessageBlock(fieldsWithErrors[i]).insertBefore($(fieldsWithErrors[i]));
    }
    fieldsWithErrors = [];
  }

  function getErrorMessageBlock(fieldId)
  {
    var MSG_INVALID_NICKNAME = "Invalid nickname";
    var MSG_INVALID_EMAIL = "Invalid email";
    var MSG_INVALID_PASSWORD = "Invalid password";
    var MSG_INVALID_PASSWORD2 = "Passwords don't match";
    var msgText;
    switch (fieldId) {
      case "#nickname_input":
        msgText = MSG_INVALID_NICKNAME;
        break;
      case "#email_input":
        msgText = MSG_INVALID_EMAIL;
        break;
      case "#password_input1":
        msgText = MSG_INVALID_PASSWORD;
        break;
      case "#password_input2":
        msgText = MSG_INVALID_PASSWORD2;
        break;
      default:
		    throw "Not defined message for field " + fieldId;
    }
    var block = "<div class=\"alert alert-danger error-block\" role=\"alert\">" + msgText + "</div>";

    return $(block);
  }
});
