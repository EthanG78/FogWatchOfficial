/* When the user clicks on the button, toggle between hiding and showing the dropdown content */

function myFunction() {
  document.getElementById("myDropdown").classList.toggle("show");
};

// Close the dropdown if the user clicks outside of it
window.onclick = function(e) {
  if (!e.target.matches('.dropbtn')) {
    var myDropdown = document.getElementById("myDropdown");
    if (myDropdown.classList.contains('show')) {
      myDropdown.classList.remove('show');
    };
  };
};

document.getElementById('showWest').onclick=function(){
  if (document.getElementById('west').style.display == "none"){
    document.getElementById('west').style.display='';
  }else{
    document.getElementById('west').style.display='none';
  };

};





//My name is Ethan and I hate javascript ;)