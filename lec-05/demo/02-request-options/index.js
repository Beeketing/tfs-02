const METHODS = {
  Get: 'GET',
  Post: 'POST',
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
      const reqInit = parseRequestOptions(options);
      return fetch(reqUrl, reqInit)
        .then((response) => response.json())
        .catch((e) => Promise.reject(e));
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }));
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }));
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

    // Request
    http.post('/posts', { body }).then((response) => {
      result.innerHTML = `<p>Result success: ${JSON.stringify(response)}`;
    }).catch((e) => {
      result.innerHTML = `<p>Result error: ${JSON.stringify(e)}`;
    });
  });
});
