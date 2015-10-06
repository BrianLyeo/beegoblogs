  <div id="MainPage">
    {{range $index, $blog := .BlogsList}}
      <li><a href="/main/{{$blog.Id}}" title="{{$blog.Title}}">{{$blog.Title}} </a></li>
    {{else}}
      <h1>No Blogs Here!</h1>
    {{end}}
  </div> <!-- MainPage --> 
  <h1>Current Operation User is {{.User}}</h1>