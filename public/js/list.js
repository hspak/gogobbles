function exitEdit(e) {
  // detect and convert &gt; stuff to >
  input = document.getElementById('listTitle').innerHTML;
  if (e.keyCode == 13) {
    e.preventDefault();
    document.activeElement.blur();
    if (input.length > 0 && /^[a-zA-Z0-9]+$/.test(input)) {
        window.location = "//" + window.location.host + "/list/" + input;
        console.log(input);
    } else {
      document.getElementById('listTitle').innerHTML = document.title;
      console.log("error msg");
    }
  }
}

function removeTodo(itemId) {

  xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", "//" + window.location.host + "/api/remove/" + document.title + "/" + itemId, false);
  xmlHttp.send(null);

  var todoItem = document.getElementById('todo' + itemId);
  var todoBut = document.getElementById('but' + itemId);
  todoItem.parentNode.classList.add('horizTranslate');
  setTimeout(function() {
    todoBut.parentNode.parentNode.removeChild(todoBut.parentNode)
    todoItem.parentNode.removeChild(todoItem);
    todoBut.parentNode.removeChild(todoBut);
  }, 400);
}

function addTodo(event) {
  var inputOrig = document.getElementById("addBox").value;
  if (event.keyCode == 13 && inputOrig.length > 0) {
    document.getElementById("addBox").value = '';
    addElem(inputOrig);
  }
}

redifyHold = false;

function addElem(text, refId) {
  var box = document.getElementById("addBox");
  thisId = refId;

  if (arguments.length == 1) {
    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "//" + window.location.host + "/api/add/" + document.title + "/" + text, false);
    xmlHttp.send(null);
    if (xmlHttp.status != 200) {
      if (redifyHold == true) {
        return;
      }
      redifyHold = true;
      box.classList.add('redify-helper');
      box.parentNode.classList.add('redify');
      setTimeout(function() {
        box.parentNode.classList.remove('redify');
        box.parentNode.classList.add('unredify');
        setTimeout(function() {
          box.classList.remove('redify-helper');
          box.parentNode.classList.remove('unredify');
          redifyHold = false;
        }, 350);
      }, 450);
      return;
    }
    thisId = xmlHttp.responseText;
    box.className = 'inputbox col c12';
  }

  if (text.length > 38) {
    text = text.slice(0, 35);
    text += "...";
  }

  var list = document.getElementById('theList');
  var entry = document.createElement('div');
  var newTodo = document.createElement('span');
  var newBut = document.createElement('img');

  entry.className = 'entry col c12'
  newTodo.id = 'todo' + thisId
  newTodo.className = 't'
  newBut.id = 'but' + thisId
  newBut.className = 'x'
  newBut.src = '/image/circlex.png'
  newBut.onclick = function() { removeTodo(newBut.id.substr(3)); };

  newTodo.appendChild(document.createTextNode(text));
  entry.appendChild(newBut);
  entry.appendChild(newTodo);
  list.appendChild(entry);
  setTimeout(function() {
    entry.className += ' load col c12';
  }, 10);
}

setInterval(function() {
    xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "//" + window.location.host + "/api/get/" + document.title, false);
    xmlHttp.send(null);
    resp = JSON.parse(xmlHttp.responseText);

    for (var i = 0; i < resp.Count; i++) {
      var obj = resp.Todos[i].Id;
      var text = resp.Todos[i].Text;
      var todoItem = document.getElementById('todo' + obj);
      if (todoItem == null) {
        addElem(text, obj);
      }
    }
}, 3000);
