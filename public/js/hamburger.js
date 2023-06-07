let hamburgerIsopen = false ;

function openHamburger() {
    let hamburgerNavContainer = document.getElementById("hamburger-nav-container");
    if(!hamburgerIsopen) {
        hamburgerNavContainer.style.display = 'block';
        hamburgerIsopen = true;
    } else {
        hamburgerNavContainer.style.display = "none";
        hamburgerIsopen = false;
    }
}