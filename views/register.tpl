
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
