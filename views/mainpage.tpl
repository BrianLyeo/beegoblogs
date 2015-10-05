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
  
  <div id="MainPage">
    {{range $index, $blog := .BlogsList}}
      <li><a href="/main/{{$blog.Id}}" title="{{$blog.Title}}">{{$blog.Title}} </a></li>
    {{else}}
      <h1>No Blogs Here!</h1>
    {{end}}
  </div> <!-- MainPage --> 
  <h1>Current Operation User is {{.User}}</h1>
</body>
</html>