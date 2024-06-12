
document.getElementById('login').addEventListener('mouseover', function () {
    document.getElementById('login').style.borderRadius = '50px 50px 0px 50px';
    document.getElementById('signup').style.borderRadius = '0px 50px 50px 50px';
    document.getElementById('btns1').style.backgroundColor = '#e9e643';
    // document.getElementById('signup').style.backgroundColor = '';
});
document.getElementById('signup').addEventListener('mouseover', function () {
    document.getElementById('signup').style.borderRadius = '50px 50px 50px 0px';
    document.getElementById('login').style.borderRadius = '50px 0px 50px 50px';
    
    // document.getElementById('login').style.backgroundColor = '';
});


document.getElementById('signup').addEventListener('mouseout', function() {
    document.getElementById('signup').style.borderRadius = '0px 50px 50px 0px';
    document.getElementById('login').style.borderRadius = '50px 0px 0px 50px';
});
document.getElementById('login').addEventListener('mouseout', function() {
    document.getElementById('signup').style.borderRadius = '0px 50px 50px 0px';
    document.getElementById('login').style.borderRadius = '50px 0px 0px 50px';
});


// const div = document.getElementById('myDiv');

// div.addEventListener('mousemove', function(event) {
//     const mouseX = event.pageX - this.offsetLeft;
//     const mouseY = event.pageY - this.offsetTop;
//     const gradientColor = `radial-gradient(circle at ${mouseX}px ${mouseY}px, #ff0000, #00ff00)`;
//     this.style.background = gradientColor;
// });

// div.addEventListener('mouseleave', function() {
//     this.style.background = '#ffffff'; // Возврат к начальному цвету при выходе за пределы div
// });


const divs = document.querySelectorAll('.area');

divs.forEach(div => {
    div.addEventListener('mousemove', function(event) {
        const mouseX = event.pageX - this.offsetLeft;
        const mouseY = event.pageY - this.offsetTop;
        const gradientColor = `radial-gradient(circle at ${mouseX}px ${mouseY}px, #ffffff, #e9e643, transparent 160%)`;
        this.style.background = gradientColor;
    });

    div.addEventListener('mouseleave', function() {
        this.style.background = '#ffffff'; // Возврат к начальному цвету при выходе за пределы div
    });
    
});


function changeLanguage(lang) {
    window.location.href = window.location.href + "?lang=" + lang;
  }