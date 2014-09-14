function newTodo() {
  input = document.getElementById("searchBox").value;
  if (event.keyCode == 13 && input.length > 0) {
    if (! /^[a-zA-Z0-9~!@$\^&\*\(\)\{\}\[\]\+\-\=\_\,\<\>\"\'\:\;\`\|]+$/.test(input)) {
      helperLabel.innerHTML = "No"
      return;
    } else if (input.length > 80) {
      helperLabel.innerHTML = "80 char limit";
      return;
    } 
    window.location = "http://" + window.location.host + "/list/" + input;
  }
}
