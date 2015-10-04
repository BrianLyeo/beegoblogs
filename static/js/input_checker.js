/*
 * checking for form input
 */
 
function check_not_empty(field, alerttxt)
{
	with (field)
	{
		if (value==null||value=="")
		{alert(alerttxt);return false}
		else {return true}
	}
}
function validate_email(field,alerttxt)
{
	with (field)
	{
		apos=value.indexOf("@")
		dotpos=value.lastIndexOf(".")
		if (apos < 1 || dotpos - apos < 2) 
		{
			alert(alerttxt);
			return false;
		}
		else
		{
			return true;
		}
	}
}

function check2pwd(p1, p2, alerttxt) {   	
		if (p1.value != p2.value) {  
			alert(alerttxt); 
			p1.value = "";  
			p2.value = ""; 
			return false;
		} 
		else
		{
			return true;
		}
} 

function validate_form(thisform)
{
	with (thisform)
	{
		if (validate_email(email, "Not a valid e-mail address!") == false)
		{
			email.focus();
			return false;
		}
		if (check_not_empty(password, "password should not be empty") == false || check2pwd(password, password2, "input password not same!") == false)
		{
			password.focus();
			return false;
		}
		
		if (check_not_empty(name, "name should not be empty") == false)
		{
			name.focus();
			return false;
		}
	}
}