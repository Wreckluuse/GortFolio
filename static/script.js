

function copyFunc() {

let copyText = document.getElementById("quoteBox");
copyText.select();
copyText.setSelectionRange(0, 99999);

navigator.clipboard.writeText(copyText.value);


}