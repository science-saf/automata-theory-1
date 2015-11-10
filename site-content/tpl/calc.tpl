<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<title>{{.title}}</title>

	<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css" />
	<link rel="stylesheet" type="text/css" href="/css/bootstrap-theme.min.css" />

	<script type="text/javascript" src="/js/jquery.js"></script>
	<script type="text/javascript" src="/js/calc_script.js"></script>
</head>
<body>
	<div class="container">
		<div class="row alert alert-info">
			<p>
				<h4>Supported functions:</h4>
				<ul>
					<li>rand(m, n) - generates random value in range from m to n</li>
					<li>sqrt(n) - square root from n</li>
				</ul>
			</p>

			<p>
				Trigonometric functions, where n - angle in <strong>degrees</strong>
				<ul>
					<li>sin(n)</li>
					<li>cos(n)</li>
					<li>tg(n)</li>
					<li>ctg(n)</li>
				</ul>
			</p>

			<p>
				Inverse trigonometric functions (return angle in <strong>degrees</strong>)
				<ul>
					<li>arcsin(n)</li>
					<li>arccos(n)</li>
					<li>arctg(n)</li>
					<li>arcctg(n)</li>
				</ul>
			</p>
		</div>

		<div class="row">
			<form action="" method="POST" id="calculator_form">
				<div class="form-group">
					<label for="expression">
						Enter expression:
					</label>
					<input type="text" class="form-control" id="expression" name="expression" />
				</div>
				<div class="form-group">
					<button type="submit" class="btn btn-default">Calculate</button>
				</div>
			</form>
		</div>

		<div class="row alert alert-danger" id="errors" {{if .errors}}style="display:block;"{{else}}style="display:none;"{{end}}>
			{{.errors}}
		</div>

		<div class="row" {{if .errors}}style="display:none;"{{else}}style="display:none;"{{end}} id="result">{{.result}}</div>
	</div>
</body>
</html>
