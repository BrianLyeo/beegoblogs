<!DOCTYPE html>

<html>
<head>
  <title>Beego Blogs Login</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">
    *{margin:0; border:5; padding:5; font-size:13pt;}
    #TopNav{height:40px;border-top:#060 2px solid;border-bottom:#060 2px solid;background-color:#690;}
    #TopNav ul{list-style:none;margin-left:50px;}
    #TopNav li{display:inline;line-height:40px;}
    #TopNav a{color:#fff;text-decoration:none;padding:20px 20px;}
    #TopNav a:hover{background-color:#060;}
    #LoginForm
    {
      margin:0;
      border:5;
      padding:5;
      padding-top: 10px;
      padding-right: 0.25em;
      padding-bottom: 2ex;
      padding-left: 20%;
    }
  </style>
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
  
  <div id="LoginForm">
    <form name="login" action="login" method="post">
      <div id="FormText">
        Your E-Mail:
      </div>
      <div id="FormInput">
        <input type="text" name="email" />
      </div>
      <br />
      <div id="FormText">
        Your Password:
      </div>
      <div id="FormInput">
        <input type="password" name="password" />
      </div>
      <br />
      <div id="FormButton">
        <input type="submit" name="Submit" />
      </div>
    </form>
  </div> <!-- LoginForm -->
</body>
</html>