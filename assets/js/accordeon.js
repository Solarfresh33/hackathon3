var acc = document.getElementsByClassName("accordion");
var accordionLast = document.getElementById("last-accordion");
var symbols = document.querySelectorAll('.material-symbols-outlined');

for (var i = 0; i < acc.length; i++) {
  acc[i].addEventListener("click", function() {
    var panel = this.nextElementSibling;
    var isActive = this.classList.contains("active");

    symbols.forEach(function(symbol) {
      symbol.classList.remove('rotated');
    });

    closeAllPanels();

    if (!isActive) {
      this.classList.add("active");
      panel.style.maxHeight = panel.scrollHeight + "px";
      this.querySelector('.material-symbols-outlined').classList.add('rotated');

      if (this === accordionLast) {
        this.style.borderBottomLeftRadius = "0";
        this.style.borderBottomRightRadius = "0";
      }
    } else {
      if (this === accordionLast) {
        this.style.borderBottomLeftRadius = "10px";
        this.style.borderBottomRightRadius = "10px";
      }
    }
  });
}

function closeAllPanels() {
  var panels = document.getElementsByClassName("panel");
  for (var j = 0; j < panels.length; j++) {
    var panel = panels[j];
    panel.style.maxHeight = null;
    panel.previousElementSibling.classList.remove("active");

    accordionLast.style.borderBottomLeftRadius = "10px";
    accordionLast.style.borderBottomRightRadius = "10px";
  }
}
