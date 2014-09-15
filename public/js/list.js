function exitEdit(e) {
  input = document.getElementById('listTitle').innerHTML;
  if (e.keyCode == 13) {
    if (input.length > 0 && /^[a-zA-Z0-9~!@$\^&\*\(\)\{\}\[\]\+\-\=\_\,\<\>\"\'\:\;\`\|]+$/.test(input)) {
        e.preventDefault();
        document.activeElement.blur();
        window.location = "http://" + window.location.host + "/list/" + input;
        console.log(input);
    } else {
      document.getElementById('listTitle').innerHTML = document.title;
      document.activeElement.blur();
      console.log("error msg");
    }
  }
}

function removeTodo(itemId) {
  xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", "http://" + window.location.host + "/api/remove/" + document.title + "/" + itemId, false);
  xmlHttp.send(null);

  var todoItem = document.getElementById('todo' + itemId);
  var todoBut = document.getElementById('but' + itemId);
  todoBut.parentNode.parentNode.removeChild(todoBut.parentNode)
  todoItem.parentNode.removeChild(todoItem);
  todoBut.parentNode.removeChild(todoBut);
}

function addTodo(event) {
  var input = document.getElementById("addBox").value;
  var helperLabel = document.getElementById("inputHelper");
  if (event.keyCode == 13 && input.length > 0) {
    if (input.length > 80) {
      helperLabel.innerHTML = "80 char limit";
      return;
    }
    helperLabel.innerHTML = ""
    document.getElementById("addBox").value = '';

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "http://" + window.location.host + "/api/add/" + document.title + "/" + input, false);
    xmlHttp.send(null);
    newId = xmlHttp.responseText

    var list = document.getElementById('theList');
    var entry = document.createElement('div');
    var newTodo = document.createElement('span');
    var newBut = document.createElement('img');

    entry.className = 'entry col c12'
    newTodo.id = 'todo' + newId
    newTodo.className = 't'
    newBut.id = 'but' + newId
    newBut.className = 'x'
    newBut.src = '/icon/circlex.png'
    newBut.onclick = function() { removeTodo(newBut.id.substr(3)); };

    newTodo.appendChild(document.createTextNode(input));
    newBut.appendChild(document.createTextNode("Remove"));
    entry.appendChild(newBut); 
    entry.appendChild(newTodo); 
    list.appendChild(entry); 
  }
}
