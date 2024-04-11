const qrCodeReader = new Html5Qrcode("qr-reader");
var audio = new Audio("beep.mp3");

const config = {
    fps: 10, 
    qrbox: { width: 300, height: 300 },
    aspectRatio: 1,
};
var wait = (ms) => {
    const start = Date.now();
    let now = start;
    while (now - start < ms) {
      now = Date.now();
    }
}
function onScanSuccess(decodedText, decodedResult) {
    qrCodeReader.stop().then(() => {
        showPopup(decodedText);
        console.log("Scan arrêté après la détection d'un texte.");
    }).catch(err => console.error('Erreur lors de l\'arrêt du scan:', err));
}

function onScanFailure(error) {
    console.warn(`Erreur de scan: ${error}`);
}

function showPopup(text) {
    audio.play();
    wait(250);
    window.location.href = `/update/${text}`;
}

document.addEventListener('DOMContentLoaded', function () {
    qrCodeReader.start({ facingMode: "environment" }, config, onScanSuccess, onScanFailure);
});
