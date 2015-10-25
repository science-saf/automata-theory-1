<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/bootstrap.min.css">
    <link rel="stylesheet" href="/css/bootstrap-theme.min.css">
    <script src="/js/jquery.js"></script>
    <script src="/js/regform-validator.js"></script>
  </head>
  <body>
    <div class="container">
      <div class="row">
        <h2>
          {{.title}}
        </h2>
      </div>
      <div class="row">
        <form action="/form" method="POST" id="regform">
          <div class="form-group">
            <label for="nickname_input">Nickname</label>
            {{if .showAlertName}}
              <div class="alert alert-danger error-block" role="alert">{{.alertName}}</div>
            {{end}}
            <input type="text" class="form-control" id="nickname_input" placeholder="Your Nickname" name="userNickname" value="{{.oldNickname}}">
            <p class="help-block">Nickname should contain only English letters, digits and underscores.</p>
          </div>
          <div class="form-group">
            <label for="email_input">Email address</label>
            {{if .showAlertEmail}}
              <div class="alert alert-danger error-block" role="alert">{{.alertEmail}}</div>
            {{end}}
            <input type="email" class="form-control" id="email_input" placeholder="Your Email" name="userEmail" value="{{.oldEmail}}">
            <p class="help-block">Only GMail, Yandex Mail and Mail.ru email addresses allowed.</p>
          </div>
          <div class="form-group">
            <label for="password_input1">Password</label>
            {{if .showAlertPassword1}}
              <div class="alert alert-danger error-block" role="alert">{{.alertPassword1}}</div>
            {{end}}
            <input type="password" class="form-control" id="password_input1" placeholder="Your New Password" name="userPassword" value="{{.oldPassword1}}">
          </div>
          <div class="form-group">
            {{if .showAlertPassword2}}
              <div class="alert alert-danger error-block" role="alert">{{.alertPassword2}}</div>
            {{end}}
            <input type="password" class="form-control" id="password_input2" placeholder="Repeat Password" name="userPasswordRepeat" value="{{.oldPassword2}}">
            <p class="help-block">Password should have at least 6 characters with letters and digits</p>
          </div>
          <!--<div class="checkbox">
            <label>
              <input type="checkbox">Don't remember me
            </label>
          </div>-->
          <button type="submit" class="btn btn-default">Submit</button>
        </form>
      </div>
    </div>
  </body>
</html>
