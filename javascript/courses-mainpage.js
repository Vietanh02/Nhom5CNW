let buttons = document.getElementsByClassName('button-courses-filter');
let contents = document.getElementsByClassName('tab-content');
function showContent(id){
    for (let i = 0; i < contents.length; i++) {
        contents[i].style.display = 'none';
    }
    let content = document.getElementsByClassName(id);
    for(let i = 0; i < content.length; i++){
        content[i].style.display = 'block';
    }
}
for (let i = 0; i < buttons.length; i++) {
    buttons[i].addEventListener("click", function(){
        let id = this.name;
        for (let i = 0; i < buttons.length; i++) {
            buttons[i].classList.remove("active");
        }
        this.className += " active";
        console.log(id);
        showContent(id);
    });
}
showContent('');  