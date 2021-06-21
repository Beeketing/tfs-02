const METHODS = {
  Get: 'GET',
  Post: 'POST',
}

/**
 * Create an http instance
 */
function createHttp() {
  const http = {
    request(url, options = {}) {
      return new Promise((resolve, reject) => {
        fetch(url, options).then((response) => {
          if (response.ok) {
            response.json().then((data) => {
              resolve(data)
            })
          }
        }).catch((e) => reject(e))
      })
    },
    get(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Get }))
    },
    post(url, options = {}) {
      return http.request(url, Object.assign(options, { method: METHODS.Post }))
    }
  }

  return http
}

export {
  createHttp
}
