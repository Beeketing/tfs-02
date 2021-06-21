import { createHttp } from "./services/http";

const btn = document.getElementById('btnRequest');
const result = document.getElementById('result');
const http = createHttp();
if (btn) {
  btn.addEventListener('click', () => {
    http.get('https://jsonplaceholder.typicode.com/posts').then((data) => {
      result.innerHTML = JSON.stringify(data);
    }).catch((e) => {
      result.innerHTML = e;
    })
  })
}
