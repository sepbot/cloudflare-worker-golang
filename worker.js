require('./polyfill.js');
require('./wasm_exec.js');

addEventListener('fetch', (event) => {
  event.respondWith(handleRequest(event.request));
});

function handleRequest(req) {
  return new Promise((async (resolve, reject) => {
    try {
      const url = new URL(req.url);
      const go = new Go();
      const instance = await WebAssembly.instantiate(WASM, go.importObject);
      go.run(instance);
      switch (url.pathname) {
        case '/sign':
          sign(url.searchParams.get('message'), (answer) => {
            resolve(new Response(answer, { status: 200 }));
          });
          break;
        case '/verify':
          verify(url.searchParams.get('token'), (answer) => {
            resolve(new Response(answer, { status: 200 }));
          });
          break;
        default:
          resolve(new Response('', { status: 404 }));
          break;
      }
    } catch (e) {
      console.log(e);
      reject(new Response(JSON.stringify(e), { status: 500 }));
    }
  }));
}
