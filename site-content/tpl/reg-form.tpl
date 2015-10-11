<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/bootstrap.min.css">
    <link rel="stylesheet" href="/css/bootstrap-theme.min.css">
  </head>
  <body>
    <div class="container">
      <div class="row">
        <h2>
          {{.title}}
        </h2>
      </div>
      <div class="row">
        <form action="/form" method="POST">
          <div class="form-group">
            <label for="exampleInputNickname">Nickname</label>
            {{if .showAlertName}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <input type="text" class="form-control" id="exampleInputEmail1" placeholder="Your Nickname" name="userNickname">
            <p class="help-block">Nickname should contain only English letters, digits and underscores.</p>
          </div>
          <div class="form-group">
            <label for="exampleInputEmail">Email address</label>
            {{if .showAlertEmail}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Your Email" name="userEmail">
            <p class="help-block">Only GMail, Yandex Mail and Mail.ru email addresses allowed.</p>
          </div>
          <div class="form-group">
            <label for="exampleInputPassword">Password</label>
            {{if .showAlertPassword}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Your New Password" name="userPassword">
          </div>
          <div class="form-group">
            <input type="password" class="form-control" id="exampleInputPassword2" placeholder="Repeat Password" name="userPasswordRepeat">
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
