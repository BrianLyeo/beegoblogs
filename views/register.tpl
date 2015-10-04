<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN"
"http://www.w3.org/TR/html4/strict.dtd"> 

<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
	<title>Register New Account</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<script src="/static/js/input_checker.js"></script>
</head>
<body>
	<form name="register" action="register" onsubmit="return validate_form(this);" method="post">
	Your E-Mail:
	<input type="text" name="email" />
	<br />
	Your Password:
	<input type="password" name="password" />
	<br />
    Your Password Again:
	<input type="password" name="password2" />
	<br />
    Your Display Name:
	<input type="text" name="name" />
	<br />
    <input type="radio" name="showemail" value="show" checked /> Show E-Mail
    <br />
    <input type="radio" name="showemail" value="hide" /> Hide E-Mail
	<input type="submit" name="Submit" />
	</form>
</body>
</html>