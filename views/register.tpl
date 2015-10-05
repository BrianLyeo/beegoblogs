<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN"
"http://www.w3.org/TR/html4/strict.dtd"> 

<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
  <title>Register New Account</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">
    *{margin:0; border:5; padding:5; font-size:13pt;}
    #TopNav{height:40px;border-top:#060 2px solid;border-bottom:#060 2px solid;background-color:#690;}
    #TopNav ul{list-style:none;margin-left:50px;}
    #TopNav li{display:inline;line-height:40px;}
    #TopNav a{color:#fff;text-decoration:none;padding:20px 20px;}
    #TopNav a:hover{background-color:#060;}
    #RegisterForm
    {
      padding-top: 10px;
      padding-right: 0.25em;
      padding-bottom: 2ex;
      padding-left: 20%;
    }
  </style>
  <script src="/static/js/input_checker.js"></script>
</head>
<body>
  <div id="TopNav">
    <ul id="TopNavList">
      <li><a href="/main" title="MainPage">Main Page </a></li>
      <li><a href="/login" title="Login">Sign in </a></li>
      <li><a href="/login?action=logout" title="Logout">Sign out </a></li>
      <li><a href="/register" title="Register">Sign up </a></li>
    </ul>
  </div> <!-- TopNav -->
  <div id="RegisterForm">
    <form name="register" action="register" onsubmit="return validate_form(this);" method="post">
      <div id="FormText">
        Your E-Mail:
      </div>
      <input type="text" name="email" />
      <br />
      <div id="FormText">
        Your Password:
      </div>
      <input type="password" name="password" />
      <br />
      <div id="FormText">
        Your Password Again:
      </div>
      <input type="password" name="password2" />
      <br />
      <div id="FormText">
        Your Display Name:
      </div>
      <input type="text" name="name" />
      <br />
      <input type="radio" name="showemail" value="show" checked /> Show E-Mail
      <br />
      <input type="radio" name="showemail" value="hide" /> Hide E-Mail
      <div id="FormButton">
        <input type="submit" name="Submit" />
      </div>
    </form>
  </div>
</body>
</html>