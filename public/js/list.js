function removeTodo(itemId) {
  var xmlHttp = null;

  xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", "http://localhost:3000/remove/" + document.title + "/" + itemId, false);
  xmlHttp.send(null);
  console.log(xmlHttp.responseText);
  
  var todoItem = document.getElementById('todo' + itemId);
  var todoBut = document.getElementById('but' + itemId);
  todoItem.parentNode.removeChild(todoItem);
  todoBut.parentNode.removeChild(todoBut);
}

function addTodo() {
  if (event.keyCode == 13) {
    var input = document.getElementById("addBox").value;
    document.getElementById("addBox").value = '';

    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "http://localhost:3000/add/" + document.title + "/" + input, false);
    xmlHttp.send(null);
    console.log(xmlHttp.responseText);
    newId = xmlHttp.responseText

    var list = document.getElementById("theList");
    var newTodo = document.createElement('li');
    var newBut = document.createElement('button');
    newTodo.id = 'todo' + newId
    newBut.id = 'but' + newId
    newBut.onclick = function() {
      removeTodo(newBut.id.substr(3));
    };
    newTodo.appendChild(document.createTextNode(input));
    newBut.appendChild(document.createTextNode("Remove"));
    list.appendChild(newBut); 
    list.appendChild(newTodo); 
  }
}
