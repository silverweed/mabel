###
 Forms-related utilities
###

# Check that passwords in a form match.
# The form should contain a 'password' and a 'passwordcheck' inputs.
matchPasswords = (form) ->
	markBad = (elem) ->
		errspan = form.querySelector 'span[name=passmatch]'
		if errspan instanceof HTMLSpanElement
			errspan.innerHTML = "Passwords don't match"
			unless elem.onkeydown?
				elem.onkeydown = ->
					errspan.innerHTML = ''
		return false
	pwdCheckEl = form.elements.namedItem 'passwordcheck'
	pwdEl = form.elements.namedItem 'password'
	return markBad pwdCheckEl unless pwdEl.value == pwdCheckEl.value
	return true

window.Forms = {
	matchPasswords: matchPasswords
}
