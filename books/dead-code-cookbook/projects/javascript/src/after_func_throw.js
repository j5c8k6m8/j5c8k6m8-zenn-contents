const f = () => {
  throw 'Error';
};

try {
  f();
  console.log("Am I dead?");
} catch (e) {
  // empty
}
  