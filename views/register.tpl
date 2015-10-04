<!DOCTYPE html>

<html>
<head>
	<title>Register New Account</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
	<form name="register" action="register" method="post">
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
    <input type="radio" name="showemail" value="show" /> Show E-Mail
    <br />
    <input type="radio" name="showemail" value="hide" /> Hide E-Mail
	<input type="submit" name="Submit" />
	</form>
</body>
</html>