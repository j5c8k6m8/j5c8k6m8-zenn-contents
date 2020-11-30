var d = "Am I dead?";
(() => {
    if (d) {
        console.log(d);
    }
    return;
    function d() {
        console.log("I am zombie!");
    }
})()