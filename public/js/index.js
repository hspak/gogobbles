redifyHold = false;

function redify(input) {
  redifyHold = true;
  input.classList.add('redify');
  setTimeout(function() {
    input.classList.remove('redify');
    input.classList.add('unredify');
    setTimeout(function() {
      input.classList.remove('unredify');
      redifyHold = false;
    }, 500);
  }, 500);
}

function newTodo(event) {
  input = document.getElementById("searchBox");
  if (event.keyCode == 13 && input.value.length > 0) {
    if (! /^[a-zA-Z0-9~!@$\^&\*\(\)\{\}\[\]\+\-\=\_\,\<\>\"\'\:\;\`\|]+$/.test(input.value)) {
      if (redifyHold == false) {
        redify(input);
      }
      return;
    } else if (input.value.length > 80) {
      if (redifyHold == false) {
        redify(input);
      }
      return;
    }
    window.location = "http://" + window.location.host + "/list/" + input.value;
  }
}
