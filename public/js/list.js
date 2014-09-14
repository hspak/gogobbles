function removeTodo(itemId) {
  xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", "http://" + window.location.host + "/api/remove/" + document.title + "/" + itemId, false);
  xmlHttp.send(null);
  
  var todoItem = document.getElementById('todo' + itemId);
  var todoBut = document.getElementById('but' + itemId);
  todoItem.parentNode.removeChild(todoItem);
  todoBut.parentNode.removeChild(todoBut);
}

function addTodo() {
  var input = document.getElementById("addBox").value;
  var helperLabel = document.getElementById("inputHelper");
  if (event.keyCode == 13 && input.length > 0) {
    if (! /^[a-zA-Z0-9~!@$\^&\*\(\)\{\}\[\]\+\-\=\_\,\<\>\"\'\:\;\`\|]+$/.test(input)) {
      helperLabel.innerHTML = "No"
      return;
    } else if (input.length > 80) {
      helperLabel.innerHTML = "80 char limit";
      return;
    }
    helperLabel.innerHTML = ""
    document.getElementById("addBox").value = '';

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "http://" + window.location.host + "/api/add/" + document.title + "/" + input, false);
    xmlHttp.send(null);
    newId = xmlHttp.responseText

    var list = document.getElementById("theList");
    var newTodo = document.createElement('p');
    var newBut = document.createElement('img');
    newTodo.id = 'todo' + newId
    newBut.id = 'but' + newId
    newBut.className = 'x'
    newBut.src = '/icon/circlex.png'
    newBut.onclick = function() {
      removeTodo(newBut.id.substr(3));
    };
    newTodo.appendChild(document.createTextNode(input));
    newBut.appendChild(document.createTextNode("Remove"));
    list.appendChild(newBut); 
    list.appendChild(newTodo); 
  }
}
