const METHODS = {
  Get: 'GET',
  Post: 'POST',
};
const RESPONSE_TYPES = {
  Json: 'json',
  Text: 'text',
  Blob: 'blob',
};

/**
 * Create a new context
 * @returns {{config: {endpoint: string}}}
 */
function createContext() {
  return {
    config: {
      endpoint: '/',
    },
    interceptors: {
      request: [],
      response: [],
    },
  };
}

/**
 * Parse request options
 * @param rawOptions
 * @returns {{}}
 */
function parseRequestOptions(rawOptions = {}) {
  const options = { ...rawOptions };
  if (!options.headers) {
    options.headers = { 'Content-Type': 'application/json; charset=UTF-8' };
  }
  if (options.body) {
    options.body = JSON.stringify(options.body);
  }
  return options;
}

/**
 * Parse response
 * @param response
 * @param options
 * @returns {any}
 */
function parseResponse(response, options = {}) {
  return new Promise((resolve, reject) => {
    if (response.ok) {
      switch (options.responseType) {
        case RESPONSE_TYPES.Text:
          response.text()
            .then((data) => {
              resolve(data);
            });
          break;
        case RESPONSE_TYPES.Blob:
          resolve.blob()
            .then((data) => {
              resolve(data);
            });
          break;
        default:
          console.log('parseResponseToJson');
          response.json()
            .then((data) => {
              resolve(data);
            });
      }
      return;
    }
    reject(response.statusText);
  });
}

/**
 * Create a http instance
 * @param config
 * @returns {{http: null, config: {endpoint: string}}|*}
 */
function createHttpInstance(config = {}) {
  const context = createContext();
  if (config) {
    context.config = Object.assign(context.config, config);
  }
  const http = {
    request(url = '/', options = {}) {
      const reqUrl = url.indexOf('http') !== -1 ? url : `${config.endpoint}${url}`;
      let reqInit = parseRequestOptions(options);
      if (context.interceptors.request.length > 0) {
        context.interceptors.request.forEach((fn) => {
          reqInit = fn(reqInit);
        });
      }
      return fetch(reqUrl, reqInit)
        .then(async (response) => {
          let data = await parseResponse(response, options);
          if (context.interceptors.response.length > 0) {
            context.interceptors.response.forEach((fn) => {
              if (data.ok) {
                data = fn(data);
              }
            });
          }

          if (data.ok) {
            return Promise.resolve(data);
          }

          return Promise.reject(data);
        })
        .catch((e) => new Promise((resolve, reject) => {
          reject(e);
        }));
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }));
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }));
    },
    interceptors: {
      request: {
        use(fn) {
          if (typeof fn === 'function') {
            context.interceptors.request.push(fn);
          } else if (fn && typeof fn === 'object' && fn.constructor === Array) { // fn is an array
            context.interceptors.request = [...fn];
          } else {
            console.log('Add interceptor request failed !');
          }
        },
      },
      response: {
        use(fn) {
          if (typeof fn === 'function') {
            context.interceptors.response.push(fn);
          } else if (fn && typeof fn === 'object' && fn.constructor === Array) { // fn is an array
            context.interceptors.response = [...fn];
          } else {
            console.log('Add interceptor response failed !');
          }
        },
      },
    },
  };
  return http;
}

document.addEventListener('DOMContentLoaded', () => {
  // Init http instance
  const http = createHttpInstance({ endpoint: 'https://jsonplaceholder.typicode.com' });
  const button = document.getElementById('btnRequest');
  const result = document.getElementById('result');
  button.addEventListener('click', () => {
    // Reset result
    result.innerHTML = 'Loading...';
    const body = {
      title: 'foo',
      body: 'bar',
      userId: 1,
    };
    const interceptorsRequest = [
      (req) => {
        if (!req.headers) {
          req.headers = {};
        }
        req.headers['X-Shop-Token'] = 'OCG';
        return req;
      },
      (req) => {
        if (!req.headers) {
          req.headers = {};
        }
        req.headers['X-Customer'] = 123;
        return req;
      },
    ];
    const interceptorsResponse = [
      (response, ok) => { // function draft
        if (!ok) return response, false;
        if (response.ok) {
          console.log(1);
          return response, true;
        }
        return response, false;
      },
      (response, ok) => {
        if (!ok) {
          return response, false;
        }
        if (response.status == 201) {
          console.log(2);
          return response, true;
          // return response, false
        }
        return response, false;
      },
    ];
    http.interceptors.request.use(interceptorsRequest);
    http.interceptors.response.use(interceptorsResponse);

    // Request
    http.post('/posts', { body })
      .then((data) => {
        console.log(data);
        result.innerHTML = `<p>Result success: ${JSON.stringify(data)}</p>`;
      })
      .catch((e) => {
        if (typeof e === 'number' && isFinite(e)) {
          result.innerHTML = `<p>Result error: Response statusCode is ${e}</p> 
                                    <br> 
                                    <p>You will be redirect to google.com</p>`;
          setTimeout(() => window.location.replace('http://www.google.com'), 1000);
        } else {
          result.innerHTML = `<p>Result error: ${JSON.stringify(e)}</p>`;
        }
      });
  });
});
