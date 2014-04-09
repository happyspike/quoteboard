function refreshQuotes() {
  var r = new XMLHttpRequest();
  r.open("GET", "/quotes", true);
  r.onreadystatechange = function () {
    if (r.readyState != 4 || r.status != 200) return;
    element = document.getElementById("quoteList");
    element.innerHTML = r.responseText;
  };
  r.send();

  setTimeout("refreshQuotes();",4000);
};

refreshQuotes();